package middleware

import (
	"testing"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

const userAndRole = enum.SystemRoleKeyManageUserAndRole

// rolesWith builds a snapshot with a single system-role grant, nothing else.
func rolesWith(global *Permissions[bool], group *Permissions[[]string]) UserAccessCache {
	var r UserAccessCache
	r.SystemRole.GlobalRole = map[enum.SystemRoleKey]Permissions[bool]{}
	r.SystemRole.GroupRoles = map[enum.SystemRoleKey]Permissions[[]string]{}
	if global != nil {
		r.SystemRole.GlobalRole[userAndRole] = *global
	}
	if group != nil {
		r.SystemRole.GroupRoles[userAndRole] = *group
	}
	return r
}

func TestHasSystemPermission(t *testing.T) {
	tests := []struct {
		name  string
		roles UserAccessCache
		verb  PermissionVerb
		want  bool
	}{
		{
			name: "total control covers everything",
			roles: UserAccessCache{
				ModuleAccess: map[enum.ModuleAccess]ModuleFlags{"fixedAsset": {IsTotalControl: true}},
			},
			verb: PermEdit,
			want: true,
		},
		{
			name:  "global edit grant",
			roles: rolesWith(&Permissions[bool]{Edit: true}, nil),
			verb:  PermEdit,
			want:  true,
		},
		{
			name:  "group edit grant (non-empty group list)",
			roles: rolesWith(nil, &Permissions[[]string]{Edit: []string{"7"}}),
			verb:  PermEdit,
			want:  true,
		},
		{
			name:  "group grant but empty list is not a grant",
			roles: rolesWith(nil, &Permissions[[]string]{Edit: []string{}}),
			verb:  PermEdit,
			want:  false,
		},
		{
			name:  "holds view but asked for edit",
			roles: rolesWith(&Permissions[bool]{View: true}, nil),
			verb:  PermEdit,
			want:  false,
		},
		{
			name:  "no grant at all",
			roles: rolesWith(nil, nil),
			verb:  PermEdit,
			want:  false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.roles.HasSystemPermission(userAndRole, tc.verb); got != tc.want {
				t.Errorf("HasSystemPermission(%q, %v) = %v, want %v", userAndRole, tc.verb, got, tc.want)
			}
		})
	}
}

// This is the case the whole rbac.Scope redesign exists for: a group-scoped grant covers ALL four
// verbs, not just View. Scope's old reduction only ever read grp.View, silently dropping a holder's
// group-scoped Edit/Create/Delete. HasGroupPermission reads the raw snapshot directly, so there is
// nothing to drop.
func TestHasGroupPermission(t *testing.T) {
	roles := rolesWith(nil, &Permissions[[]string]{
		View: []string{"7"},
		Edit: []string{"7"}, // the exact grant Scope used to lose
	})

	if !roles.HasGroupPermission(userAndRole, 7, PermEdit) {
		t.Error("group-scoped Edit grant on group 7 must be visible — this is the CRUD-drop bug")
	}
	if roles.HasGroupPermission(userAndRole, 8, PermEdit) {
		t.Error("group 8 was never granted Edit; must not pass")
	}
	if !roles.HasGroupPermission(userAndRole, 7, PermView) {
		t.Error("View grant on group 7 should still work as before")
	}

	tc := UserAccessCache{ModuleAccess: map[enum.ModuleAccess]ModuleFlags{"fixedAsset": {IsTotalControl: true}}}
	if !tc.HasGroupPermission(userAndRole, 999, PermEdit) {
		t.Error("total control covers every group, including ones with no grant row at all")
	}
}

func TestGroupIDsWithPermission_UnionsAcrossKeys(t *testing.T) {
	var roles UserAccessCache
	roles.SystemRole.GroupRoles = map[enum.SystemRoleKey]Permissions[[]string]{
		enum.SystemRoleKeyManageUserAndRole:     {View: []string{"1", "2"}},
		enum.SystemRoleKeyManageMasterReference: {View: []string{"2", "3"}}, // 2 overlaps — must dedupe
		enum.SystemRoleKeyImportData:            {Edit: []string{"4"}},      // wrong verb, must not appear
	}

	got := roles.GroupIDsWithPermission(
		PermView,
		enum.SystemRoleKeyManageUserAndRole,
		enum.SystemRoleKeyManageMasterReference,
		enum.SystemRoleKeyImportData,
	)

	want := map[int64]bool{1: true, 2: true, 3: true}
	if len(got) != len(want) {
		t.Fatalf("GroupIDsWithPermission = %v, want exactly %v", got, want)
	}
	for _, id := range got {
		if !want[id] {
			t.Errorf("unexpected group id %d in result %v", id, got)
		}
	}
}

func TestHasAnySystemPermission_GlobalOnly(t *testing.T) {
	roles := rolesWith(nil, &Permissions[[]string]{View: []string{"7"}}) // GROUP grant only

	if roles.HasAnySystemPermission(PermView, userAndRole) {
		t.Error("HasAnySystemPermission must ignore group-scoped grants — it answers the GLOBAL question only")
	}

	globalOnly := rolesWith(&Permissions[bool]{View: true}, nil)
	if !globalOnly.HasAnySystemPermission(PermView, userAndRole) {
		t.Error("a global View grant should pass")
	}
}

func TestTransactionRoleHelpers(t *testing.T) {
	const tx = enum.FixedAssetTransactionRoleTransactionKeyBorrowAssignRole
	var roles UserAccessCache
	roles.TransactionRole = map[enum.FixedAssetTransactionRoleTransactionKey]map[enum.TransactionRolesType][]string{
		tx: {
			enum.TransactionRolesTypeManager:    {"10"},
			enum.TransactionRolesTypeMonitoring: {"20"},
		},
	}

	if !roles.HasTransactionRole(tx, 10, enum.TransactionRolesTypeManager) {
		t.Error("should hold Manager on group 10")
	}
	if roles.HasTransactionRole(tx, 10, enum.TransactionRolesTypeMonitoring) {
		t.Error("group 10 is Manager-only, not Monitoring")
	}
	if !roles.HasTransactionRole(tx, 20, enum.TransactionRolesTypeManager, enum.TransactionRolesTypeMonitoring) {
		t.Error("group 20 should pass an OR-check that includes Monitoring")
	}

	ids := roles.GroupIDsWithTransactionRole(tx, enum.TransactionRolesTypeManager, enum.TransactionRolesTypeMonitoring)
	if len(ids) != 2 {
		t.Errorf("GroupIDsWithTransactionRole = %v, want both 10 and 20", ids)
	}
}

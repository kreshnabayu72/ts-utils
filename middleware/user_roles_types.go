package middleware

import (
	"fmt"
	"slices"
	"strconv"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

// UserAccessCache is the fixed-asset RBAC snapshot: a user's roles already reduced to the
// shape authorization needs, so a consumer filters without joining the role tables. It is the
// contract between iam (which builds, caches and serves it) and every consumer (which fetches
// it via LoadUserRoles). The shape follows BETS-V1 so the gating logic stays portable:
//   - GroupRoles carry the list of group_ids the user may act on per permission (group-scoped).
//   - GlobalRole carries booleans (tenant-wide, not group-scoped).
//   - TransactionRole maps transaction -> role type -> group_ids.
type UserAccessCache struct {
	ID                   int                               `json:"_id"`
	FullName             string                            `json:"fullName"`
	FirstName            string                            `json:"firstName"`
	LastName             string                            `json:"lastName"`
	Email                string                            `json:"email"`
	ClientID             string                            `json:"clientId"`
	ModuleAccess         map[enum.ModuleAccess]ModuleFlags `json:"moduleAccess"`
	FixedAssetSystemRole struct {
		GlobalRole map[enum.SystemRoleKey]Permissions[bool]     `json:"globalRole"`
		GroupRoles map[enum.SystemRoleKey]Permissions[[]string] `json:"groupRoles"`
	} `json:"fixedAssetSystemRole"`
	GlobalSettingsSystemRole struct {
		GlobalRole map[enum.SystemRoleKey]Permissions[bool] `json:"globalRole"`
	} `json:"globalSettingsSystemRole"`
	AdminConsoleSystemRole struct {
		GlobalRole map[enum.SystemRoleKey]Permissions[bool]     `json:"globalRole"`
		GroupRoles map[enum.SystemRoleKey]Permissions[[]string] `json:"groupRoles"`
	} `json:"adminConsoleSystemRole"`
	FixedAssetTransactionRole map[enum.FixedAssetTransactionRoleTransactionKey]map[enum.TransactionRolesType][]string `json:"fixedAssetTransactionRole"`
}

// ModuleFlags is the per-module top-level access (e.g. moduleAccess.fixedAsset).
type ModuleFlags struct {
	IsTotalControl bool   `json:"isTotalControl"`
	IsReadOnly     bool   `json:"isReadOnly"`
	Plan           string `json:"plan"`
}

// PermissionVerb selects one column of a Permissions tuple.
type PermissionVerb int

const (
	PermCreate PermissionVerb = iota
	PermView
	PermEdit
	PermDelete
)

func pickPermission[T any](p Permissions[T], verb PermissionVerb) T {
	switch verb {
	case PermCreate:
		return p.Create
	case PermView:
		return p.View
	case PermEdit:
		return p.Edit
	default: // PermDelete
		return p.Delete
	}
}

// IsTotalControl reports whether the user is a total-control admin of `product` — the grant that
// covers every action in that module, including managing users and roles.
func (r UserAccessCache) IsTotalControl(product enum.ModuleAccess) bool {
	m, ok := r.ModuleAccess[product]
	return ok && m.IsTotalControl
}

// IsReadOnly reports whether the user holds the read-only grant of `product`.
func (r UserAccessCache) IsReadOnly(product enum.ModuleAccess) bool {
	m, ok := r.ModuleAccess[product]
	return ok && m.IsReadOnly
}

// HasSystemPermissionFixedAsset reports whether the user holds `verb` on system role `key` ANYWHERE — either
// tenant-wide (global) or in at least one group. Total control covers everything.
//
// This is the snapshot equivalent of BETS-V1's GlobalSystemRoleRestrictionSQL, minus the SQL: the
// answer is already in the cached snapshot iam served, so no query runs. Gating logic lives here,
// on the shared contract type, so iam (which builds the snapshot itself) and every consumer (which
// fetches it via LoadUserRoles) ask the same question the same way.
func (r UserAccessCache) HasSystemPermissionFixedAsset(key enum.SystemRoleKey, verb PermissionVerb) bool {
	if r.IsTotalControl(enum.ModuleAccessFixedAsset) {
		return true
	}
	if g, ok := r.FixedAssetSystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
		return true
	}
	if g, ok := r.FixedAssetSystemRole.GroupRoles[key]; ok && len(pickPermission(g, verb)) > 0 {
		return true
	}
	return false
}

// HasGroupPermissionFixedAsset reports whether the user holds `verb` on system role `key` in the SPECIFIC
// group `groupID` — unlike HasSystemPermissionFixedAsset's "anywhere" answer, this checks membership in one
// named group. Total control still covers everything.
//
// This is the check a per-row write decision needs (e.g. "may I edit this asset, which is in group
// 7"): HasSystemPermissionFixedAsset("do you hold this anywhere") would wrongly admit someone who holds it in
// a DIFFERENT group.
func (r UserAccessCache) HasGroupPermissionFixedAsset(
	key enum.SystemRoleKey,
	groupID int64,
	verb PermissionVerb,
) bool {
	if r.IsTotalControl(enum.ModuleAccessFixedAsset) {
		return true
	}
	g, ok := r.FixedAssetSystemRole.GroupRoles[key]
	if !ok {
		return false
	}
	return slices.Contains(pickPermission(g, verb), strconv.FormatInt(groupID, 10))
}

// HasAnyGroupPermissionFixedAsset reports whether the user holds `verb` on the SPECIFIC group `groupID`, via
// ANY of `keys`. Variadic sibling of HasGroupPermissionFixedAsset for a caller ORing several system-role keys
// (V1's available-asset-list breadth) against one group, without materializing a full id list first.
func (r UserAccessCache) HasAnyGroupPermissionFixedAsset(
	groupID int64,
	verb PermissionVerb,
	keys ...enum.SystemRoleKey,
) bool {
	if r.IsTotalControl(enum.ModuleAccessFixedAsset) {
		return true
	}
	gid := strconv.FormatInt(groupID, 10)
	for _, key := range keys {
		if g, ok := r.FixedAssetSystemRole.GroupRoles[key]; ok && slices.Contains(pickPermission(g, verb), gid) {
			return true
		}
	}
	return false
}

// GroupIDsWithPermission returns the group ids where the user holds `verb` on ANY of `keys` —
// union across all of them, deduped. Variadic so a caller ORing several system-role keys (V1's
// available-asset-list gate ORs purchaseInformation/tag/purchaseDocument alongside assets) writes
// one call instead of one append per key.
//
// Total control is NOT expanded into "every group id" here — there is no group list to enumerate
// for a tenant-wide grant. Callers needing a see-all shortcut check IsTotalControl/IsReadOnly (or
// HasAnySystemPermissionFixedAsset for a global key) separately, and skip the group filter entirely.
func (r UserAccessCache) GroupIDsWithPermission(verb PermissionVerb, keys ...enum.SystemRoleKey) []int64 {
	seen := map[int64]bool{}
	var out []int64
	for _, key := range keys {
		g, ok := r.FixedAssetSystemRole.GroupRoles[key]
		if !ok {
			continue
		}
		for _, raw := range pickPermission(g, verb) {
			id, err := strconv.ParseInt(raw, 10, 64)
			if err != nil || seen[id] {
				continue
			}
			seen[id] = true
			out = append(out, id)
		}
	}
	return out
}

// PermissionEntry is the unit of lookup for GroupIDsFor: either a system-role permission or a
// transaction role. Fields are unexported and the isTransaction discriminant is set only by the
// two constructors below, so a caller cannot build one that names both or neither.
type PermissionEntry struct {
	isTransaction   bool
	verb            PermissionVerb
	systemRole      enum.SystemRoleKey
	transactionRole enum.FixedAssetTransactionRoleTransactionKey
	transactionType enum.TransactionRolesType
}

// ForSystemRole builds a PermissionEntry that looks up group ids by system-role permission.
func ForSystemRole(verb PermissionVerb, key enum.SystemRoleKey) PermissionEntry {
	return PermissionEntry{verb: verb, systemRole: key}
}

// ForTransactionRole builds a PermissionEntry that looks up group ids by transaction role.
func ForTransactionRole(
	transactionRole enum.FixedAssetTransactionRoleTransactionKey,
	transactionType enum.TransactionRolesType,
) PermissionEntry {
	return PermissionEntry{isTransaction: true, transactionRole: transactionRole, transactionType: transactionType}
}

// GroupIDsFor returns the group ids where the user holds each entry — system-role permission or
// transaction role — deduped across all entries.
//
//	roles.GroupIDsFor(
//	    middleware.ForSystemRole(PermView, Assets),
//	    middleware.ForTransactionRole(borrowTx, Manager),
//	)
func (r UserAccessCache) GroupIDsFor(entries ...PermissionEntry) []int64 {
	seen := map[int64]bool{}
	var out []int64
	for _, e := range entries {
		if e.isTransaction {
			gids := r.GroupIDsWithTransactionRole(e.transactionRole, e.transactionType)
			for _, id := range gids {
				if !seen[id] {
					seen[id] = true
					out = append(out, id)
				}
			}
			continue
		}
		g, ok := r.FixedAssetSystemRole.GroupRoles[e.systemRole]
		if !ok {
			continue
		}
		for _, raw := range pickPermission(g, e.verb) {
			id, err := strconv.ParseInt(raw, 10, 64)
			if err != nil || seen[id] {
				continue
			}
			seen[id] = true
			out = append(out, id)
		}
	}
	return out
}

// HasAnySystemPermissionFixedAsset reports whether the user holds `verb` GLOBALLY (tenant-wide, no group_id)
// on any of `keys` in the fixed-asset module. Unlike HasSystemPermissionFixedAsset, this does NOT look at
// group-scoped grants — it answers "does one of these keys, held globally, waive the group filter
// entirely", which is what a global key means: it has no group_id to filter by, so holding it
// covers every row (V1's depreciationGroup/depreciationMethod, joined with no group condition at
// all).
func (r UserAccessCache) HasAnySystemPermissionFixedAsset(verb PermissionVerb, keys ...enum.SystemRoleKey) bool {
	if r.IsTotalControl(enum.ModuleAccessFixedAsset) {
		return true
	}
	for _, key := range keys {
		if g, ok := r.FixedAssetSystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
			return true
		}
	}
	return false
}

// HasGlobalSettingsPermission reports whether the user holds `verb` on system role `key` in the
// global-settings module. Unlike HasSystemPermissionFixedAsset, this reads GlobalSettingsSystemRole,
// not FixedAssetSystemRole, so a fixed-asset grant cannot accidentally satisfy a global-settings
// gate. Global-settings total control passes everything.
func (r UserAccessCache) HasGlobalSettingsPermission(key enum.SystemRoleKey, verb PermissionVerb) bool {
	if r.IsTotalControl(enum.ModuleAccessGlobalSettings) {
		return true
	}
	if g, ok := r.GlobalSettingsSystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
		return true
	}
	return false
}

// HasSystemPermissionAdminConsole reports whether the user holds `verb` on system role `key` ANYWHERE — either
// tenant-wide (global) or in at least one group. Total control covers everything.
func (r UserAccessCache) HasSystemPermissionAdminConsole(key enum.SystemRoleKey, verb PermissionVerb) bool {
	if r.IsTotalControl(enum.ModuleAccessAdminConsole) {
		return true
	}
	if g, ok := r.AdminConsoleSystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
		return true
	}
	if g, ok := r.AdminConsoleSystemRole.GroupRoles[key]; ok && len(pickPermission(g, verb)) > 0 {
		return true
	}
	return false
}

// HasGroupPermissionAdminConsole reports whether the user holds `verb` on system role `key` in the SPECIFIC
// group `groupID` — unlike HasSystemPermissionAdminConsole's "anywhere" answer, this checks membership in one
// named group. Total control still covers everything.
func (r UserAccessCache) HasGroupPermissionAdminConsole(
	key enum.SystemRoleKey,
	groupID int64,
	verb PermissionVerb,
) bool {
	if r.IsTotalControl(enum.ModuleAccessAdminConsole) {
		return true
	}
	g, ok := r.AdminConsoleSystemRole.GroupRoles[key]
	if !ok {
		return false
	}
	return slices.Contains(pickPermission(g, verb), strconv.FormatInt(groupID, 10))
}

// HasAnyGroupPermissionAdminConsole reports whether the user holds `verb` on the SPECIFIC group `groupID`, via
// ANY of `keys`. Variadic sibling of HasGroupPermissionAdminConsole for a caller ORing several system-role keys
// against one group, without materializing a full id list first.
func (r UserAccessCache) HasAnyGroupPermissionAdminConsole(
	groupID int64,
	verb PermissionVerb,
	keys ...enum.SystemRoleKey,
) bool {
	if r.IsTotalControl(enum.ModuleAccessAdminConsole) {
		return true
	}
	gid := strconv.FormatInt(groupID, 10)
	for _, key := range keys {
		if g, ok := r.AdminConsoleSystemRole.GroupRoles[key]; ok && slices.Contains(pickPermission(g, verb), gid) {
			return true
		}
	}
	return false
}

// HasTransactionRoleFixedAsset reports whether the user holds ANY of `types` on `transaction` in the
// specific group `groupID`.
func (r UserAccessCache) HasTransactionRoleFixedAsset(
	transaction enum.FixedAssetTransactionRoleTransactionKey,
	groupID int64,
	types ...enum.TransactionRolesType,
) bool {
	if r.IsTotalControl(enum.ModuleAccessFixedAsset) {
		return true
	}
	byType, ok := r.FixedAssetTransactionRole[transaction]
	if !ok {
		return false
	}
	gid := strconv.FormatInt(groupID, 10)
	for _, t := range types {
		if slices.Contains(byType[t], gid) {
			return true
		}
	}
	return false
}

// GroupIDsWithTransactionRole returns the group ids where the user holds ANY of `types` on
// `transaction` — union across all given types, deduped.
func (r UserAccessCache) GroupIDsWithTransactionRole(
	transaction enum.FixedAssetTransactionRoleTransactionKey,
	types ...enum.TransactionRolesType,
) []int64 {
	byType, ok := r.FixedAssetTransactionRole[transaction]
	if !ok {
		return nil
	}
	seen := map[int64]bool{}
	var out []int64
	for _, t := range types {
		for _, raw := range byType[t] {
			id, err := strconv.ParseInt(raw, 10, 64)
			if err != nil || seen[id] {
				continue
			}
			seen[id] = true
			out = append(out, id)
		}
	}
	return out
}

// Permissions is the CRUD permission tuple. T is bool for global (tenant-wide) roles
// and []string (group_ids) for group-scoped roles.
type Permissions[T any] struct {
	Create T `json:"create"`
	View   T `json:"view"`
	Edit   T `json:"edit"`
	Delete T `json:"delete"`
}

// UserRolesCacheKey is the Valkey key holding a user's UserAccessCache snapshot.
//
// Only iam reads or writes this key — it serves the snapshot from the cache, rebuilds it from
// the role tables on a miss, and drops it when a user's roles change. Consumers ask iam over
// HTTP (see LoadUserRoles) rather than touching the cache. The key lives here because it is
// part of the contract between iam and the shared snapshot type.
func UserRolesCacheKey(userID string) string {
	return fmt.Sprintf("user-fixed-asset-roles:%s", userID)
}

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
	ID           int                               `json:"_id"`
	FullName     string                            `json:"fullName"`
	Email        string                            `json:"email"`
	ClientID     string                            `json:"clientId"`
	ModuleAccess map[enum.ModuleAccess]ModuleFlags `json:"moduleAccess"`
	SystemRole   struct {
		GlobalRole map[enum.FixedAssetSystemRoleKey]Permissions[bool]     `json:"globalRole"`
		GroupRoles map[enum.FixedAssetSystemRoleKey]Permissions[[]string] `json:"groupRoles"`
	} `json:"systemRole"`
	TransactionRole map[enum.FixedAssetTransactionRoleTransactionKey]map[enum.TransactionRolesType][]string `json:"transactionRole"`
}

// ModuleFlags is the per-module top-level access (e.g. moduleAccess.fixedAsset).
type ModuleFlags struct {
	IsTotalControl bool `json:"isTotalControl"`
	IsReadOnly     bool `json:"isReadOnly"`
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

// IsTotalControl reports whether the user is a total-control admin of the fixed-asset module —
// the grant that covers every action, including managing users and roles.
func (r UserAccessCache) IsTotalControl() bool {
	m, ok := r.ModuleAccess[enum.ModuleAccessFixedAsset]
	return ok && m.IsTotalControl
}

// IsReadOnly reports whether the user holds the read-only grant of the fixed-asset module.
func (r UserAccessCache) IsReadOnly() bool {
	m, ok := r.ModuleAccess[enum.ModuleAccessFixedAsset]
	return ok && m.IsReadOnly
}

// HasSystemPermission reports whether the user holds `verb` on system role `key` ANYWHERE — either
// tenant-wide (global) or in at least one group. Total control covers everything.
//
// This is the snapshot equivalent of BETS-V1's GlobalSystemRoleRestrictionSQL, minus the SQL: the
// answer is already in the cached snapshot iam served, so no query runs. Gating logic lives here,
// on the shared contract type, so iam (which builds the snapshot itself) and every consumer (which
// fetches it via LoadUserRoles) ask the same question the same way.
func (r UserAccessCache) HasSystemPermission(key enum.FixedAssetSystemRoleKey, verb PermissionVerb) bool {
	if r.IsTotalControl() {
		return true
	}
	if g, ok := r.SystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
		return true
	}
	if g, ok := r.SystemRole.GroupRoles[key]; ok && len(pickPermission(g, verb)) > 0 {
		return true
	}
	return false
}

// HasGroupPermission reports whether the user holds `verb` on system role `key` in the SPECIFIC
// group `groupID` — unlike HasSystemPermission's "anywhere" answer, this checks membership in one
// named group. Total control still covers everything.
//
// This is the check a per-row write decision needs (e.g. "may I edit this asset, which is in group
// 7"): HasSystemPermission("do you hold this anywhere") would wrongly admit someone who holds it in
// a DIFFERENT group.
func (r UserAccessCache) HasGroupPermission(
	key enum.FixedAssetSystemRoleKey,
	groupID int64,
	verb PermissionVerb,
) bool {
	if r.IsTotalControl() {
		return true
	}
	g, ok := r.SystemRole.GroupRoles[key]
	if !ok {
		return false
	}
	return slices.Contains(pickPermission(g, verb), strconv.FormatInt(groupID, 10))
}

// HasAnyGroupPermission reports whether the user holds `verb` on the SPECIFIC group `groupID`, via
// ANY of `keys`. Variadic sibling of HasGroupPermission for a caller ORing several system-role keys
// (V1's available-asset-list breadth) against one group, without materializing a full id list first.
func (r UserAccessCache) HasAnyGroupPermission(
	groupID int64,
	verb PermissionVerb,
	keys ...enum.FixedAssetSystemRoleKey,
) bool {
	if r.IsTotalControl() {
		return true
	}
	gid := strconv.FormatInt(groupID, 10)
	for _, key := range keys {
		if g, ok := r.SystemRole.GroupRoles[key]; ok && slices.Contains(pickPermission(g, verb), gid) {
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
// HasAnySystemPermission for a global key) separately, and skip the group filter entirely.
func (r UserAccessCache) GroupIDsWithPermission(verb PermissionVerb, keys ...enum.FixedAssetSystemRoleKey) []int64 {
	seen := map[int64]bool{}
	var out []int64
	for _, key := range keys {
		g, ok := r.SystemRole.GroupRoles[key]
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
	systemRole      enum.FixedAssetSystemRoleKey
	transactionRole enum.FixedAssetTransactionRoleTransactionKey
	transactionType enum.TransactionRolesType
}

// ForSystemRole builds a PermissionEntry that looks up group ids by system-role permission.
func ForSystemRole(verb PermissionVerb, key enum.FixedAssetSystemRoleKey) PermissionEntry {
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
		g, ok := r.SystemRole.GroupRoles[e.systemRole]
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

// HasAnySystemPermission reports whether the user holds `verb` GLOBALLY (tenant-wide, no group_id)
// on any of `keys`. Unlike HasSystemPermission, this does NOT look at group-scoped grants — it
// answers "does one of these keys, held globally, waive the group filter entirely", which is what a
// global key means: it has no group_id to filter by, so holding it covers every row (V1's
// depreciationGroup/depreciationMethod, joined with no group condition at all).
func (r UserAccessCache) HasAnySystemPermission(verb PermissionVerb, keys ...enum.FixedAssetSystemRoleKey) bool {
	if r.IsTotalControl() {
		return true
	}
	for _, key := range keys {
		if g, ok := r.SystemRole.GlobalRole[key]; ok && pickPermission(g, verb) {
			return true
		}
	}
	return false
}

// HasTransactionRole reports whether the user holds ANY of `types` on `transaction` in the
// specific group `groupID`.
func (r UserAccessCache) HasTransactionRole(
	transaction enum.FixedAssetTransactionRoleTransactionKey,
	groupID int64,
	types ...enum.TransactionRolesType,
) bool {
	if r.IsTotalControl() {
		return true
	}
	byType, ok := r.TransactionRole[transaction]
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
	byType, ok := r.TransactionRole[transaction]
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

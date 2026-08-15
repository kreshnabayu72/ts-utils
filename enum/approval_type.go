// Package enum defines the shared string-typed constants (role keys,
// transaction types, module identifiers, statuses) used across every
// service and by the RBAC snapshot.
package enum

type ApprovalType string

const (
	ApprovalTypeOr  ApprovalType = "or"
	ApprovalTypeAnd ApprovalType = "and"
)

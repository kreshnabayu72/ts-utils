package enum

type FixedAssetTransactionRoleTransaction string

const (
	FixedAssetTransactionRoleTransactionBorrowingAssignment FixedAssetTransactionRoleTransaction = "Borrowing/Assignment"
	FixedAssetTransactionRoleTransactionTransfer            FixedAssetTransactionRoleTransaction = "Transfer"
	FixedAssetTransactionRoleTransactionDisposal            FixedAssetTransactionRoleTransaction = "Disposal"
	FixedAssetTransactionRoleTransactionAudit               FixedAssetTransactionRoleTransaction = "Audit"
	FixedAssetTransactionRoleTransactionMaintenanceRoutine  FixedAssetTransactionRoleTransaction = "Maintenance Routine"
	FixedAssetTransactionRoleTransactionRepairTicketing     FixedAssetTransactionRoleTransaction = "Repair Ticketing"
	FixedAssetTransactionRoleTransactionTracking            FixedAssetTransactionRoleTransaction = "Tracking"
)

package enum

type TransactionRolesType string

const (
	TransactionRolesTypeManager    TransactionRolesType = "Manager"
	TransactionRolesTypeMonitoring TransactionRolesType = "Monitoring"
	TransactionRolesTypeApprover   TransactionRolesType = "Approver"
	TransactionRolesTypeStaff      TransactionRolesType = "Staff"
)

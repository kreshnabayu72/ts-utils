package enum

type AssetStatus string

const (
	AssetStatusNonTransactional   AssetStatus = "Non-Transactional"
	AssetStatusAvailable          AssetStatus = "Available"
	AssetStatusAssigned           AssetStatus = "Assigned"
	AssetStatusBorrowed           AssetStatus = "Borrowed"
	AssetStatusDisposed           AssetStatus = "Disposed"
	AssetStatusDamaged            AssetStatus = "Damaged"
	AssetStatusMissing            AssetStatus = "Missing"
	AssetStatusOnAssignProcess    AssetStatus = "On Assign Process"
	AssetStatusOnBorrowingProcess AssetStatus = "On Borrowing Process"
	AssetStatusOnDisposalProcess  AssetStatus = "On Disposal Process"
	AssetStatusReportedDisposal   AssetStatus = "Reported Disposal"
	AssetStatusOnTransfer         AssetStatus = "On Transfer"
	AssetStatusExpiredLicense     AssetStatus = "Expired License"
	AssetStatusTAGReported        AssetStatus = "TAG Reported"
	AssetStatusTAGNotSet          AssetStatus = "TAG Not Set"
	AssetStatusNotPairedYet       AssetStatus = "Not Paired Yet"
)

type ApprovalStatus string

const (
	ApprovalStatusPending          ApprovalStatus = "Pending"
	ApprovalStatusNeedApproval     ApprovalStatus = "Need Approval"
	ApprovalStatusFinishedApproval ApprovalStatus = "Finished Approval"
)

type CustomFieldType string

const (
	CustomFieldTypeGeneralInformation = CustomFieldType("General Information")
	CustomFieldTypePurchase           = CustomFieldType("Purchase")
	CustomFieldTypeAccounting         = CustomFieldType("Accounting")
	CustomFieldTypeMaintenanceRoutine = CustomFieldType("Maintenance Routine")
	CustomFieldTypeAudit              = CustomFieldType("Audit")
)

type CustomFieldDataType string

const (
	CustomFieldDataTypeText          = CustomFieldDataType("Text")
	CustomFieldDataTypeTextarea      = CustomFieldDataType("Textarea")
	CustomFieldDataTypeDropdown      = CustomFieldDataType("Dropdown")
	CustomFieldDataTypeMultiDropdown = CustomFieldDataType("MultiDropdown")
	CustomFieldDataTypeRadioButton   = CustomFieldDataType("Radio button")
	CustomFieldDataTypeCheckbox      = CustomFieldDataType("Checkbox")
	CustomFieldDataTypeDate          = CustomFieldDataType("Date")
	CustomFieldDataTypeDatetime      = CustomFieldDataType("Datetime")
	CustomFieldDataTypeCurrency      = CustomFieldDataType("Currency")
	CustomFieldDataTypeNumeric       = CustomFieldDataType("Numeric")
	CustomFieldDataTypePhone         = CustomFieldDataType("Phone")
	CustomFieldDataTypePercentage    = CustomFieldDataType("Percentage")
	CustomFieldDataTypeEmail         = CustomFieldDataType("Email")
	CustomFieldDataTypeURL           = CustomFieldDataType("URL")
	CustomFieldDataTypeDocument      = CustomFieldDataType("Document")
)



type FixedAssetTransactionRoleTransactionKey string

const (
	FixedAssetTransactionRoleTransactionKeyBorrowAssignRole       FixedAssetTransactionRoleTransactionKey = "borrowAssignRole"
	FixedAssetTransactionRoleTransactionKeyTransferRole           FixedAssetTransactionRoleTransactionKey = "transferRole"
	FixedAssetTransactionRoleTransactionKeyDisposalRole           FixedAssetTransactionRoleTransactionKey = "disposalRole"
	FixedAssetTransactionRoleTransactionKeyAuditRole              FixedAssetTransactionRoleTransactionKey = "auditRole"
	FixedAssetTransactionRoleTransactionKeyMaintenanceRoutineRole FixedAssetTransactionRoleTransactionKey = "maintenanceRoutineRole"
	FixedAssetTransactionRoleTransactionKeyRepairTicketingRole    FixedAssetTransactionRoleTransactionKey = "repairTicketingRole"
	FixedAssetTransactionRoleTransactionKeyTrackingRole           FixedAssetTransactionRoleTransactionKey = "trackingRole"
	FixedAssetTransactionRoleTransactionKeyRegistrationRole       FixedAssetTransactionRoleTransactionKey = "registrationRole"
)

type FixedAssetTransactionRoleTransaction string

const (
	FixedAssetTransactionRoleTransactionBorrowingAssignment FixedAssetTransactionRoleTransaction = "Borrowing/Assignment"
	FixedAssetTransactionRoleTransactionTransfer            FixedAssetTransactionRoleTransaction = "Transfer"
	FixedAssetTransactionRoleTransactionDisposal            FixedAssetTransactionRoleTransaction = "Disposal"
	FixedAssetTransactionRoleTransactionAudit               FixedAssetTransactionRoleTransaction = "Audit"
	FixedAssetTransactionRoleTransactionMaintenanceRoutine  FixedAssetTransactionRoleTransaction = "Maintenance Routine"
	FixedAssetTransactionRoleTransactionRepairTicketing     FixedAssetTransactionRoleTransaction = "Repair Ticketing"
	FixedAssetTransactionRoleTransactionTracking            FixedAssetTransactionRoleTransaction = "Tracking"
	FixedAssetTransactionRoleTransactionRegistration        FixedAssetTransactionRoleTransaction = "Registration"
)

// TransactionStatus tracks the transaction as a whole; RequestStatus tracks one asset line
// within it. They move together, but a single line can be cancelled without ending the
// transaction, so they are separate.
type TransactionBorrowStatus string

const (
	TransactionBorrowStatusWaitingForApproval TransactionBorrowStatus = "Waiting for Approval"
	TransactionBorrowStatusWaitingForHandover TransactionBorrowStatus = "Waiting for Handover"
	TransactionBorrowStatusBorrowed           TransactionBorrowStatus = "Borrowed"
	TransactionBorrowStatusRejected           TransactionBorrowStatus = "Rejected"
	TransactionBorrowStatusReturned           TransactionBorrowStatus = "Returned"
	TransactionBorrowStatusCanceled           TransactionBorrowStatus = "Canceled"
	TransactionBorrowStatusRequestExtension   TransactionBorrowStatus = "Request Extension"
	TransactionBorrowStatusExtendApproval     TransactionBorrowStatus = "Extend Approval"
)

type RequestBorrowStatus string

const (
	RequestBorrowStatusWaitingForApproval RequestBorrowStatus = "Waiting for Approval"
	RequestBorrowStatusWaitingForHandover RequestBorrowStatus = "Waiting for Handover"
	RequestBorrowStatusApproved           RequestBorrowStatus = "Approved"
	RequestBorrowStatusBorrowed           RequestBorrowStatus = "Borrowed"
	RequestBorrowStatusRejected           RequestBorrowStatus = "Rejected"
	RequestBorrowStatusCanceled           RequestBorrowStatus = "Canceled"
	RequestBorrowStatusReturned           RequestBorrowStatus = "Returned"
	RequestBorrowStatusRequestExtension   RequestBorrowStatus = "Request Extension"
	RequestBorrowStatusExtendApproval     RequestBorrowStatus = "Extend Approval"
)

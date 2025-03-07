package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TransactionRoleType represents the possible types of transaction role
type TransactionRoleType string

const (
	TransactionRoleTypeReceptionSupplierReturn TransactionRoleType = "receptionSupplierReturn"
	TransactionRoleTypeCheckoutInternalReturn  TransactionRoleType = "checkoutInternalReturn"
	TransactionRoleTypeTransfer                TransactionRoleType = "transfer"
	TransactionRoleTypeDefect                  TransactionRoleType = "defect"
	TransactionRoleTypeExpired                 TransactionRoleType = "expired"
	TransactionRoleTypeTracking                TransactionRoleType = "tracking"
	TransactionRoleTypeMissing                 TransactionRoleType = "missing"
	TransactionRoleTypeDisposal                TransactionRoleType = "disposal"
	TransactionRoleTypeStockOpname             TransactionRoleType = "stockOpname"
)

// TransactionRoleName represents the name of the transaction role's TransactionRoleType
type TransactionRoleName string

var TransactionRoleNameMap = map[TransactionRoleType]TransactionRoleName{
	TransactionRoleTypeReceptionSupplierReturn: "Reception / Supplier Return",
	TransactionRoleTypeCheckoutInternalReturn:  "Checkout / Internal Return",
	TransactionRoleTypeTransfer:                "Transfer",
	TransactionRoleTypeDefect:                  "Defect",
	TransactionRoleTypeExpired:                 "Expired",
	TransactionRoleTypeMissing:                 "Missing",
	TransactionRoleTypeDisposal:                "Disposal",
	TransactionRoleTypeStockOpname:             "Stock Opname",
	TransactionRoleTypeTracking:                "Tracking",
}

// TransactionManageType represents the transaction role's manage type
type TransactionManageType string

const (
	TransactionManageTypeAdmin            TransactionManageType = "Admin"
	TransactionManageTypeMonitoringReport TransactionManageType = "Monitoring & Report"
	TransactionManageTypeStaff            TransactionManageType = "Staff"
	TransactionManageTypeApprover         TransactionManageType = "Approver"
)

// TransactionApprovalType represents the possible values for approval type
type TransactionApprovalType string

const (
	TransactionApprovalTypeOr  TransactionApprovalType = "Or"
	TransactionApprovalTypeAnd TransactionApprovalType = "And"
)

// UserTransactionRole represents the user's individual transaction role data
type UserTransactionRole struct {
	ID                  primitive.ObjectID       `json:"_id,omitempty"       bson:"_id,omitempty"`
	User                primitive.ObjectID       `json:"user"                bson:"user"`
	IsActive            bool                     `json:"isActive"            bson:"isActive"`
	IsTotalControl      bool                     `json:"isTotalControl"      bson:"isTotalControl"`
	IsReadOnly          bool                     `json:"isReadOnly"          bson:"isReadOnly"`
	AdministrativeGroup primitive.ObjectID       `json:"administrativeGroup" bson:"administrativeGroup"`
	RoleName            TransactionRoleName      `json:"roleName"            bson:"roleName"`
	RoleType            TransactionRoleType      `json:"roleType"            bson:"roleType"`
	ManageType          TransactionManageType    `json:"manageType"          bson:"manageType"`
	ApprovalLevel       *int                     `json:"approvalLevel"       bson:"approvalLevel"`
	ApprovalType        *TransactionApprovalType `json:"approvalType"        bson:"approvalType"`
	UpdatedAt           time.Time                `json:"updatedAt"           bson:"updatedAt"`
	CreatedAt           time.Time                `json:"createdAt"           bson:"createdAt"`
}

package new_model

import (
	"time"
)

type Group struct {
	ID                                 int        `json:"id"                                 column:"groups.id"`
	Name                               string     `json:"name"                               column:"groups.name"`
	Code                               *string    `json:"code"                               column:"groups.code"`
	Level                              int        `json:"level"                              column:"groups.level"`
	IsDisposable                       bool       `json:"isDisposable"                       column:"groups.is_disposable"`
	AncestorID                         int        `json:"ancestorId"                         column:"groups.ancestor_id"`
	ParentID                           *int       `json:"parentId"                           column:"groups.parent_id"`
	ParentQuotaSourceID                int        `json:"parentQuotaSourceId"                column:"groups.parent_quota_source_id"`
	BorrowingRoleManagedByParent       bool       `json:"borrowingRoleManagedByParent"       column:"groups.borrowing_role_managed_by_parent"`
	TransferRoleManagedByParent        bool       `json:"transferRoleManagedByParent"        column:"groups.transfer_role_managed_by_parent"`
	DisposalRoleManagedByParent        bool       `json:"disposalRoleManagedByParent"        column:"groups.disposal_role_managed_by_parent"`
	DamagedRoleManagedByParent         bool       `json:"damagedRoleManagedByParent"         column:"groups.damaged_role_managed_by_parent"`
	MissingRoleManagedByParent         bool       `json:"missingRoleManagedByParent"         column:"groups.missing_role_managed_by_parent"`
	RepairTicketingRoleManagedByParent bool       `json:"repairTicketingRoleManagedByParent" column:"groups.repair_ticketing_role_managed_by_parent"`
	TrackingRoleManagedByParent        bool       `json:"trackingRoleManagedByParent"        column:"groups.tracking_role_managed_by_parent"`
	AuditRoleManagedByParent           bool       `json:"auditRoleManagedByParent"           column:"groups.audit_role_managed_by_parent"`
	MaintenanceRoleManagedByParent     bool       `json:"maintenanceRoleManagedByParent"     column:"groups.maintenance_role_managed_by_parent"`
	AssetControlRoleManagedByParent    bool       `json:"assetControlRoleManagedByParent"    column:"groups.asset_control_role_managed_by_parent"`
	IsDeleted                          bool       `json:"isDeleted"                          column:"groups.is_deleted"`
	DeletedAt                          *time.Time `json:"deletedAt"                          column:"groups.deleted_at"`
	CreatedAt                          time.Time  `json:"createdAt"                          column:"groups.created_at"`
	UpdatedAt                          time.Time  `json:"updatedAt"                          column:"groups.updated_at"`
}

type GroupQuota struct {
	ID                   int            `json:"id"                   column:"group_quotas.id"`
	GroupID              int            `json:"groupId"              column:"group_quotas.group_id"`
	Type                 GroupQuotaType `json:"type"                 column:"group_quotas.type"`
	Total                int            `json:"total"                column:"group_quotas.total"`
	Used                 int            `json:"used"                 column:"group_quotas.used"`
	UsedByThisGroup      int            `json:"usedByThisGroup"      column:"group_quotas.used_by_this_group"`
	AllocatedToSubgroups int            `json:"allocatedToSubgroups" column:"group_quotas.allocated_to_sub_groups"`
	Available            *int           `json:"available"            column:"group_quotas.available"` // nil is "Not Set"
	CreatedAt            time.Time      `json:"createdAt"            column:"group_quotas.created_at"`
	UpdatedAt            time.Time      `json:"updatedAt"            column:"group_quotas.updated_at"`
}

type GroupApprovalType struct {
	ID           int                  `json:"id"           column:"id"`
	GroupID      int                  `json:"groupId"      column:"group_id"`
	Level        int                  `json:"level"        column:"level"`
	Transaction  GroupTransactionType `json:"transaction"  column:"transaction"`
	ApprovalType ApprovalType         `json:"approvalType" column:"approval_type"`
	CreatedAt    time.Time            `json:"createdAt"    column:"created_at"`
	UpdatedAt    time.Time            `json:"updatedAt"    column:"updated_at"`
}

package new_model

import "time"

type Group struct {
	ID   int     `json:"id" column:"id"`
	Name string  `json:"name" column:"name"`
	Code *string `json:"code" column:"code"`
	Key  int     `json:"key"  column:"key"`

	Level *int `json:"level" column:"level"`

	IsDisposable        bool `json:"isDisposable" column:"is_disposable"`
	ParentID            int  `json:"parentId" column:"parent_id"`
	ParentQuotaSourceID int  `json:"parentQuotaSourceId" column:"parent_quota_source_id"`

	BorrowingRoleManagedByParent       bool `json:"borrowingRoleManagedByParent" column:"borrowing_role_managed_by_parent"`
	TransferRoleManagedByParent        bool `json:"transferRoleManagedByParent" column:"transfer_role_managed_by_parent"`
	DisposalRoleManagedByParent        bool `json:"disposalRoleManagedByParent" column:"disposal_role_managed_by_parent"`
	DamagedRoleManagedByParent         bool `json:"damagedRoleManagedByParent" column:"damaged_role_managed_by_parent"`
	MissingRoleManagedByParent         bool `json:"missingRoleManagedByParent" column:"missing_role_managed_by_parent"`
	RepairTicketingRoleManagedByParent bool `json:"repairTicketingRoleManagedByParent" column:"repair_ticketing_role_managed_by_parent"`
	TrackingRoleManagedByParent        bool `json:"trackingRoleManagedByParent" column:"tracking_role_managed_by_parent"`
	AuditRoleManagedByParent           bool `json:"auditRoleManagedByParent" column:"audit_role_managed_by_parent"`
	MaintenanceRoleManagedByParent     bool `json:"maintenanceRoleManagedByParent" column:"maintenance_role_managed_by_parent"`
	AssetControlRoleManagedByParent    bool `json:"assetControlRoleManagedByParent" column:"asset_control_role_managed_by_parent"`

	IsDeleted bool       `json:"isDeleted" column:"is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type Group struct {
	ID                                 int        `json:"id"                                 column:"groups.id"                                      sql:"BIGINT PRIMARY KEY"`
	Name                               string     `json:"name"                               column:"groups.name"                                    sql:"VARCHAR UNIQUE NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	Code                               *string    `json:"code"                               column:"groups.code"                                    sql:"VARCHAR NULL"`
	Level                              int        `json:"level"                              column:"groups.level"                                   sql:"BIGINT"`
	IsDisposable                       bool       `json:"isDisposable"                       column:"groups.is_disposable"                           sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AncestorID                         int        `json:"ancestorId"                         column:"groups.ancestor_id"                             sql:"BIGINT"`
	ParentID                           *int       `json:"parentId"                           column:"groups.parent_id"                               sql:"BIGINT"`
	ParentQuotaSourceID                int        `json:"parentQuotaSourceId"                column:"groups.parent_quota_source_id"                  sql:"BIGINT"`
	BorrowingRoleManagedByParent       bool       `json:"borrowingRoleManagedByParent"       column:"groups.borrowing_role_managed_by_parent"        sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	TransferRoleManagedByParent        bool       `json:"transferRoleManagedByParent"        column:"groups.transfer_role_managed_by_parent"         sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	DisposalRoleManagedByParent        bool       `json:"disposalRoleManagedByParent"        column:"groups.disposal_role_managed_by_parent"         sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	DamagedRoleManagedByParent         bool       `json:"damagedRoleManagedByParent"         column:"groups.damaged_role_managed_by_parent"          sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	MissingRoleManagedByParent         bool       `json:"missingRoleManagedByParent"         column:"groups.missing_role_managed_by_parent"          sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	RepairTicketingRoleManagedByParent bool       `json:"repairTicketingRoleManagedByParent" column:"groups.repair_ticketing_role_managed_by_parent" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	TrackingRoleManagedByParent        bool       `json:"trackingRoleManagedByParent"        column:"groups.tracking_role_managed_by_parent"         sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	AuditRoleManagedByParent           bool       `json:"auditRoleManagedByParent"           column:"groups.audit_role_managed_by_parent"            sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	MaintenanceRoleManagedByParent     bool       `json:"maintenanceRoleManagedByParent"     column:"groups.maintenance_role_managed_by_parent"      sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	AssetControlRoleManagedByParent    bool       `json:"assetControlRoleManagedByParent"    column:"groups.asset_control_role_managed_by_parent"    sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsDeleted                          bool       `json:"isDeleted"                          column:"groups.is_deleted"                              sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt                          *time.Time `json:"deletedAt"                          column:"groups.deleted_at"                              sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt                          time.Time  `json:"createdAt"                          column:"groups.created_at"                              sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                          time.Time  `json:"updatedAt"                          column:"groups.updated_at"                              sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Group) TableName() string {
	return "groups"
}

func (m *Group) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("idx_groups_name", "name", false, "")
	builder.AddIndex("groups_code_uindex", "code", true, "code IS NOT NULL")

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.GroupPostQuery)
}

type GroupQuota struct {
	ID                   int            `json:"id" column:"id" sql:"BIGINT"`
	GroupID              int            `json:"groupId" column:"group_id" sql:"BIGINT"`
	Type                 GroupQuotaType `json:"type" column:"type" sql:"group_quota_type_enum"`
	Total                int            `json:"total" column:"total" sql:"INTEGER DEFAULT 0"`
	Used                 int            `json:"used" column:"used" sql:"INTEGER DEFAULT 0"`
	UsedByThisGroup      int            `json:"usedByThisGroup" column:"used_by_this_group" sql:"INTEGER DEFAULT 0"`
	AllocatedToSubgroups int            `json:"allocatedToSubgroups" column:"allocated_to_sub_groups" sql:"INTEGER DEFAULT 0"`
	Available            int            `json:"available" column:"available" sql:"INTEGER DEFAULT 0"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *GroupQuota) TableName() string {
	return "group_quotas"
}

func (m *GroupQuota) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("group_quotas_primary_key", "group_id,type", true, "")
}

type GroupApprovalType struct {
	ID           int                  `json:"id" column:"id" sql:"BIGINT"`
	GroupID      int                  `json:"groupId" column:"group_id" sql:"BIGINT"`
	Level        int                  `json:"level" column:"level" sql:"INTEGER CHECK (level > 0) DEFAULT 1"`
	Transaction  GroupTransactionType `json:"transaction" column:"transaction" sql:"group_transaction_type_enum"`
	ApprovalType ApprovalType         `json:"approvalType" column:"approval_type" sql:"approval_type_enum DEFAULT 'Or'"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *GroupApprovalType) TableName() string {
	return "group_approval_types"
}

func (m *GroupApprovalType) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("group_approval_types_primary_key", "group_id,transaction", true, "")
}

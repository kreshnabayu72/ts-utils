package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditTaskAssets struct {
	ID           int                     `json:"_id"          column:"audit_task_assets.id"             sql:"BIGINT PRIMARY KEY"`
	Asset        int                     `json:"asset"        column:"audit_task_assets.asset"          sql:"BIGINT"`
	AuditTaskId  int                     `json:"auditTaskId"  column:"audit_task_assets.audit_task_id"  sql:"BIGINT NOT NULL"`
	FoundById    *int                    `json:"foundById"    column:"audit_task_assets.found_by_id"    sql:"BIGINT NULL"`
	Status       AuditTaskAssetStatus    `json:"status"       column:"audit_task_assets.status"         sql:"VARCHAR"`
	Condition    AuditTaskAssetCondition `json:"condition"    column:"audit_task_assets.condition"      sql:"VARCHAR DEFAULT NULL"`
	Notes        string                  `json:"notes"        column:"audit_task_assets.notes"          sql:"VARCHAR NOT NULL DEFAULT ''"`
	FoundGroupId *int                    `json:"foundGroupId" column:"audit_task_assets.found_group_id" sql:"BIGINT NULL"`
	Address      string                  `json:"address"         column:"audit_task_assets.address"           sql:"VARCHAR NULL DEFAULT NULL"`
	Coordinate   string                  `json:"coordinate"      column:"audit_task_assets.coordinate"        sql:"VARCHAR NULL DEFAULT NULL"`
	CreatedAt    time.Time               `json:"createdAt"    column:"audit_task_assets.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt    time.Time               `json:"updatedAt"    column:"audit_task_assets.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditTaskAssets) TableName() string {
	return "audit_task_assets"
}

func (b *AuditTaskAssets) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_tasks_assets_assets_id",
		"asset",
		"assets",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_audit_tasks_assets_task_id",
		"audit_task_id",
		"audit_tasks",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_audit_tasks_assets_found_by_id",
		"found_by_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_audit_tasks_assets_found_group_id",
		"found_group_id",
		"groups",
		"id",
		"",
	)
}

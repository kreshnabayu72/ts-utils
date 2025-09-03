package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditTaskAssets struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset       int  `json:"asset" column:"asset" sql:"BIGINT"`
	AuditTaskId int  `json:"auditTaskId" column:"audit_task_id" sql:"BIGINT NOT NULL"`
	FoundById   *int `json:"foundById" column:"found_by_id" sql:"BIGINT NULL"`

	Status    AuditTaskAssetStatus    `json:"status" column:"status" sql:"VARCHAR"`
	Condition AuditTaskAssetCondition `json:"condition" column:"condition" sql:"VARCHAR DEFAULT NULL"`
	Notes     string                  `json:"notes" column:"notes" sql:"VARCHAR NOT NULL DEFAULT ''"`

	FoundGroupId *int `json:"foundGroupId" column:"found_group_id" sql:"BIGINT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
		"audit_found_by_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_audit_tasks_assets_found_group_id",
		"audit_found_group_id",
		"groups",
		"id",
		"",
	)
}

package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditOtherTransactionAsset struct {
	ID          int                   `json:"_id"         column:"audit_other_transaction_assets.id"            sql:"BIGINT PRIMARY KEY"`
	Asset       int                   `json:"asset"       column:"audit_other_transaction_assets.asset"         sql:"BIGINT"`
	AuditTaskId int                   `json:"auditTaskId" column:"audit_other_transaction_assets.audit_task_id" sql:"BIGINT NOT NULL"`
	Type        AuditOtherTransaction `json:"type"        column:"audit_other_transaction_assets.type"          sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	CreatedAt   time.Time             `json:"createdAt"   column:"audit_other_transaction_assets.created_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time             `json:"updatedAt"   column:"audit_other_transaction_assets.updated_at"    sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditOtherTransactionAsset) TableName() string {
	return "audit_other_transaction_assets"
}

func (b *AuditOtherTransactionAsset) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_other_transaction_assets_asset_id",
		"asset",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_audit_other_transaction_assets_audit_task_id",
		"audit_task_id",
		"audit_tasks",
		"id",
		"",
	)
}

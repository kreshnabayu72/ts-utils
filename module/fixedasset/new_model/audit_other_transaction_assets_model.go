package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditOtherTransactionAsset struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset       int                   `json:"asset" column:"asset" sql:"BIGINT"`
	AudittaskId int                   `json:"auditTaskId" column:"audit_task_id" sql:"BIGINT NOT NULL"`
	Type        AuditOtherTransaction `json:"type" column:"type" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

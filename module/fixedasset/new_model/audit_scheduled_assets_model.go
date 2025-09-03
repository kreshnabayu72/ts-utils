package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditScheduledAsset struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset           int  `json:"asset" column:"asset" sql:"BIGINT"`
	AuditScheduleId int  `json:"auditScheduleId" column:"audit_schedule_id" sql:"BIGINT NOT NULL"`
	IsActive        bool `json:"isActive" column:"is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditScheduledAsset) TableName() string {
	return "audit_scheduled_assets"
}

func (b *AuditScheduledAsset) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_scheduled_assets_schedule_id",
		"audit_schedule_id",
		"audit_schedules",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_audit_schedules_categories_category_id",
		"category_id",
		"categories",
		"id",
		"CASCADE",
	)
}

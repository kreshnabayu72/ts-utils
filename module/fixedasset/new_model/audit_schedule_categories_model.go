package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditScheduleCategories struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	CategoryID      int `json:"categoryId" column:"category_id" sql:"BIGINT"`
	AuditScheduleId int `json:"auditScheduleId" column:"audit_schedule_id" sql:"BIGINT NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditScheduleCategories) TableName() string {
	return "audit_schedule_categories"
}

func (b *AuditScheduleCategories) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_schedules_categories_schedule_id",
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

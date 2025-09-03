package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditTaskCategories struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	CategoryID  int `json:"categoryId" column:"category_id" sql:"BIGINT"`
	AuditTaskId int `json:"auditTaskId" column:"audit_task_id" sql:"BIGINT NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditTaskCategories) TableName() string {
	return "audit_task_categories"
}

func (b *AuditTaskCategories) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_tasks_categories_task_id",
		"audit_task_id",
		"audit_tasks",
		"id",
		"CASCADE",
	)
	builder.AddForeignKey(
		"fk_audit_tasks_categories_category_id",
		"category_id",
		"categories",
		"id",
		"CASCADE",
	)
}

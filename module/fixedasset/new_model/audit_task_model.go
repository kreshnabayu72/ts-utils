package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditTask struct {
	ID            int             `json:"_id"            column:"audit_tasks.id"              sql:"BIGINT PRIMARY KEY"`
	ScheduleID    int             `json:"scheduleId"     column:"audit_tasks.schedule_id"     sql:"BIGINT"`
	TransactionId string          `json:"transaction_id" column:"audit_tasks.transaction_id"  sql:"VARCHAR"`
	TagType       TagType         `json:"tagType"        column:"audit_tasks.tag_type"        sql:"tag_type_enum"`
	AuditType     AuditType       `json:"auditType"      column:"audit_tasks.audit_type"      sql:"audit_type_enum"`
	Status        AuditTaskStatus `json:"status"         column:"audit_tasks.status"          sql:"audit_task_status_enum"`
	Duration      int             `json:"duration"       column:"audit_tasks.duration"        sql:"INTEGER NOT NULL"`
	ModifiedById  int             `json:"modifiedById"   column:"audit_tasks.modified_by_id"  sql:"BIGINT"`
	ApprovalRound int             `json:"approvalRound"  column:"audit_tasks.approval_round"  sql:"INTEGER DEFAULT 1"`
	StartedAt     *time.Time      `json:"startedAt"      column:"audit_tasks.started_at"      sql:"TIMESTAMPTZ NULL"`
	GracePeriodAt *time.Time      `json:"gracePeriodAt"  column:"audit_tasks.grace_period_at" sql:"TIMESTAMPTZ NULL"`
	FinishedAt    *time.Time      `json:"finishedAt"     column:"audit_tasks.finished_at"     sql:"TIMESTAMPTZ NULL"`
	CreatedAt     time.Time       `json:"createdAt"      column:"audit_tasks.created_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time       `json:"updatedAt"      column:"audit_tasks.updated_at"      sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditTask) TableName() string {
	return "audit_tasks"
}

func (b *AuditTask) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_tasks_modified_by",
		"modified_by_id",
		"users",
		"id",
		"SET NULL",
	)

	builder.AddForeignKey(
		"fk_audit_tasks_schedule",
		"schedule_id",
		"audit_schedules",
		"id",
		"",
	)
}

package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditTask struct {
	ID            int    `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	ScheduleID    int    `json:"scheduleId" column:"schedule_id" sql:"BIGINT"`
	TransactionId string `json:"transaction_id" column:"transaction_id" sql:"VARCHAR"`

	TagType   TagType         `json:"tagType" column:"tag_type" sql:"tag_type_enum"`
	AuditType AuditType       `json:"auditType" column:"audit_type" sql:"audit_type_enum"`
	Status    AuditTaskStatus `json:"status" column:"status" sql:"audit_task_status_enum"`

	Duration      int `json:"duration" column:"duration" sql:"INTEGER NOT NULL"`
	ModifiedById  int `json:"modifiedById" column:"modified_by_id" sql:"BIGINT NOT NULL"`
	ApprovalRound int `json:"approvalRound" column:"approval_round" sql:"INTEGER DEFAULT 1"`

	StartedAt     *time.Time `json:"startedAt" column:"started_at" sql:"TIMESTAMPTZ NULL"`
	GracePeriodAt *time.Time `json:"gracePeriodAt" column:"grace_period_at" sql:"TIMESTAMPTZ NULL"`
	FinishedAt    *time.Time `json:"finishedAt" column:"finished_at" sql:"TIMESTAMPTZ NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

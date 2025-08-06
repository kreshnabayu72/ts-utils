package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineTask struct {
	ID            int               `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	ScheduleID    int               `json:"scheduleId" column:"schedule_id" sql:"BIGINT NOT NULL"`
	AssetID       int               `json:"assetId" column:"asset_id" sql:"BIGINT NOT NULL"`
	Status        RoutineTaskStatus `json:"status" column:"status" sql:"routine_task_status_enum NOT NULL"`
	Duration      int               `json:"duration" column:"duration" sql:"INT NOT NULL"`
	ModifiedByID  int               `json:"modifiedById" column:"modified_by_id" sql:"BIGINT NOT NULL"`
	ApprovalRound int               `json:"approvalRound" column:"approval_round" sql:"INT NOT NULL"`
	StartAt       time.Time         `json:"startAt" column:"start_at" sql:"TIMESTAMP NOT NULL"`
	CreatedAt     time.Time         `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time         `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *RoutineTask) TableName() string {
	return "routine_tasks"
}

func (m *RoutineTask) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddEnum(
		"routine_task_status_enum",
		[]string{
			string(RoutineTaskStatusOnSchedule),
			string(RoutineTaskStatusWaitingForAssignment),
			string(RoutineTaskStatusUnderMaintenance),
			string(RoutineTaskStatusWaitingForReview),
			string(RoutineTaskStatusWaitingForApproval),
			string(RoutineTaskStatusOverdue),
			string(RoutineTaskStatusCompleted),
			string(RoutineTaskStatusCanceled),
			string(RoutineTaskStatusReportedDamaged),
			string(RoutineTaskStatusStopped),
		},
	)

	builder.AddForeignKey(
		"fk_routine_task_schedule_id",
		"schedule_id",
		"routine_schedules",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_task_asset_id",
		"asset_id",
		"assets",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_task_modified_by_id",
		"modified_by_id",
		"users",
		"id",
		"",
	)
}

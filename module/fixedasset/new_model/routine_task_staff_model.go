package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineTaskStaff struct {
	ID            int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	UserID        int       `json:"userId" column:"user_id" sql:"BIGINT NOT NULL"`
	RoutineTaskID int       `json:"routineTaskId" column:"routine_task_id" sql:"BIGINT NOT NULL"`
	CreatedAt     time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *RoutineTaskStaff) TableName() string {
	return "routine_task_staffs"
}

func (m *RoutineTaskStaff) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_routine_schedule_staff_user_id",
		"user_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_schedule_staff_routine_task_id",
		"routine_task_id",
		"routine_tasks",
		"id",
		"",
	)
}

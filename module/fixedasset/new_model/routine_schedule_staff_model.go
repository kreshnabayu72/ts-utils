package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineScheduleStaff struct {
	ID                int       `json:"_id"               column:"routine_schedule_staffs.id"                  sql:"BIGINT PRIMARY KEY"`
	UserID            int       `json:"userId"            column:"routine_schedule_staffs.user_id"             sql:"BIGINT NOT NULL"`
	RoutineScheduleID int       `json:"routineScheduleId" column:"routine_schedule_staffs.routine_schedule_id" sql:"BIGINT NOT NULL"`
	CreatedAt         time.Time `json:"createdAt"         column:"routine_schedule_staffs.created_at"          sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt         time.Time `json:"updatedAt"         column:"routine_schedule_staffs.updated_at"          sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RoutineScheduleStaff) TableName() string {
	return "routine_schedule_staffs"
}

func (m *RoutineScheduleStaff) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_routine_schedule_staff_user_id",
		"user_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_schedule_staff_routine_schedule_id",
		"routine_schedule_id",
		"routine_schedules",
		"id",
		"",
	)
}

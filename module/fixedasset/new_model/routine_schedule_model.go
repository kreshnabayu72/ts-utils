package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineSchedule struct {
	ID                   int                      `json:"_id"                  column:"routine_schedules.id"                     sql:"BIGINT PRIMARY KEY"`
	Name                 string                   `json:"name"                 column:"routine_schedules.name"                   sql:"VARCHAR NOT NULL"`
	AssetNameID          int                      `json:"assetNameId"          column:"routine_schedules.asset_name_id"          sql:"BIGINT NOT NULL"`
	GroupID              int                      `json:"groupId"              column:"routine_schedules.group_id"               sql:"BIGINT NOT NULL"`
	FrequencyLabel       RoutineScheduleFrequency `json:"frequencyLabel"       column:"routine_schedules.frequency_label"        sql:"frequency_enum NOT NULL"`
	FrequencyCustomValue *int                     `json:"frequencyCustomValue" column:"routine_schedules.frequency_custom_value" sql:"INT DEFAULT NULL"`
	Duration             int                      `json:"duration"             column:"routine_schedules.duration"               sql:"INT NOT NULL"`
	ModifiedByID         int                      `json:"modifiedById"         column:"routine_schedules.modified_by_id"         sql:"BIGINT NOT NULL"`
	IsDeleted            bool                     `json:"isDeleted"            column:"routine_schedules.is_deleted"             sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt            *time.Time               `json:"deletedAt"            column:"routine_schedules.deleted_at"             sql:"TIMESTAMPTZ DEFAULT NULL"`
	NextSchedule         time.Time                `json:"nextSchedule"         column:"routine_schedules.next_schedule"          sql:"TIMESTAMPTZ NOT NULL"`
	CreatedAt            time.Time                `json:"createdAt"            column:"routine_schedules.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time                `json:"updatedAt"            column:"routine_schedules.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RoutineSchedule) TableName() string {
	return "routine_schedules"
}

func (m *RoutineSchedule) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddEnum(
		"frequency_enum",
		[]string{
			string(RoutineScheduleFrequencyOnce),
			string(RoutineScheduleFrequencyDaily),
			string(RoutineScheduleFrequencyWeekly),
			string(RoutineScheduleFrequencyMonthly),
			string(RoutineScheduleFrequencyAnnually),
			string(RoutineScheduleFrequencyCustom),
		},
	)

	builder.AddForeignKey(
		"fk_routine_schedule_asset_name_id",
		"asset_name_id",
		"asset_names",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_schedule_group_id",
		"group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_routine_schedule_modified_by_id",
		"modified_by_id",
		"users",
		"id",
		"",
	)
}

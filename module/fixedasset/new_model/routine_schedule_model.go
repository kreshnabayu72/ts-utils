package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineSchedule struct {
	ID                   int                      `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name                 string                   `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	AssetNameID          int                      `json:"assetNameId" column:"asset_name_id" sql:"BIGINT NOT NULL"`
	GroupID              int                      `json:"groupId" column:"group_id" sql:"BIGINT NOT NULL"`
	FrequencyLabel       RoutineScheduleFrequency `json:"frequencyLabel" column:"frequency_label" sql:"frequency_enum NOT NULL"`
	FrequencyCustomValue *int                     `json:"frequencyCustomValue" column:"frequency_custom_value" sql:"INT DEFAULT NULL"`
	Duration             int                      `json:"duration" column:"duration" sql:"INT NOT NULL"`
	ModifiedByID         int                      `json:"modifiedById" column:"modified_by_id" sql:"BIGINT NOT NULL"`
	IsDeleted            bool                     `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt            *time.Time               `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMP DEFAULT NULL"`
	NextSchedule         time.Time                `json:"nextSchedule" column:"next_schedule" sql:"TIMESTAMP NOT NULL"`
	CreatedAt            time.Time                `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time                `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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

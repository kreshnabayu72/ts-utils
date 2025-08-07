package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineTaskIndicator struct {
	ID                  int               `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TaskID              int               `json:"taskId" column:"task_id" sql:"BIGINT NOT NULL"`
	Name                string            `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	Measurement         string            `json:"measurement" column:"measurement" sql:"VARCHAR NOT NULL"`
	IndicatorType       IndicatorType     `json:"indicatorType" column:"indicator_type" sql:"indicator_type_enum NOT NULL"`
	DataType            IndicatorDataType `json:"dataType" column:"data_type" sql:"data_type_enum NOT NULL"`
	Standard            string            `json:"standard" column:"standard" sql:"VARCHAR NOT NULL"`
	IsMandatory         bool              `json:"isMandatory" column:"is_mandatory" sql:"BOOLEAN NOT NULL"`
	IsHaveStandardValue bool              `json:"isHaveStandardValue" column:"is_have_standard_value" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Options             []string          `json:"options" column:"options" sql:"TEXT[]"`
	CreatedAt           time.Time         `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt           time.Time         `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *RoutineTaskIndicator) TableName() string {
	return "routine_task_indicators"
}

func (m *RoutineTaskIndicator) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddEnum(
		"indicator_type_enum",
		[]string{
			string(IndicatorTypeSelection),
			string(IndicatorTypeSingleValue),
			string(IndicatorTypeRangeValue),
		},
	)
	builder.AddEnum(
		"data_type_enum",
		[]string{
			string(IndicatorDataTypeText),
			string(IndicatorDataTypeNumber),
		},
	)

	builder.AddForeignKey(
		"fk_routine_task_indicator_task_id",
		"task_id",
		"routine_tasks",
		"id",
		"",
	)
}

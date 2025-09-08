package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type RoutineTaskIndicator struct {
	ID                  int               `json:"_id"                 column:"routine_task_indicators.id"                     sql:"BIGINT PRIMARY KEY"`
	TaskID              int               `json:"taskId"              column:"routine_task_indicators.task_id"                sql:"BIGINT NOT NULL"`
	Name                string            `json:"name"                column:"routine_task_indicators.name"                   sql:"VARCHAR NOT NULL"`
	Measurement         string            `json:"measurement"         column:"routine_task_indicators.measurement"            sql:"VARCHAR NOT NULL"`
	IndicatorType       IndicatorType     `json:"indicatorType"       column:"routine_task_indicators.indicator_type"         sql:"indicator_type_enum NOT NULL"`
	DataType            IndicatorDataType `json:"dataType"            column:"routine_task_indicators.data_type"              sql:"data_type_enum NOT NULL"`
	Standard            string            `json:"standard"            column:"routine_task_indicators.standard"               sql:"VARCHAR NOT NULL"`
	IsMandatory         bool              `json:"isMandatory"         column:"routine_task_indicators.is_mandatory"           sql:"BOOLEAN NOT NULL"`
	IsHaveStandardValue bool              `json:"isHaveStandardValue" column:"routine_task_indicators.is_have_standard_value" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Options             []string          `json:"options"             column:"routine_task_indicators.options"                sql:"TEXT[]"`
	CreatedAt           time.Time         `json:"createdAt"           column:"routine_task_indicators.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt           time.Time         `json:"updatedAt"           column:"routine_task_indicators.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

package new_model

import (
	"time"
)

type RoutineScheduleIndicator struct {
	ID                  int               `json:"_id"                 column:"routine_schedule_indicators.id"`
	ScheduleID          int               `json:"scheduleId"          column:"routine_schedule_indicators.schedule_id"`
	Name                string            `json:"name"                column:"routine_schedule_indicators.name"`
	Measurement         string            `json:"measurement"         column:"routine_schedule_indicators.measurement"`
	IndicatorType       IndicatorType     `json:"indicatorType"       column:"routine_schedule_indicators.indicator_type"`
	DataType            IndicatorDataType `json:"dataType"            column:"routine_schedule_indicators.data_type"`
	Standard            string            `json:"standard"            column:"routine_schedule_indicators.standard"`
	IsMandatory         bool              `json:"isMandatory"         column:"routine_schedule_indicators.is_mandatory"`
	IsHaveStandardValue bool              `json:"isHaveStandardValue" column:"routine_schedule_indicators.is_have_standard_value"`
	Options             []string          `json:"options"             column:"routine_schedule_indicators.options"`
	CreatedAt           time.Time         `json:"createdAt"           column:"routine_schedule_indicators.created_at"`
	UpdatedAt           time.Time         `json:"updatedAt"           column:"routine_schedule_indicators.updated_at"`
}

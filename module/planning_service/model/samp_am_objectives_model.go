package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/new_model"
)

type SampAmObjective struct {
	ID     int `json:"id" column:"samp_am_objectives.id"`
	SampID int `json:"sampId" column:"samp_am_objectives.samp_id"` // Foreign Key to samp_current_versions.id

	Objectives        string              `json:"objectives" column:"samp_am_objectives.objectives"`
	KpiOrTarget       string              `json:"kpiOrTarget" column:"samp_am_objectives.kpi_or_target"`
	MeasurementMethod string              `json:"measurementMethod" column:"samp_am_objectives.measurement_method"`
	Timeline          new_model.TimeRange `json:"timeline" column:"samp_am_objectives.timeline"`

	CreatedAt time.Time `json:"createdAt" column:"samp_am_objectives.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"samp_am_objectives.updated_at"`
}

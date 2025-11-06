package new_model

import (
	"time"
)

type AmpDemandAndPerformanceRequirements struct {
	ID    int `json:"id" column:"amp_demand_and_performance_requirements.id"`
	AmpID int `json:"sampId" column:"amp_demand_and_performance_requirements.amp_id"` // Foreign Key to amp

	Description       string `json:"description" column:"amp_demand_and_performance_requirements.description"`
	PerformanceTarget string `json:"performanceTarget" column:"amp_demand_and_performance_requirements.performance_target"`
	MeasurementMethod string `json:"measurementMethod" column:"amp_demand_and_performance_requirements.measurement_method"`
	Notes             string `json:"notes" column:"amp_demand_and_performance_requirements.notes"`

	IsDeleted bool       `json:"isDeleted"        column:"amp_demand_and_performance_requirements.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"amp_demand_and_performance_requirements.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"amp_demand_and_performance_requirements.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_demand_and_performance_requirements.updated_at"`
}

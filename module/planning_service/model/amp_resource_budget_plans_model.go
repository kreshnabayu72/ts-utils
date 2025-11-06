package new_model

import (
	"time"
)

type AmpResourceBudgetPlans struct {
	ID    int `json:"id" column:"amp_lifecycle_management_plans.id"`
	AmpID int `json:"sampId" column:"amp_lifecycle_management_plans.amp_id"` // Foreign Key to amp

	Resource              string  `json:"resource" column:"amp_lifecycle_management_plans.resource"`
	Quantity              string  `json:"quantity" column:"amp_lifecycle_management_plans.quantity"`
	EstimatedCostValue    float64 `json:"estimatedCostValue" column:"amp_lifecycle_management_plans.estimated_cost_value"`
	EstimatedCostCurrency string  `json:"estimatedCostCurrency" column:"amp_lifecycle_management_plans.estimated_cost_currency"`
	Pic                   string  `json:"pic" column:"amp_lifecycle_management_plans.pic"`
	Notes                 string  `json:"notes" column:"amp_lifecycle_management_plans.notes"`

	IsDeleted bool       `json:"isDeleted"        column:"amp_lifecycle_management_plans.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"amp_lifecycle_management_plans.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"amp_lifecycle_management_plans.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_lifecycle_management_plans.updated_at"`
}

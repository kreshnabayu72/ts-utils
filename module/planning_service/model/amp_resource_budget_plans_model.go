package new_model

import (
	"time"
)

type AmpResourceBudgetPlans struct {
	ID           int `json:"id" column:"amp_resource_budget_plans.id"`
	AmpID        int `json:"ampId" column:"amp_resource_budget_plans.amp_id"`                // Foreign Key to amp
	AmpHistoryID int `json:"ampHistoryId" column:"amp_resource_budget_plans.amp_history_id"` // Foreign Key to amp history

	Resource              string  `json:"resource" column:"amp_resource_budget_plans.resource"`
	Quantity              string  `json:"quantity" column:"amp_resource_budget_plans.quantity"`
	EstimatedCostValue    float64 `json:"estimatedCostValue" column:"amp_resource_budget_plans.estimated_cost_value"`
	EstimatedCostCurrency string  `json:"estimatedCostCurrency" column:"amp_resource_budget_plans.estimated_cost_currency"`
	Pic                   string  `json:"pic" column:"amp_resource_budget_plans.pic"`
	Notes                 string  `json:"notes" column:"amp_resource_budget_plans.notes"`

	IsDeleted bool       `json:"isDeleted"        column:"amp_resource_budget_plans.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"amp_resource_budget_plans.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"amp_resource_budget_plans.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_resource_budget_plans.updated_at"`
}

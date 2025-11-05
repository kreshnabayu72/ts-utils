package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/new_model"
)

type AmpLifecycleManagementPlans struct {
	ID    int `json:"id" column:"amp_lifecycle_management_plans.id"`
	AmpID int `json:"sampId" column:"amp_lifecycle_management_plans.amp_id"` // Foreign Key to amp

	Stage    string              `json:"stage" column:"amp_lifecycle_management_plans.stage"`
	Activity string              `json:"activity" column:"amp_lifecycle_management_plans.activity"`
	Timeline new_model.TimeRange `json:"timeline" column:"amp_lifecycle_management_plans.timeline"`
	Pic      string              `json:"pic" column:"amp_lifecycle_management_plans.pic"`
	Notes    string              `json:"notes" column:"amp_lifecycle_management_plans.notes"`

	IsDeleted bool       `json:"isDeleted"        column:"amp_lifecycle_management_plans.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"amp_lifecycle_management_plans.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"amp_lifecycle_management_plans.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"amp_lifecycle_management_plans.updated_at"`
}

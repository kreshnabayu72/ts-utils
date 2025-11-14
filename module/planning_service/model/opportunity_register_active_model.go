package new_model

import (
	"time"
)

type OpportunityRegisterActive struct {
	ID int `json:"id" column:"opportunity_registers_active.id"`

	OpportunityID string    `json:"opportunityId" column:"opportunity_registers_active.opportunity_id"`
	Opportunity   string    `json:"opportunity" column:"opportunity_registers_active.opportunity"`
	Benefit       string    `json:"benefit" column:"opportunity_registers_active.benefit"`
	Action        string    `json:"action" column:"opportunity_registers_active.action"`
	TargetDate    time.Time `json:"targetDate" column:"opportunity_registers_active.target_date"`
	Status        string    `json:"status" column:"opportunity_registers_active.status"`
	PicID         int       `json:"picId" column:"opportunity_registers_active.pic_id"`

	CreatedAt time.Time `json:"createdAt" column:"opportunity_registers_active.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"opportunity_registers_active.updated_at"`

	IsDeleted bool       `json:"isDeleted" column:"opportunity_registers_active.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"opportunity_registers_active.deleted_at"`
}

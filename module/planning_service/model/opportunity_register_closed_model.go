package new_model

import (
	"time"
)

type OpportunityRegisterClosed struct {
	ID int `json:"id" column:"opportunity_registers_closed.id"`

	OpportunityID string    `json:"opportunityId" column:"opportunity_registers_closed.opportunity_id"`
	Opportunity   string    `json:"opportunity" column:"opportunity_registers_closed.opportunity"`
	Benefit       string    `json:"benefit" column:"opportunity_registers_closed.benefit"`
	Action        string    `json:"action" column:"opportunity_registers_closed.action"`
	TargetDate    time.Time `json:"targetDate" column:"opportunity_registers_closed.target_date"`
	Status        string    `json:"status" column:"opportunity_registers_closed.status"`
	PicName       string    `json:"picName" column:"opportunity_registers_closed.pic_name"`

	CreatedAt time.Time `json:"createdAt" column:"opportunity_registers_closed.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"opportunity_registers_closed.updated_at"`
}

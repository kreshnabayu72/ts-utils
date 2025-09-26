package new_model

import (
	"time"
)

// DisposalPrelist represents the disposal_prelists table in the database.
type DisposalPrelist struct {
	ID              int                   `json:"id"              column:"disposal_prelists.id"`
	OriginGroupID   int                   `json:"originGroupId"   column:"disposal_prelists.origin_group_id"`
	DisposalGroupID int                   `json:"disposalGroupId" column:"disposal_prelists.disposal_group_id"`
	Method          DisposalMethod        `json:"method"          column:"disposal_prelists.method"`
	CustomMethod    string                `json:"customMethod"    column:"disposal_prelists.custom_method"`
	LastCondition   DisposalLastCondition `json:"lastCondition"   column:"disposal_prelists.last_condition"`
	Organization    string                `json:"organization"    column:"disposal_prelists.organization"`
	CreatedAt       time.Time             `json:"createdAt"       column:"disposal_prelists.created_at"`
	UpdatedAt       time.Time             `json:"updatedAt"       column:"disposal_prelists.updated_at"`
}

package new_model

import "time"

type DisposalPrelist struct {
	ID              int            `json:"id" column:"id"`
	DisposalGroupId int            `json:"disposalGroupId" column:"disposal_group_id"`
	Method          DisposalMethod `json:"method" column:"method"`
	Organization    string         `json:"organization" column:"organization"`
	CreatedAt       time.Time      `json:"createdAt" column:"created_at"`
	UpdatedAt       time.Time      `json:"updatedAt" column:"updated_at"`
}

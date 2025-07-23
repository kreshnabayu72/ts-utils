package new_model

import "time"

type TransferPrelist struct {
	ID                 int       `json:"id" column:"id"`
	OriginGroupId      int       `json:"originGroupId" column:"origin_group_id"`
	DestinationGroupId int       `json:"destinationGroupId" column:"destination_group_id"`
	CreatedAt          time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt          time.Time `json:"updatedAt" column:"updated_at"`
}

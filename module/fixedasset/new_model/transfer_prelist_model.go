package new_model

import (
	"time"
)

// TransferPrelist represents the transfer_prelists table in the database.
type TransferPrelist struct {
	ID                 int       `json:"id"                 column:"transfer_prelists.id"`
	OriginGroupID      int       `json:"originGroupId"      column:"transfer_prelists.origin_group_id"`
	DestinationGroupID int       `json:"destinationGroupId" column:"transfer_prelists.destination_group_id"`
	CreatedAt          time.Time `json:"createdAt"          column:"transfer_prelists.created_at"`
	UpdatedAt          time.Time `json:"updatedAt"          column:"transfer_prelists.updated_at"`
}

package new_model

import (
	"time"
)

// AssignmentRequestHistory represents the assignment_request_histories table in the database.
type AssignmentRequestHistory struct {
	ID               int                     `json:"id"               column:"assignment_request_histories.id"`
	ParentID         *int                    `json:"parentId"         column:"assignment_request_histories.parent_id"`
	TagType          TagType                 `json:"tagType"          column:"assignment_request_histories.tag_type"`
	Asset            int                     `json:"asset"            column:"assignment_request_histories.asset"`
	AssetID          string                  `json:"assetId"          column:"assignment_request_histories.asset_id"`
	AliasCode        *string                 `json:"aliasCode"                column:"assignment_request_histories.alias_code"`
	Name             string                  `json:"name"             column:"assignment_request_histories.name"`
	Brand            string                  `json:"brand"            column:"assignment_request_histories.brand"`
	Model            string                  `json:"model"            column:"assignment_request_histories.model"`
	NameWithSequence string                  `json:"nameWithSequence" column:"assignment_request_histories.name_with_sequence"`
	SequenceNumber   int                     `json:"sequenceNumber"   column:"assignment_request_histories.sequence_number"`
	Image            *string                 `json:"image"            column:"assignment_request_histories.image"`
	Group            string                  `json:"group"            column:"assignment_request_histories.group"`
	GroupID          int                     `json:"groupId"          column:"assignment_request_histories.group_id"`
	Transaction      *int                    `json:"transaction"      column:"assignment_request_histories.transaction"`
	Status           AssignmentRequestStatus `json:"status"           column:"assignment_request_histories.status"`
	Notes            string                  `json:"notes"            column:"assignment_request_histories.notes"`
	TransactionID    string                  `json:"transactionId"              column:"assignment_request_histories.transaction_id"`
	AssignedToId     int                     `json:"assignedToId"   column:"assignment_request_histories.assigned_to_id"`
	AssignedTo       string                  `json:"assignedTo"       column:"assignment_request_histories.assigned_to"`
	Manager          string                  `json:"manager"          column:"assignment_request_histories.manager"`
	CreatedAt        time.Time               `json:"createdAt"        column:"assignment_request_histories.created_at"`
	UpdatedAt        time.Time               `json:"updatedAt"        column:"assignment_request_histories.updated_at"`
}

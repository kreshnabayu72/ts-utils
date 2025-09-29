package new_model

import (
	"time"
)

type BorrowingRequestHistory struct {
	ID               int                    `json:"id"               column:"borrowing_request_histories.id"`
	ParentID         *int                   `json:"parentId"         column:"borrowing_request_histories.parent_id"`
	TagType          TagType                `json:"tagType"          column:"borrowing_request_histories.tag_type"`
	Asset            int                    `json:"asset"            column:"borrowing_request_histories.asset"`
	AssetID          string                 `json:"assetId"          column:"borrowing_request_histories.asset_id"`
	AliasCode        *string                `json:"aliasCode"                column:"borrowing_request_histories.alias_code"`
	Name             string                 `json:"name"             column:"borrowing_request_histories.name"`
	Brand            string                 `json:"brand"            column:"borrowing_request_histories.brand"`
	Model            string                 `json:"model"            column:"borrowing_request_histories.model"`
	NameWithSequence string                 `json:"nameWithSequence" column:"borrowing_request_histories.name_with_sequence"`
	SequenceNumber   int                    `json:"sequenceNumber"   column:"borrowing_request_histories.sequence_number"`
	Image            *string                `json:"image"            column:"borrowing_request_histories.image"`
	Group            string                 `json:"group"            column:"borrowing_request_histories.group"`
	GroupID          int                    `json:"groupId"          column:"borrowing_request_histories.group_id"`
	Transaction      int                    `json:"transaction"      column:"borrowing_request_histories.transaction"`
	Status           BorrowingRequestStatus `json:"status"           column:"borrowing_request_histories.status"`
	Notes            string                 `json:"notes"            column:"borrowing_request_histories.notes"`
	Borrower         string                 `json:"borrower"         column:"borrowing_request_histories.borrower"`
	Manager          string                 `json:"manager"          column:"borrowing_request_histories.manager"`
	Duration         int                    `json:"duration"         column:"borrowing_request_histories.duration"`
	RequestedAt      time.Time              `json:"requestDate"      column:"borrowing_request_histories.requested_at"`
	ReturnedAt       time.Time              `json:"returnDate"       column:"borrowing_request_histories.returned_at"`
	DueAt            time.Time              `json:"dueDate"          column:"borrowing_request_histories.due_at"`
	CreatedAt        time.Time              `json:"createdAt"        column:"borrowing_request_histories.created_at"`
	UpdatedAt        time.Time              `json:"updatedAt"        column:"borrowing_request_histories.updated_at"`
}

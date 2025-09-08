package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingRequestHistory struct {
	ID               int                    `json:"id"               column:"borrowing_request_histories.id"                 sql:"BIGINT PRIMARY KEY"`
	Asset            int                    `json:"asset"            column:"borrowing_request_histories.asset"              sql:"BIGINT"`
	AssetID          string                 `json:"assetId"          column:"borrowing_request_histories.asset_id"           sql:"VARCHAR"`
	Name             string                 `json:"name"             column:"borrowing_request_histories.name"               sql:"VARCHAR"`
	Brand            string                 `json:"brand"            column:"borrowing_request_histories.brand"              sql:"VARCHAR"`
	Model            string                 `json:"model"            column:"borrowing_request_histories.model"              sql:"VARCHAR"`
	NameWithSequence string                 `json:"nameWithSequence" column:"borrowing_request_histories.name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                    `json:"sequenceNumber"   column:"borrowing_request_histories.sequence_number"    sql:"BIGINT"`
	Image            *string                `json:"image"            column:"borrowing_request_histories.image"              sql:"VARCHAR"`
	Group            string                 `json:"group"            column:"borrowing_request_histories.group"              sql:"VARCHAR"`
	GroupID          int                    `json:"groupId"          column:"borrowing_request_histories.group_id"           sql:"BIGINT"`
	Transaction      int                    `json:"transaction"      column:"borrowing_request_histories.transaction"        sql:"BIGINT"`
	Status           BorrowingRequestStatus `json:"status"           column:"borrowing_request_histories.status"             sql:"borrowing_request_status_enum"`
	Notes            string                 `json:"notes"            column:"borrowing_request_histories.notes"              sql:"VARCHAR"`
	Borrower         string                 `json:"borrower"         column:"borrowing_request_histories.borrower"           sql:"VARCHAR"`
	Manager          string                 `json:"manager"          column:"borrowing_request_histories.manager"            sql:"VARCHAR"`
	Duration         int                    `json:"duration"         column:"borrowing_request_histories.duration"           sql:"INTEGER"`
	RequestedAt      time.Time              `json:"requestDate"      column:"borrowing_request_histories.requested_at"       sql:"TIMESTAMPTZ NULL"`
	ReturnedAt       time.Time              `json:"returnDate"       column:"borrowing_request_histories.returned_at"        sql:"TIMESTAMPTZ NULL"`
	DueAt            time.Time              `json:"dueDate"          column:"borrowing_request_histories.due_at"             sql:"TIMESTAMPTZ NULL"`
	CreatedAt        time.Time              `json:"createdAt"        column:"borrowing_request_histories.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time              `json:"updatedAt"        column:"borrowing_request_histories.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingRequestHistory) TableName() string {
	return "borrowing_request_histories"
}

func (m *BorrowingRequestHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}

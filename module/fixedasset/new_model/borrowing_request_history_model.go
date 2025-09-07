package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingRequestHistory struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	// Foreign key to parent request history
	ParentID *int `json:"parentId"            column:"parent_id"           sql:"BIGINT"`

	// Foreign key to assets
	Asset int `json:"asset"            column:"asset"           sql:"BIGINT"`
	// Asset ID based on alias code
	AssetID          string  `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	Name             string  `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string  `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string  `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string  `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int     `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string `json:"image"         column:"image"        sql:"VARCHAR"`
	Group            string  `json:"group"   column:"group" sql:"VARCHAR"`
	GroupID          int     `json:"groupId"   column:"group_id" sql:"BIGINT"`

	Transaction int                    `json:"transaction" column:"transaction" sql:"BIGINT"`
	Status      BorrowingRequestStatus `json:"status" column:"status" sql:"borrowing_request_status_enum"`
	Notes       string                 `json:"notes" column:"notes" sql:"VARCHAR"`
	Borrower    string                 `json:"borrower"   column:"borrower" sql:"VARCHAR"`
	Manager     string                 `json:"manager"   column:"manager" sql:"VARCHAR"`

	Duration    int       `json:"duration" column:"duration" sql:"INTEGER"`
	RequestedAt time.Time `json:"requestDate" column:"requested_at" sql:"TIMESTAMPTZ NULL"`
	ReturnedAt  time.Time `json:"returnDate" column:"returned_at" sql:"TIMESTAMPTZ NULL"`
	DueAt       time.Time `json:"dueDate" column:"due_at" sql:"TIMESTAMPTZ NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingRequestHistory) TableName() string {
	return "borrowing_request_histories"
}

func (m *BorrowingRequestHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_request_histories_parent",
		"parent_id",
		"borrowing_requests",
		"id",
		"",
	)
}

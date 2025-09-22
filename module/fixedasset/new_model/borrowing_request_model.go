package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingRequest struct {
	ID               int                    `json:"id"               column:"borrowing_requests.id"                 sql:"BIGINT PRIMARY KEY"`
	Asset            int                    `json:"asset"            column:"borrowing_requests.asset"              sql:"BIGINT"`
	Transaction      int                    `json:"transaction"      column:"borrowing_requests.transaction"        sql:"BIGINT"`
	ApprovalRound    int                    `json:"approval_round"   column:"borrowing_requests.approval_round"     sql:"INTEGER DEFAULT 1"`
	ParentID         *int                   `json:"parentId"         column:"borrowing_requests.parent_id"          sql:"BIGINT"`
	RepairID         *int                   `json:"repairId"         column:"borrowing_requests.repair_id"          sql:"BIGINT"`
	TrackingID       *int                   `json:"trackingId"       column:"borrowing_requests.tracking_id"        sql:"BIGINT"`
	Status           BorrowingRequestStatus `json:"status"           column:"borrowing_requests.status"             sql:"borrowing_request_status_enum"`
	PreviousStatus   BorrowingRequestStatus `json:"previousStatus"   column:"borrowing_requests.previous_status"    sql:"borrowing_request_status_enum NULL DEFAULT NULL"`
	IsVerified       bool                   `json:"isVerified"       column:"borrowing_requests.is_verified"        sql:"BOOLEAN"`
	Notes            string                 `json:"notes"            column:"borrowing_requests.notes"              sql:"VARCHAR"`
	Duration         int                    `json:"duration"         column:"borrowing_requests.duration"           sql:"BIGINT"`
	ActualReturnDate *time.Time             `json:"actualReturnDate" column:"borrowing_requests.actual_return_date" sql:"TIMESTAMPTZ DEFAULT NULL"`
	DueReturnDate    time.Time              `json:"dueReturnDate"    column:"borrowing_requests.due_return_date"    sql:"TIMESTAMPTZ NOT NULL"`
	SerialNumber     *string                `json:"serialNumber"     column:"borrowing_requests.serial_number"`
	CreatedAt        time.Time              `json:"createdAt"        column:"borrowing_requests.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time              `json:"updatedAt"        column:"borrowing_requests.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingRequest) TableName() string {
	return "borrowing_requests"
}

func (m *BorrowingRequest) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_request_asset",
		"asset",
		"assets",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_request_transaction",
		"transaction",
		"borrowing_transactions",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_request_parent",
		"parent_id",
		"borrowing_requests",
		"id",
		"",
	)
}

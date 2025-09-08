package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingPrelistRequest struct {
	ID        int       `json:"id"        column:"borrowing_prelist_requests.id"         sql:"BIGINT PRIMARY KEY"`
	AssetId   int       `json:"assetId"   column:"borrowing_prelist_requests.asset_id"   sql:"BIGINT"`
	PrelistId int       `json:"prelistId" column:"borrowing_prelist_requests.prelist_id" sql:"BIGINT"`
	CreatedAt time.Time `json:"createdAt" column:"borrowing_prelist_requests.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"borrowing_prelist_requests.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingPrelistRequest) TableName() string {
	return "borrowing_prelist_requests"
}

func (m *BorrowingPrelistRequest) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_prelist_req_asset",
		"asset_id",
		"assets",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_prelist_req_prelist",
		"prelist_id",
		"borrowing_prelists",
		"id",
		"CASCADE",
	)

	// Add the unique index for prelist id and asset id
	builder.AddIndex(
		"borrowing_prelist_requests_prelist_id_asset_id",
		"prelist_id,asset_id",
		true,
		"",
	)
}

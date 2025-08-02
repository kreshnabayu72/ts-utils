package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingPrelistRequest struct {
	ID        int       `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetId   int       `json:"assetId" column:"asset_id" sql:"BIGINT"`
	PrelistId int       `json:"prelistId" column:"prelist_id" sql:"BIGINT"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (BorrowingPrelistRequest) TableName() string {
	return "borrowing_prelist_requests"
}

func (BorrowingPrelistRequest) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(BorrowingPrelistRequest{})

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
		"",
	)
}

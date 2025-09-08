package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingPrelist struct {
	ID         int       `json:"id"         column:"borrowing_prelists.id"          sql:"BIGINT PRIMARY KEY"`
	GroupId    int       `json:"groupId"    column:"borrowing_prelists.group_id"    sql:"BIGINT"`
	BorrowerId int       `json:"borrowerId" column:"borrowing_prelists.borrower_id" sql:"BIGINT"`
	CreatedAt  time.Time `json:"createdAt"  column:"borrowing_prelists.created_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt"  column:"borrowing_prelists.updated_at"  sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *BorrowingPrelist) TableName() string {
	return "borrowing_prelists"
}

func (m *BorrowingPrelist) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"fk_borrowing_prelists_group",
		"group_id",
		"groups",
		"id",
		"",
	)

	builder.AddForeignKey(
		"fk_borrowing_prelists_borrower",
		"borrower_id",
		"users",
		"id",
		"",
	)

	// Add the unique index for user id, group id
	builder.AddIndex(
		"borrowing_prelists_group_id_borrower_id",
		"group_id,borrower_id",
		true,
		"",
	)
}

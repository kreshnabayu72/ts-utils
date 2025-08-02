package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type BorrowingPrelist struct {
	ID         int       `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	GroupId    int       `json:"groupId" column:"group_id" sql:"BIGINT"`
	BorrowerId int       `json:"borrowerId" column:"borrower_id" sql:"BIGINT"`
	CreatedAt  time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
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
}

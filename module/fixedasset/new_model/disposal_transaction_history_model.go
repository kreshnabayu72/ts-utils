package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalTransactionHistory represents the disposal_transaction_histories table in the database.
type DisposalTransactionHistory struct {
	ID                 int                       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	TransactionID      string                    `json:"transactionId" column:"transaction_id" sql:"VARCHAR UNIQUE"`
	Manager            int                       `json:"managerId" column:"manager_id" sql:"BIGINT"`
	OriginGroupID      int                       `json:"originGroupId" column:"origin_group_id" sql:"BIGINT"`
	DisposalGroupID    int                       `json:"disposalGroupId" column:"disposal_group_id" sql:"BIGINT"`
	OriginGroup        string                    `json:"originGroup" column:"origin_group" sql:"VARCHAR"`
	DisposalGroup      string                    `json:"disposalGroup" column:"disposal_group" sql:"VARCHAR"`
	Status             DisposalTransactionStatus `json:"status" column:"status" sql:"disposal_transaction_status_enum"`
	Notes              string                    `json:"notes" column:"notes" sql:"VARCHAR"`
	CompletionNotes    string                    `json:"completionNotes" column:"completion_notes" sql:"VARCHAR"`
	LastCondition      DisposalLastCondition     `json:"lastCondition" column:"last_condition" sql:"disposal_last_condition_enum"`
	Method             DisposalMethod            `json:"method" column:"method" sql:"disposal_method_enum"`
	Organization       string                    `json:"organization" column:"organization" sql:"VARCHAR"`
	DocumentationPhoto string                    `json:"documentationPhoto" column:"documentation_photo" sql:"VARCHAR"`
	TotalAssets        int                       `json:"totalAssets" column:"total_assets" sql:"BIGINT"`
	CreatedAt          time.Time                 `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time                 `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalTransactionHistory) TableName() string {
	return "disposal_transaction_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalTransactionHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_disposal_transaction_manager",
		"manager_id",
		"users",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_transaction_origin_group",
		"origin_group_id",
		"groups",
		"id",
		"",
	)
	builder.AddForeignKey(
		"fk_disposal_transaction_destination_group",
		"destination_group_id",
		"groups",
		"id",
		"",
	)

	// Add the unique index for transaction_id
	builder.AddIndex("disposal_transactions_transaction_id", "transaction_id", true, "")
}

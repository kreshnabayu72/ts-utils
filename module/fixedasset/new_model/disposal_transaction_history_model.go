package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

// DisposalTransactionHistory represents the disposal_transaction_histories table in the database.
type DisposalTransactionHistory struct {
	ID                 int                       `json:"_id"                column:"disposal_transaction_histories.id"                  sql:"BIGINT PRIMARY KEY"`
	TransactionID      string                    `json:"transactionId"      column:"disposal_transaction_histories.transaction_id"      sql:"VARCHAR UNIQUE"`
	Manager            int                       `json:"managerId"          column:"disposal_transaction_histories.manager_id"          sql:"BIGINT"`
	OriginGroupID      int                       `json:"originGroupId"      column:"disposal_transaction_histories.origin_group_id"     sql:"BIGINT"`
	DisposalGroupID    int                       `json:"disposalGroupId"    column:"disposal_transaction_histories.disposal_group_id"   sql:"BIGINT"`
	OriginGroup        string                    `json:"originGroup"        column:"disposal_transaction_histories.origin_group"        sql:"VARCHAR"`
	DisposalGroup      string                    `json:"disposalGroup"      column:"disposal_transaction_histories.disposal_group"      sql:"VARCHAR"`
	Status             DisposalTransactionStatus `json:"status"             column:"disposal_transaction_histories.status"              sql:"disposal_transaction_status_enum"`
	Notes              string                    `json:"notes"              column:"disposal_transaction_histories.notes"               sql:"VARCHAR"`
	CompletionNotes    string                    `json:"completionNotes"    column:"disposal_transaction_histories.completion_notes"    sql:"VARCHAR"`
	LastCondition      DisposalLastCondition     `json:"lastCondition"      column:"disposal_transaction_histories.last_condition"      sql:"disposal_last_condition_enum"`
	Method             DisposalMethod            `json:"method"             column:"disposal_transaction_histories.method"              sql:"disposal_method_enum"`
	CustomMethod       string                    `json:"customMethod"       column:"disposal_transaction_histories.custom_method"       sql:"VARCHAR"`
	Organization       string                    `json:"organization"       column:"disposal_transaction_histories.organization"        sql:"VARCHAR"`
	DocumentationPhoto string                    `json:"documentationPhoto" column:"disposal_transaction_histories.documentation_photo" sql:"VARCHAR"`
	TotalAssets        int                       `json:"totalAssets"        column:"disposal_transaction_histories.total_assets"        sql:"BIGINT"`
	CreatedAt          time.Time                 `json:"createdAt"          column:"disposal_transaction_histories.created_at"          sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time                 `json:"updatedAt"          column:"disposal_transaction_histories.updated_at"          sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

// TableName returns the name of the database table for this model.
func (m *DisposalTransactionHistory) TableName() string {
	return "disposal_transaction_histories"
}

// Up defines the migration logic to create the table and its constraints.
func (m *DisposalTransactionHistory) Up(builder *migration.MigrationBuilder) {
	// Add all columns from the struct
	builder.AddColumn(m)

	// Add the unique index for transaction_id
	builder.AddIndex("disposal_transaction_histories_transaction_id", "transaction_id", true, "")
}

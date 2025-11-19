package model

import (
	"time"
)

// DisposalTransactionHistory represents the disposal_transaction_histories table in the database.
type DisposalTransactionHistory struct {
	ID                 int                       `json:"_id"                column:"disposal_transaction_histories.id"`
	TransactionID      string                    `json:"transactionId"      column:"disposal_transaction_histories.transaction_id"`
	Manager            int                       `json:"managerId"          column:"disposal_transaction_histories.manager_id"`
	OriginGroupID      int                       `json:"originGroupId"      column:"disposal_transaction_histories.origin_group_id"`
	DisposalGroupID    int                       `json:"disposalGroupId"    column:"disposal_transaction_histories.disposal_group_id"`
	OriginGroup        string                    `json:"originGroup"        column:"disposal_transaction_histories.origin_group"`
	DisposalGroup      string                    `json:"disposalGroup"      column:"disposal_transaction_histories.disposal_group"`
	Status             DisposalTransactionStatus `json:"status"             column:"disposal_transaction_histories.status"`
	Notes              string                    `json:"notes"              column:"disposal_transaction_histories.notes"`
	CompletionNotes    string                    `json:"completionNotes"    column:"disposal_transaction_histories.completion_notes"`
	LastCondition      DisposalLastCondition     `json:"lastCondition"      column:"disposal_transaction_histories.last_condition"`
	Method             DisposalMethod            `json:"method"             column:"disposal_transaction_histories.method"`
	CustomMethod       string                    `json:"customMethod"       column:"disposal_transaction_histories.custom_method"`
	Organization       string                    `json:"organization"       column:"disposal_transaction_histories.organization"`
	DocumentationPhoto string                    `json:"documentationPhoto" column:"disposal_transaction_histories.documentation_photo"`
	TotalAssets        int                       `json:"totalAssets"        column:"disposal_transaction_histories.total_assets"`
	CreatedAt          time.Time                 `json:"createdAt"          column:"disposal_transaction_histories.created_at"`
	UpdatedAt          time.Time                 `json:"updatedAt"          column:"disposal_transaction_histories.updated_at"`
}

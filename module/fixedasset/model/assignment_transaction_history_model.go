package model

import (
	"time"
)

// AssignmentTransactionHistory represents the assignment_transaction_histories table in the database.
type AssignmentTransactionHistory struct {
	ID            int                         `json:"_id"           column:"assignment_transaction_histories.id"`
	TransactionID string                      `json:"transactionId" column:"assignment_transaction_histories.transaction_id"`
	GroupID       int                         `json:"groupId"       column:"assignment_transaction_histories.group_id"`
	Group         string                      `json:"group"         column:"assignment_transaction_histories.group"`
	AssignedTo    string                      `json:"assignedTo"    column:"assignment_transaction_histories.assigned_to"`
	Manager       string                      `json:"manager"       column:"assignment_transaction_histories.manager"`
	TotalAssets   int                         `json:"totalAssets"   column:"assignment_transaction_histories.total_assets"`
	Status        AssignmentTransactionStatus `json:"status"        column:"assignment_transaction_histories.status"`
	CreatedAt     time.Time                   `json:"createdAt"     column:"assignment_transaction_histories.created_at"`
	UpdatedAt     time.Time                   `json:"updatedAt"     column:"assignment_transaction_histories.updated_at"`
}

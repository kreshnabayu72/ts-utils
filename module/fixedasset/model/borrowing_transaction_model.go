package model

import (
	"time"
)

type BorrowingTransaction struct {
	ID                           int                        `json:"_id"                          column:"borrowing_transactions.id"`
	TransactionID                string                     `json:"transactionId"                column:"borrowing_transactions.transaction_id"`
	BorrowerID                   int                        `json:"borrowerId"                   column:"borrowing_transactions.borrower_id"`
	ManagerID                    int                        `json:"managerId"                    column:"borrowing_transactions.manager_id"`
	GroupID                      int                        `json:"groupId"                      column:"borrowing_transactions.group_id"`
	Status                       BorrowingTransactionStatus `json:"status"                       column:"borrowing_transactions.status"`
	IsSelfService                bool                       `json:"isSelfService"                column:"borrowing_transactions.is_self_service"`
	IsVerified                   bool                       `json:"isVerified"                   column:"borrowing_transactions.is_verified"`
	ConfirmationEmail            string                     `json:"confirmationEmail"            column:"borrowing_transactions.confirmation_email"`
	ConfirmationEmailConfirmed   bool                       `json:"confirmationEmailConfirmed"   column:"borrowing_transactions.confirmation_email_confirmed"`
	ConfirmationEmailSentAt      *time.Time                 `json:"confirmationEmailSentAt"      column:"borrowing_transactions.confirmation_email_sent_at"`
	ConfirmationEmailExpiredAt   *time.Time                 `json:"confirmationEmailExpiredAt"   column:"borrowing_transactions.confirmation_email_expired_at"`
	ConfirmationEmailToken       string                     `json:"confirmationEmailToken"       column:"borrowing_transactions.confirmation_email_token"`
	ConfirmationEmailConfirmedAt *time.Time                 `json:"confirmationEmailConfirmedAt" column:"borrowing_transactions.confirmation_email_confirmed_at"`
	CurrentApprovalLevel         int                        `json:"currentApprovalLevel"         column:"borrowing_transactions.current_approval_level"`
	CreatedAt                    time.Time                  `json:"createdAt"                    column:"borrowing_transactions.created_at"`
	UpdatedAt                    time.Time                  `json:"updatedAt"                    column:"borrowing_transactions.updated_at"`
}

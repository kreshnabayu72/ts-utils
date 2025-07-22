package new_model

import "time"

type AssignmentTransaction struct {
	ID            int                         `json:"_id" column:"id"`
	TransactionID string                      `json:"transactionId" column:"transaction_id"`
	AssignedToID  int                         `json:"assignedToId" column:"assigned_to_id"`
	ManagerID     int                         `json:"managerId" column:"manager_id"`
	GroupID       int                         `json:"groupId" column:"group_id"`
	Status        AssignmentTransactionStatus `json:"status" column:"status"`
	IsVerified    bool                        `json:"isVerified" column:"is_verified"`

	ConfirmationEmail            string    `json:"confirmationEmail" column:"confirmation_email"`
	ConfirmationEmailConfirmed   bool      `json:"confirmationEmailConfirmed" column:"confirmation_email_confirmed"`
	ConfirmationEmailSentAt      time.Time `json:"confirmationEmailSentAt" column:"confirmation_email_sent_at"`
	ConfirmationEmailExpiredAt   time.Time `json:"confirmationEmailExpiredAt" column:"confirmation_email_expired_at"`
	ConfirmationEmailToken       string    `json:"confirmationEmailToken" column:"confirmation_email_token"`
	ConfirmationEmailConfirmedAt time.Time `json:"confirmationEmailConfirmedAt" column:"confirmation_email_confirmed_at"`

	CurrentApprovalLevel int `json:"currentApprovalLevel" column:"current_approval_level"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

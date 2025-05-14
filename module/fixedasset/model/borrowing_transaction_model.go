package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BorrowingTransactionStatus string

const (
	BorrowingTransactionStatusRejected           = BorrowingTransactionStatus("Rejected")
	BorrowingTransactionStatusReturned           = BorrowingTransactionStatus("Returned")
	BorrowingTransactionStatusWaitingForApproval = BorrowingTransactionStatus("Waiting for Approval")
	BorrowingTransactionStatusWaitingForHandover = BorrowingTransactionStatus("Waiting for Handover")
	BorrowingTransactionStatusBorrowed           = BorrowingTransactionStatus("Borrowed") // Previously On Borrowing
)

type BorrowingTransaction struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	TransactionId string                     `json:"transactionId" bson:"transactionId"`
	Status        BorrowingTransactionStatus `json:"status" bson:"status"`
	Group         model.AttributeEmbedded    `json:"group" bson:"group"`

	Manager  model.UserEmbedded `json:"manager" bson:"manager"`
	Borrower model.UserEmbedded `json:"borrower" bson:"borrower"`

	ConfirmationEmail string `json:"confirmationEmail" bson:"confirmationEmail"`

	DueReturnDate    time.Time `json:"dueReturnDate" bson:"dueReturnDate"`
	ActualReturnDate time.Time `json:"actualReturnDate" bson:"actualReturnDate"`
	CreatedAt        time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt" bson:"updatedAt"`
}

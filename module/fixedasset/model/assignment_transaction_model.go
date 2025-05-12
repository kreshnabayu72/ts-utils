package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssignmentTransactionStatus string

const (
	AssignmentTransactionStatusRejected           = AssignmentTransactionStatus("Rejected")
	AssignmentTransactionStatusReturned           = AssignmentTransactionStatus("Returned")
	AssignmentTransactionStatusWaitingForApproval = AssignmentTransactionStatus("Waiting for Approval")
	AssignmentTransactionStatusWaitingForHandover = AssignmentTransactionStatus("Waiting for Handover")
	AssignmentTransactionStatusAssigned           = AssignmentTransactionStatus("Assigned")
)

type AssignmentTransaction struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	TransactionId string                      `json:"transactionId" bson:"transactionId"`
	Status        AssignmentTransactionStatus `json:"status" bson:"status"`
	Group         model.AttributeEmbedded     `json:"group" bson:"group"`

	Manager  model.UserEmbedded `json:"manager" bson:"manager"`
	Assigner model.UserEmbedded `json:"borrower" bson:"borrower"`

	ConfirmationEmail string `json:"confirmationEmail" bson:"confirmationEmail"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

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
	AssignmentTransactionStatusWaitingForApproval = AssignmentTransactionStatus(
		"Waiting for Approval",
	)
	AssignmentTransactionStatusWaitingForHandover = AssignmentTransactionStatus(
		"Waiting for Handover",
	)
	AssignmentTransactionStatusAssigned   = AssignmentTransactionStatus("Assigned")
	AssignmentTransactionStatusUnassigned = AssignmentTransactionStatus("Unassigned")
)

type AssignmentTransaction struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	TransactionId string                      `json:"transactionId" bson:"transactionId"`
	Status        AssignmentTransactionStatus `json:"status" bson:"status"`
	Group         model.AttributeEmbedded     `json:"group" bson:"group"`

	Manager    model.UserEmbedded `json:"manager" bson:"manager"`
	AssignedTo model.UserEmbedded `json:"assignedTo" bson:"assignedTo"`

	ConfirmationEmail string `json:"confirmationEmail" bson:"confirmationEmail"`
	TotalAssets       int    `json:"totalAssets" bson:"totalAssets"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

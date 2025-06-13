package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReceptionTransactionStatus string

const (
	ReceptionTransactionStatusInReception        ReceptionTransactionStatus = "In Reception"
	ReceptionTransactionStatusWaitingForApproval ReceptionTransactionStatus = "Waiting for Approval"
	ReceptionTransactionStatusCompleted          ReceptionTransactionStatus = "Completed"
	ReceptionTransactionStatusCanceled           ReceptionTransactionStatus = "Canceled"
	ReceptionTransactionStatusReceived           ReceptionTransactionStatus = "Received"
)

// ReceptionTransaction represents the active reception transaction that groups ReceptionRequest
//
// Reception transaction is a process to initiate a transaction to receive stock(s) from supplier
type ReceptionTransaction struct {
	ID             primitive.ObjectID         `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  string                     `json:"transactionId"  bson:"transactionId"`
	Supplier       primitive.ObjectID         `json:"supplier"       bson:"supplier"`
	Status         ReceptionTransactionStatus `json:"status"         bson:"status"`
	ReceivingGroup primitive.ObjectID         `json:"receivingGroup" bson:"receivingGroup"`
	Document       *string                    `json:"document"       bson:"document"`
	Note           *string                    `json:"note"           bson:"note"`
	Photo          *string                    `json:"photo"          bson:"photo"`
	PIC            primitive.ObjectID         `json:"pic"            bson:"pic"`
	Log            []TransactionLogDetail     `json:"log"            bson:"log"`
	TotalItem      int                        `json:"totalItem"      bson:"totalItem"`
	TotalStock     int                        `json:"totalStock"     bson:"totalStock"`
	ApprovalNote   *string                    `json:"approvalNote"   bson:"approvalNote"`
	UpdatedAt      time.Time                  `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                  `json:"createdAt"      bson:"createdAt"`
}

// ReceptionTransactionHistory represents the canceled/completed transaction of ReceptionTransaction
type ReceptionTransactionHistory struct {
	ID             primitive.ObjectID         `json:"_id,omitempty"  bson:"_id,omitempty"`
	TransactionId  string                     `json:"transactionId"  bson:"transactionId"`
	Supplier       AttributeEmbedded          `json:"supplier"       bson:"supplier"`
	Status         ReceptionTransactionStatus `json:"status"         bson:"status"`
	ReceivingGroup TreeAttributeEmbedded      `json:"receivingGroup" bson:"receivingGroup"`
	Document       *string                     `json:"document"       bson:"document"`
	Note           *string                    `json:"note"           bson:"note"`
	Photo          *string                    `json:"photo"          bson:"photo"`
	PIC            UserAttributeEmbedded      `json:"pic"            bson:"pic"`
	Log            []TransactionLogDetail     `json:"log"            bson:"log"`
	TotalItem      int                        `json:"totalItem"      bson:"totalItem"`
	TotalStock     int                        `json:"totalStock"     bson:"totalStock"`
	ApprovalNote   *string                    `json:"approvalNote"   bson:"approvalNote"`
	UpdatedAt      time.Time                  `json:"updatedAt"      bson:"updatedAt"`
	CreatedAt      time.Time                  `json:"createdAt"      bson:"createdAt"`
}

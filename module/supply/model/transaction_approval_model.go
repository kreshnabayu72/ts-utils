package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TransactionApprovalFor represents the possible values to "which object this approval belongs for"
type TransactionApprovalFor string

// TODO(refan): completes the TransactionApprovalFor possible values
const (
	TransactionApprovalForReception TransactionApprovalFor = "Reception"
	TransactionApprovalForExpired   TransactionApprovalFor = "Expired"
)

// TransactionApprovalType represents the possible values for approval type
type TransactionApprovalType string

const (
	TransactionApprovalTypeOr  TransactionApprovalType = "Or"
	TransactionApprovalTypeAnd TransactionApprovalType = "And"
)

// TransactionApprovalAction represents the possible values for actionable approval
type TransactionApprovalAction string

const (
	TransactionApprovalActionApproved TransactionApprovalAction = "Approved"
	TransactionApprovalActionRejected TransactionApprovalAction = "Rejected"
)

// TransactionApprovalApprover stores users and it's actions (rejected/approved/null)
type TransactionApprovalApprover[IO any] struct {
	User         primitive.ObjectID         `json:"user"         bson:"user"`
	UserFullName string                     `json:"userFullName" bson:"userFullName"`
	ActionDate   *time.Time                 `json:"actionAt"     bson:"actionAt"`
	Action       *TransactionApprovalAction `json:"action"       bson:"action"`
	Data         *IO                        `json:"data"         bson:"data"`
}

// TransactionApproval stores users and its actions data for each approval level
//
// Generic type "O" represents the struct object to be used in the approval list
//
// While Generic type "IO" represents the struct object of the user's input on each approval
type TransactionApproval[O any, IO any] struct {
	ID              primitive.ObjectID                `json:"_id,omitempty"   bson:"_id,omitempty"`
	For             TransactionApprovalFor            `json:"for"             bson:"for"`
	ApprovalBatchId primitive.ObjectID                `json:"approvalBatchId" bson:"approvalBatchId"`
	ObjectId        primitive.ObjectID                `json:"objectId"        bson:"objectId"`
	Object          O                                 `json:"object"          bson:"object"`
	Level           int                               `json:"level"           bson:"level"`
	IsActivated     bool                              `json:"isActivated"     bson:"isActivated"`
	Type            TransactionApprovalType           `json:"type"            bson:"type"`
	Approvers       []TransactionApprovalApprover[IO] `json:"approvers"       bson:"approvers"`
	UpdatedAt       time.Time                         `json:"updatedAt"       bson:"updatedAt"`
	CreatedAt       time.Time                         `json:"createdAt"       bson:"createdAt"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ApprovalFor represents the possible values to "which object this approval belongs for"
type ApprovalFor string

// TODO(refan): completes the ApprovalFor possible values
const (
	TransactionApprovalForReception ApprovalFor = "Reception"
	TransactionApprovalForExpired   ApprovalFor = "Expired"
)

// ApprovalType represents the possible values for approval type
type ApprovalType string

const (
	ApprovalTypeOr  ApprovalType = "Or"
	ApprovalTypeAnd ApprovalType = "And"
)

// ApprovalAction represents the possible values for actionable approval
type ApprovalAction string

const (
	ApprovalActionApproved ApprovalAction = "Approved"
	ApprovalActionRejected ApprovalAction = "Rejected"
)

// ApprovalApprover stores users and it's actions (rejected/approved/null)
type ApprovalApprover[IO any] struct {
	User         primitive.ObjectID `json:"user"         bson:"user"`
	UserFullName string             `json:"userFullName" bson:"userFullName"`
	ActionDate   *time.Time         `json:"actionAt"     bson:"actionAt"`
	Action       *ApprovalAction    `json:"action"       bson:"action"`
	Data         *IO                `json:"data"         bson:"data"`
}

// Approval stores users and its actions data for each approval level
//
// Generic type "O" represents the struct object to be used in the approval list
//
// While Generic type "IO" represents the struct object of the user's input on each approval
type Approval[O any, IO any] struct {
	ID              primitive.ObjectID     `json:"_id,omitempty"   bson:"_id,omitempty"`
	For             ApprovalFor            `json:"for"             bson:"for"`
	ApprovalBatchId primitive.ObjectID     `json:"approvalBatchId" bson:"approvalBatchId"`
	ObjectId        primitive.ObjectID     `json:"objectId"        bson:"objectId"`
	Object          O                      `json:"object"          bson:"object"`
	Level           int                    `json:"level"           bson:"level"`
	IsActivated     bool                   `json:"isActivated"     bson:"isActivated"`
	Type            ApprovalType           `json:"type"            bson:"type"`
	Approvers       []ApprovalApprover[IO] `json:"approvers"       bson:"approvers"`
	UpdatedAt       time.Time              `json:"updatedAt"       bson:"updatedAt"`
	CreatedAt       time.Time              `json:"createdAt"       bson:"createdAt"`
}

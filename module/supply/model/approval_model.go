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

// ApprovalStatus represents the possible values for the status of approval
type ApprovalStatus string

const (
	ApprovalStatusWaitingForApproval ApprovalStatus = "Waiting for Approval"
	ApprovalStatusRejected           ApprovalStatus = "Rejected"
	ApprovalStatusApproved           ApprovalStatus = "Approved"
)

// Approval stores users and its actions data for each approval level
//
// Generic type "O" represents the struct object to be used in the approval list
//
// While Generic type "IO" represents the struct object of the user's input on each approval
//
// # Approval Flow Pseudo-code
//
// 1. Create product > Send Create Approval
//
// In Approval Service
//  2. Check for approver in settings approval product
//  3. If approver in settings approval product exists:
//     create approval product of all level for every approver
//     3.1. If approval type is "Or" and an approver approved:
//     3.1.1. If higher settings approval product level exists:
//     jump to 3, activating higher level than the current level
//     3.1.2 Else:
//     > Finish Approval as Approved
//     3.2. Else If approval type is "Or" and an approver rejected:
//     3.2.1. If other approver exists:
//     jump to 3, without the creation process
//     3.2.2. Else:
//     > Finish Approval as Rejected
//     3.3. Else If approval type is "And" and an approver approved:
//     3.3.1. If other approver exists:
//     jump to 3, without the creation process
//     3.3.2. Else:
//     3.3.2.1. If higher settings approval product level exists:
//     jump to 3, activating higher level than the current level
//     3.3.2.2. Else:
//     > Finish Approval as Approved
//     3.4. Else If approval type is "And" and an approver rejected:
//     > Finish Approval as Rejected
//  4. Else:
//     > Finish Approval as Approved
//
// In Product Service
//  5. Subscribe for Finish Approval, and do whatever it wants with it.
//
// Note:
// - Approver who were absent in the approval process is marked as "Not Involved"
type Approval[O any, IO any] struct {
	ID              primitive.ObjectID     `json:"_id,omitempty"   bson:"_id,omitempty"`
	For             ApprovalFor            `json:"for"             bson:"for"`
	ApprovalBatchId primitive.ObjectID     `json:"approvalBatchId" bson:"approvalBatchId"`
	ObjectId        primitive.ObjectID     `json:"objectId"        bson:"objectId"`
	Object          O                      `json:"object"          bson:"object"`
	Level           int                    `json:"level"           bson:"level"`
	Status          ApprovalStatus         `json:"status"          bson:"status"`
	IsActivated     bool                   `json:"isActivated"     bson:"isActivated"`
	Type            ApprovalType           `json:"type"            bson:"type"`
	Approvers       []ApprovalApprover[IO] `json:"approvers"       bson:"approvers"`
	UpdatedAt       time.Time              `json:"updatedAt"       bson:"updatedAt"`
	CreatedAt       time.Time              `json:"createdAt"       bson:"createdAt"`
}

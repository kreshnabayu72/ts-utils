package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// TransactionApproval represents the transactional approval data
type TransactionApproval struct {
	ID                  primitive.ObjectID   `json:"_id,omitempty"       bson:"_id,omitempty"`
	AdministrativeGroup primitive.ObjectID   `json:"administrativeGroup" bson:"administrativeGroup"`
	ApprovalLevel       int                  `json:"approvalLevel"       bson:"approvalLevel"`
	ApprovalType        ApprovalType         `json:"approvalType"        bson:"approvalType"`
	TransactionRoleType TransactionRoleType  `json:"transactionRoleType" bson:"transactionRoleType"`
	Approvers           []primitive.ObjectID `json:"approvers"           bson:"approvers"`
}

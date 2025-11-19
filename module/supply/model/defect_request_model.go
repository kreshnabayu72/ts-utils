package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
)

type DefectRequestStatus string

const (
	DefectRequestStatusReported           DefectRequestStatus = "Reported"
	DefectRequestStatusWaitingForApproval DefectRequestStatus = "Waiting for Approval"
	DefectRequestStatusUpdated            DefectRequestStatus = "Updated"
	DefectRequestStatusDisposed           DefectRequestStatus = "Disposed"
)

// DefectRequest represents the individual request of a stock reported as defect
type DefectRequest struct {
	ID            primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	BatchId       primitive.ObjectID  `json:"batchId"       bson:"BatchId"`
	Status        DefectRequestStatus `json:"status"        bson:"status"`
	ItemName      primitive.ObjectID  `json:"itemName"      bson:"itemName"`
	ItemSKU       primitive.ObjectID  `json:"itemSKU"       bson:"itemSKU"`
	Brand         primitive.ObjectID  `json:"brand"         bson:"brand"`
	Model         primitive.ObjectID  `json:"model"         bson:"model"`
	Stock         primitive.ObjectID  `json:"stock"         bson:"stock"`
	TagType       model.TagType       `json:"tagType"       bson:"tagType"`
	Note          *string             `json:"note"          bson:"note"`
	ReportedBy    primitive.ObjectID  `json:"reportedBy"    bson:"reportedBy"`
	ReportedDate  time.Time           `json:"reportedDate"  bson:"reportedDate"`
	Photo         string              `json:"photo"         bson:"photo"`
	ApprovalNote  *string             `json:"approvalNote"  bson:"approvalNote"`
	AvailableDate *time.Time          `json:"availableDate" bson:"availableDate"`
	UpdatedAt     time.Time           `json:"updatedAt"     bson:"updatedAt"`
	CreatedAt     time.Time           `json:"createdAt"     bson:"createdAt"`
}

// DefectRequestHistory represents the completed request of DefectRequest
type DefectRequestHistory struct {
	ID            primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	BatchId       primitive.ObjectID    `json:"batchId"       bson:"batchId"`
	Status        DefectRequestStatus   `json:"status"        bson:"status"`
	ItemName      AttributeEmbedded     `json:"itemName"      bson:"itemName"`
	ItemSKU       SKUAttributeEmbedded  `json:"itemSKU"       bson:"itemSKU"`
	Brand         AttributeEmbedded     `json:"brand"         bson:"brand"`
	Model         AttributeEmbedded     `json:"model"         bson:"model"`
	Stock         Stock                 `json:"stock"         bson:"stock"`
	TagType       model.TagType         `json:"tagType"       bson:"tagType"`
	Note          *string               `json:"note"          bson:"note"`
	ReportedBy    UserAttributeEmbedded `json:"reportedBy"    bson:"reportedBy"`
	ReportedDate  time.Time             `json:"reportedDate"  bson:"reportedDate"`
	Photo         *string               `json:"photo"         bson:"photo"`
	ApprovalNote  *string               `json:"approvalNote"  bson:"approvalNote"`
	AvailableDate *time.Time            `json:"availableDate" bson:"availableDate"`
	UpdatedAt     time.Time             `json:"updatedAt"     bson:"updatedAt"`
	CreatedAt     time.Time             `json:"createdAt"     bson:"createdAt"`
}

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
)

type ExpiredRequestStatus string

const (
	ExpiredRequestStatusReported           ExpiredRequestStatus = "Reported"
	ExpiredRequestStatusWaitingForApproval ExpiredRequestStatus = "Waiting for Approval"
	ExpiredRequestStatusUpdated            ExpiredRequestStatus = "Updated"
	ExpiredRequestStatusDisposed           ExpiredRequestStatus = "Disposed"
)

// ExpiredRequest represents the individual request of a stock reported as expired
type ExpiredRequest struct {
	ID           primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	BatchId      primitive.ObjectID   `json:"batchId"       bson:"BatchId"`
	Status       ExpiredRequestStatus `json:"status"        bson:"status"`
	ItemName     primitive.ObjectID   `json:"itemName"      bson:"itemName"`
	ItemSKU      primitive.ObjectID   `json:"itemSKU"       bson:"itemSKU"`
	Brand        primitive.ObjectID   `json:"brand"         bson:"brand"`
	Model        primitive.ObjectID   `json:"model"         bson:"model"`
	Stock        primitive.ObjectID   `json:"stock"         bson:"stock"`
	TagType      model.TagType        `json:"tagType"       bson:"tagType"`
	Note         *string              `json:"note"          bson:"note"`
	ReportedBy   primitive.ObjectID   `json:"reportedBy"    bson:"reportedBy"`
	ReportedDate time.Time            `json:"reportedDate"  bson:"reportedDate"`
	UpdatedDate  *time.Time           `json:"updatedDate"   bson:"updatedDate"`
	UpdatedAt    time.Time            `json:"updatedAt"     bson:"updatedAt"`
	CreatedAt    time.Time            `json:"createdAt"     bson:"createdAt"`
}

// ExpiredRequestHistory represents the completed request of ExpiredRequest
type ExpiredRequestHistory struct {
	ID           primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	BatchId      primitive.ObjectID    `json:"batchId"       bson:"batchId"`
	Status       ExpiredRequestStatus  `json:"status"        bson:"status"`
	ItemName     AttributeEmbedded     `json:"itemName"      bson:"itemName"`
	ItemSKU      SKUAttributeEmbedded  `json:"itemSKU"       bson:"itemSKU"`
	Brand        AttributeEmbedded     `json:"brand"         bson:"brand"`
	Model        AttributeEmbedded     `json:"model"         bson:"model"`
	Stock        Stock                 `json:"stock"         bson:"stock"`
	TagType      model.TagType         `json:"tagType"       bson:"tagType"`
	Note         *string               `json:"note"          bson:"note"`
	ReportedBy   UserAttributeEmbedded `json:"reportedBy"    bson:"reportedBy"`
	ReportedDate time.Time             `json:"reportedDate"  bson:"reportedDate"`
	UpdatedDate  *time.Time            `json:"updatedDate"   bson:"updatedDate"`
	UpdatedAt    time.Time             `json:"updatedAt"     bson:"updatedAt"`
	CreatedAt    time.Time             `json:"createdAt"     bson:"createdAt"`
}

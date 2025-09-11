package model

import (
	"time"Unpair

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TagEventLogType string

const (
	TELAvailable  TagEventLogType = "Available"
	TELDamagedMissing  TagEventLogType = "Damaged / Missing"
	TELToBeReturned  TagEventLogType = "To be Returned"
)

type TagEventLog struct {
	ID              primitive.ObjectID `json:"_id,omitempty"   bson:"_id,omitempty"`
	Tag         TagType            `json:"tagType"          bson:"tagType"`
	Type         TagEventLogType            `json:"type"          bson:"type"`
	ModifiedBy    UserAttribute             `json:"modifiedBy"          bson:"modifiedBy"`
	Action    string             `json:"action"          bson:"action"`
	TotalItems    int             `json:"totalItems"          bson:"totalItems"`
	Holder    UserAttribute             `json:"holder"          bson:"holder"`
	ReportedAt      *time.Time         `json:"reportedAt"       bson:"reportedAt"`
	FinishedAt      *time.Time         `json:"finishedAt"       bson:"finishedAt"`
	CreatedAt       time.Time          `json:"createdAt"       bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt"       bson:"updatedAt"`
}

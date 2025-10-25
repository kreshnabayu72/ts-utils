package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Changelog struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Action        string             `json:"action"        bson:"action"`
	Field         *string            `json:"field"         bson:"field"`
	OldValue      string             `json:"oldValue"      bson:"oldValue"`
	NewValue      string             `json:"newValue"      bson:"newValue"`
	Value         string             `json:"value"         bson:"value"`
	ModifiedBy    string             `json:"modifiedBy"    bson:"modifiedBy"`
	ModifiedById  primitive.ObjectID `json:"modifiedById"  bson:"modifiedById"`
	ModifiedByKey string             `json:"modifiedByKey" bson:"modifiedByKey"`
	Object        string             `json:"object"        bson:"object"`
	ObjectId      primitive.ObjectID `json:"objectId"      bson:"objectId"`
	ObjectKey     string             `json:"objectKey"     bson:"objectKey"`

	ObjectName  string              `json:"objectName"  bson:"objectName"`
	AssetName   string              `json:"assetName"   bson:"assetName"`
	GroupId     *primitive.ObjectID `json:"groupId"     bson:"groupId"`
	GroupKey    *string             `json:"groupKey"    bson:"groupKey"`
	Module      string              `json:"module"      bson:"module"`
	ActivityLog bool                `json:"activityLog" bson:"activityLog"`
	CreatedAt   time.Time           `json:"createdAt"   bson:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"   bson:"updatedAt"`
}

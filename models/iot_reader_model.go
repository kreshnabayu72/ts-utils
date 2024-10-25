package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IOTReader struct {
	ID                primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	SerialNumber      string             `json:"serialNumber" bson:"serialNumber"`
	Name              AssetName          `json:"name" bson:"name"`
	IsAutoTransaction bool               `json:"isAutoTransaction" bson:"isAutoTransaction"`
	IsRoleRestricted  bool               `json:"isRoleRestricted" bson:"isRoleRestricted"`
	GroupInFrom       AssetAttribute     `json:"groupInFrom" bson:"groupInFrom"`
	GroupOutTo        AssetAttribute     `json:"groupOutTo" bson:"groupOutTo"`
	Key               int32              `json:"key" bson:"key"`
}

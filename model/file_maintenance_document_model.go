package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FileMaintenanceDocument struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AssetName  AssetName          `bson:"assetName" json:"assetName"`
	FileName   string             `bson:"fileName" json:"fileName"`
	FilePath   string             `bson:"filePath" json:"filePath"`
	FileFormat string             `bson:"fileFormat" json:"fileFormat"`
	FileSize   string             `bson:"fileSize" json:"fileSize"`
	FileSizeKb float64            `bson:"fileSizeKb" json:"fileSizeKb"`
	IsDeleted  bool               `bson:"isDeleted" json:"isDeleted"`
	ModifiedBy ModifiedBy         `bson:"modifiedBy" json:"modifiedBy"`
	DeletedAt  time.Time          `bson:"deletedAt" json:"deletedAt"`
	Group      GroupAttribute     `bson:"group" json:"group"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt  time.Time          `bson:"updatedAt" json:"updatedAt"`
}

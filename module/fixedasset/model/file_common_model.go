package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// For creating file (currently only planned for basic)
type FileCommon struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"    json:"_id,omitempty"`
	AssetName  AssetNameAttribute `bson:"assetName"        json:"assetName"`
	FileName   string             `bson:"fileName"         json:"fileName"`
	FilePath   string             `bson:"filePath"         json:"filePath"`
	FileFormat string             `bson:"fileFormat"       json:"fileFormat"`
	FileSize   string             `bson:"fileSize"         json:"fileSize"`
	FileSizeKb float64            `bson:"fileSizeKb"       json:"fileSizeKb"`
	IsDeleted  bool               `bson:"isDeleted"        json:"isDeleted"`
	ModifiedBy model.UserEmbedded `bson:"modifiedBy"       json:"modifiedBy"`
	DeletedAt  *time.Time         `bson:"deletedAt"        json:"deletedAt"`
	Group      GroupAttribute     `bson:"group"            json:"group"`

	//for: attachment _id, report damage _id (so it can easily deleted when the relatedData is canceled/deleted)
	RelatedDataId primitive.ObjectID `bson:"relatedDataId,omitempty"    json:"relatedDataId,omitempty"`

	CreatedAt time.Time `bson:"createdAt"        json:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"        json:"updatedAt"`
}

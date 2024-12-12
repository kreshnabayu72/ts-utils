package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AssetNameTransactionAttr struct {
	Borrow      bool `json:"borrow" bson:"borrow"`
	Assign      bool `json:"assign" bson:"assign"`
	Transfer    bool `json:"transfer" bson:"transfer"`
	SelfService bool `json:"selfService" bson:"selfService"`
}

type AssetNameCategoryAttr struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int32              `json:"key" bson:"key"`
}

type AssetNameMeasurementAttr struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type AssetNameAccountCodeAttr struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int32              `json:"key" bson:"key"`
	Code string             `json:"code" bson:"code"`
}

type AssetNameAddOnAttr struct {
	RepairTicketing bool `json:"repairTicketing" bson:"repairTicketing"`
	Maintenance     bool `json:"maintenance" bson:"maintenance"`
	Tracking        bool `json:"tracking" bson:"tracking"`
	Audit           bool `json:"audit" bson:"audit"`
}

type AssetNameAssetStatusAttr struct {
	Status  string     `json:"status" bson:"status"`
	AddedAt *time.Time `json:"addedAt" bson:"addedAt"`
}

type AssetNameLicenseAttr struct {
	FixedAsset AssetNameAssetStatusAttr `json:"fixedAsset" bson:"fixedAsset"`
}

type AssetNameDepreciationGroupAttr = AssetNameCategoryAttr

type AssetName struct {
	ID                   primitive.ObjectID              `json:"_id,omitempty" bson:"_id,omitempty"`
	Key                  int32                           `json:"key" bson:"key"`
	Name                 string                          `json:"name" bson:"name"`
	AliasCode            string                          `json:"aliasCode" bson:"aliasCode"`
	TotalAssets          int32                           `json:"totalAssets" bson:"totalAssets"`
	IsDeleted            bool                            `json:"isDeleted" bson:"isDeleted"`
	TagType              *string                         `json:"tagType" bson:"tagType"`
	DeleteFromDatabaseAt *time.Time                      `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	Transactions         AssetNameTransactionAttr        `json:"transactions" bson:"transactions"`
	Category             *AssetNameCategoryAttr          `json:"category" bson:"category"`
	Measurement          *AssetNameMeasurementAttr       `json:"measurement" bson:"measurement"`
	DepreciationGroup    *AssetNameDepreciationGroupAttr `json:"depreciationGroup" bson:"depreciationGroup"`
	AccountCode          *AssetNameAccountCodeAttr       `json:"accountCode" bson:"accountCode"`
	AddOn                *AssetNameAddOnAttr             `json:"addOn" bson:"addOn"`
	License              AssetNameLicenseAttr            `json:"license" bson:"license"`
	Maintenance          AssetNameAssetStatusAttr        `json:"maintenance" bson:"maintenance"`
	Tracking             AssetNameAssetStatusAttr        `json:"tracking" bson:"tracking"`
	Audit                AssetNameAssetStatusAttr        `json:"audit" bson:"audit"`
	RepairTicketing      AssetNameAssetStatusAttr        `json:"repairTicketing" bson:"repairTicketing"`
	CreatedAt            time.Time                       `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time                       `json:"updatedAt" bson:"updatedAt"`
}

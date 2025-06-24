package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewAssetName() AssetName {
	activeStatus := "Active"
	return AssetName{
		Transactions: AssetNameTransactionAttr{
			Borrow:      false,
			Assign:      false,
			Transfer:    false,
			SelfService: false,
		},
		AddOn: &AssetNameAddOnAttr{
			RepairTicketing: false,
			Maintenance:     false,
			Tracking:        false,
			Audit:           false,
		},
		IsDeleted:   false,
		TotalAssets: 0,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
		TagType:     nil,
		AccountCode: nil,
		Measurement: nil,
		Category:    nil,
		License: AssetNameLicenseAttr{
			FixedAsset: AssetNameLicenseStatusAttr{
				Status:  activeStatus,
				AddedAt: nil,
			},
			Maintenance: AssetNameLicenseStatusAttr{
				Status:  activeStatus,
				AddedAt: nil,
			},
			Tracking: AssetNameLicenseStatusAttr{
				Status:  activeStatus,
				AddedAt: nil,
			},
			Audit: AssetNameLicenseStatusAttr{
				Status:  activeStatus,
				AddedAt: nil,
			},
			RepairTicketing: AssetNameLicenseStatusAttr{
				Status:  activeStatus,
				AddedAt: nil,
			},
		},
	}
}

type AssetNameEmbedAttr struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name"`
	Key              int                `json:"key" bson:"key"`
	Sequence         int                `json:"sequence" bson:"sequence"`
	NameWithSequence string             `json:"nameWithSequence" bson:"nameWithSequence"`
	AliasCode        string             `json:"aliasCode" bson:"aliasCode"`
}

type AssetNameTransactionAttr struct {
	Borrow      bool `json:"borrow"      bson:"borrow"`
	Assign      bool `json:"assign"      bson:"assign"`
	Transfer    bool `json:"transfer"    bson:"transfer"`
	SelfService bool `json:"selfService" bson:"selfService"`
}

type AssetNameCategoryAttr struct {
	ID   primitive.ObjectID `json:"_id"  bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int                `json:"key"  bson:"key"`
}

type AssetNameMeasurementAttr struct {
	ID   primitive.ObjectID `json:"_id"  bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type AssetNameAccountCodeAttr struct {
	ID   primitive.ObjectID `json:"_id"  bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Key  int                `json:"key"  bson:"key"`
	Code string             `json:"code" bson:"code"`
}

type AssetNameAddOnAttr struct {
	RepairTicketing bool `json:"repairTicketing" bson:"repairTicketing"`
	Maintenance     bool `json:"maintenance"     bson:"maintenance"`
	Tracking        bool `json:"tracking"        bson:"tracking"`
	Audit           bool `json:"audit"           bson:"audit"`
}

type AssetNameLicenseStatusAttr struct {
	Status  string     `json:"status"  bson:"status"`
	AddedAt *time.Time `json:"addedAt" bson:"addedAt"`
}

type AssetNameLicenseAttr struct {
	FixedAsset      AssetNameLicenseStatusAttr `json:"fixedAsset"      bson:"fixedAsset"`
	Maintenance     AssetNameLicenseStatusAttr `json:"maintenance"     bson:"maintenance"`
	Tracking        AssetNameLicenseStatusAttr `json:"tracking"        bson:"tracking"`
	Audit           AssetNameLicenseStatusAttr `json:"audit"           bson:"audit"`
	RepairTicketing AssetNameLicenseStatusAttr `json:"repairTicketing" bson:"repairTicketing"`
}

type AssetNameDepreciationGroupAttr = AssetNameCategoryAttr

type AssetName struct {
	ID                   primitive.ObjectID              `json:"_id,omitempty"        bson:"_id,omitempty"`
	Key                  int                             `json:"key"                  bson:"key"`
	Name                 string                          `json:"name"                 bson:"name"`
	AliasCode            string                          `json:"aliasCode"            bson:"aliasCode"`
	TotalAssets          int                             `json:"totalAssets"          bson:"totalAssets"`
	IsDeleted            bool                            `json:"isDeleted"            bson:"isDeleted"`
	TagType              *string                         `json:"tagType"              bson:"tagType"`
	DeleteFromDatabaseAt *time.Time                      `json:"deleteFromDatabaseAt" bson:"deleteFromDatabaseAt"`
	Transactions         AssetNameTransactionAttr        `json:"transactions"         bson:"transactions"`
	Category             *AssetNameCategoryAttr          `json:"category"             bson:"category"`
	Measurement          *AssetNameMeasurementAttr       `json:"measurement"          bson:"measurement"`
	DepreciationGroup    *AssetNameDepreciationGroupAttr `json:"depreciationGroup"    bson:"depreciationGroup"`
	AccountCode          *AssetNameAccountCodeAttr       `json:"accountCode"          bson:"accountCode"`
	AddOn                *AssetNameAddOnAttr             `json:"addOn"                bson:"addOn"`
	License              AssetNameLicenseAttr            `json:"license"              bson:"license"`
	CreatedAt            time.Time                       `json:"createdAt"            bson:"createdAt"`
	UpdatedAt            time.Time                       `json:"updatedAt"            bson:"updatedAt"`
	Code                 string                          `json:"code" bson:"code"`
}

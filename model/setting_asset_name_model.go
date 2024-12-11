package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewSettingAssetName() SettingAssetName {
	activeStatus := "Active"
	return SettingAssetName{
		Transactions: transactions{
			Borrow:      false,
			Assign:      false,
			Transfer:    false,
			SelfService: false,
		},
		AddOn: addOn{
			RepairTicketing: false,
			Maintenance:     false,
			Tracking:        false,
			Audit:           false,
		},
		IsDeleted:          false,
		SchemaVersion:      1,
		TotalAssets:        0,
		UpdatedAt:          time.Now(),
		CreatedAt:          time.Now(),
		DeleteFromDatabase: nil,
		TagType:            nil,
		AccountCode:        nil,
		Measurement:        nil,
		Category:           nil,
		License: license{
			FixedAsset: licenseDetail{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Maintenance: licenseDetail{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Tracking: licenseDetail{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Audit: licenseDetail{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			RepairTicketing: licenseDetail{
				Status:  &activeStatus,
				AddedAt: nil,
			},
		},
	}
}

type SettingAssetName struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	SchemaVersion      int                `bson:"schemaVersion" json:"schemaVersion"`
	Key                int                `bson:"key" json:"key"`
	Name               string             `bson:"name" json:"name"`
	AliasCode          *string            `bson:"aliasCode,omitempty" json:"aliasCode,omitempty"`
	Transactions       transactions       `bson:"transactions" json:"transactions"`
	Category           *category          `bson:"category,omitempty" json:"category,omitempty"`
	Measurement        *measurement       `bson:"measurement,omitempty" json:"measurement,omitempty"`
	DepreciationGroup  *DepreciationGroup `bson:"depreciationGroup,omitempty" json:"depreciationGroup,omitempty"`
	AccountCode        *accountCode       `bson:"accountCode,omitempty" json:"accountCode,omitempty"`
	TotalAssets        int                `bson:"totalAssets" json:"totalAssets"`
	AddOn              addOn              `bson:"addOn" json:"addOn"`
	IsDeleted          bool               `bson:"isDeleted" json:"isDeleted"`
	TagType            *string            `bson:"tagType,omitempty" json:"tagType,omitempty"`
	DeleteFromDatabase *time.Time         `bson:"deleteFromDatabaseAt,omitempty" json:"deleteFromDatabaseAt,omitempty"`
	License            license            `bson:"license" json:"license"`
	CreatedAt          time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type transactions struct {
	Borrow      bool `bson:"borrow" json:"borrow"`
	Assign      bool `bson:"assign" json:"assign"`
	Transfer    bool `bson:"transfer" json:"transfer"`
	SelfService bool `bson:"selfService" json:"selfService"`
}

type category struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
	Key  *int32              `bson:"key,omitempty" json:"key,omitempty"`
}

type measurement struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
}

type depreciationGroup struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UsefulLife *int                `bson:"usefulLife,omitempty" json:"usefulLife,omitempty"`
}

type accountCode struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
	Code *string             `bson:"code,omitempty" json:"code,omitempty"`
	Key  *int                `bson:"key,omitempty" json:"key,omitempty"`
}

type addOn struct {
	RepairTicketing bool `bson:"repairTicketing" json:"repairTicketing"`
	Maintenance     bool `bson:"maintenance" json:"maintenance"`
	Tracking        bool `bson:"tracking" json:"tracking"`
	Audit           bool `bson:"audit" json:"audit"`
}

type license struct {
	FixedAsset      licenseDetail `bson:"fixedAsset" json:"fixedAsset"`
	Maintenance     licenseDetail `bson:"maintenance" json:"maintenance"`
	Tracking        licenseDetail `bson:"tracking" json:"tracking"`
	Audit           licenseDetail `bson:"audit" json:"audit"`
	RepairTicketing licenseDetail `bson:"repairTicketing" json:"repairTicketing"`
}

type licenseDetail struct {
	Status  *string    `bson:"status,omitempty" json:"status,omitempty"`
	AddedAt *time.Time `bson:"addedAt,omitempty" json:"addedAt,omitempty"`
}

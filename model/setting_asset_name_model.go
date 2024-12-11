package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func NewSettingAssetName() SettingAssetName {
	activeStatus := "Active"
	return SettingAssetName{
		Transactions: TransactionsAttributeSettingAssetName{
			Borrow:      false,
			Assign:      false,
			Transfer:    false,
			SelfService: false,
		},
		AddOn: AddOnAttributeSettingAssetName{
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
		License: LicenseAttributeSettingAssetName{
			FixedAsset: LicenseDetailAttributeSettingAssetName{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Maintenance: LicenseDetailAttributeSettingAssetName{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Tracking: LicenseDetailAttributeSettingAssetName{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			Audit: LicenseDetailAttributeSettingAssetName{
				Status:  &activeStatus,
				AddedAt: nil,
			},
			RepairTicketing: LicenseDetailAttributeSettingAssetName{
				Status:  &activeStatus,
				AddedAt: nil,
			},
		},
	}
}

type SettingAssetName struct {
	ID                 primitive.ObjectID                          `bson:"_id,omitempty" json:"_id,omitempty"`
	SchemaVersion      int                                         `bson:"schemaVersion" json:"schemaVersion"`
	Key                int                                         `bson:"key" json:"key"`
	Name               string                                      `bson:"name" json:"name"`
	AliasCode          *string                                     `bson:"aliasCode,omitempty" json:"aliasCode,omitempty"`
	Transactions       TransactionsAttributeSettingAssetName       `bson:"transactions" json:"transactions"`
	Category           *CategoryAttributeSettingAssetName          `bson:"category,omitempty" json:"category,omitempty"`
	Measurement        *MeasurementAttributeSettingAssetName       `bson:"measurement,omitempty" json:"measurement,omitempty"`
	DepreciationGroup  *DepreciationGroupAttributeSettingAssetName `bson:"depreciationGroup,omitempty" json:"depreciationGroup,omitempty"`
	AccountCode        *AccountCodeAttributeSettingAssetName       `bson:"accountCode,omitempty" json:"accountCode,omitempty"`
	TotalAssets        int                                         `bson:"totalAssets" json:"totalAssets"`
	AddOn              AddOnAttributeSettingAssetName              `bson:"addOn" json:"addOn"`
	IsDeleted          bool                                        `bson:"isDeleted" json:"isDeleted"`
	TagType            *string                                     `bson:"tagType,omitempty" json:"tagType,omitempty"`
	DeleteFromDatabase *time.Time                                  `bson:"deleteFromDatabaseAt,omitempty" json:"deleteFromDatabaseAt,omitempty"`
	License            LicenseAttributeSettingAssetName            `bson:"license" json:"license"`
	CreatedAt          time.Time                                   `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAt          time.Time                                   `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type TransactionsAttributeSettingAssetName struct {
	Borrow      bool `bson:"borrow" json:"borrow"`
	Assign      bool `bson:"assign" json:"assign"`
	Transfer    bool `bson:"transfer" json:"transfer"`
	SelfService bool `bson:"selfService" json:"selfService"`
}

type CategoryAttributeSettingAssetName struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
	Key  *int32              `bson:"key,omitempty" json:"key,omitempty"`
}

type MeasurementAttributeSettingAssetName struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
}

type DepreciationGroupAttributeSettingAssetName struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UsefulLife *int                `bson:"usefulLife,omitempty" json:"usefulLife,omitempty"`
}

type AccountCodeAttributeSettingAssetName struct {
	ID   *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name *string             `bson:"name,omitempty" json:"name,omitempty"`
	Code *string             `bson:"code,omitempty" json:"code,omitempty"`
	Key  *int                `bson:"key,omitempty" json:"key,omitempty"`
}

type AddOnAttributeSettingAssetName struct {
	RepairTicketing bool `bson:"repairTicketing" json:"repairTicketing"`
	Maintenance     bool `bson:"maintenance" json:"maintenance"`
	Tracking        bool `bson:"tracking" json:"tracking"`
	Audit           bool `bson:"audit" json:"audit"`
}

type LicenseAttributeSettingAssetName struct {
	FixedAsset      LicenseDetailAttributeSettingAssetName `bson:"fixedAsset" json:"fixedAsset"`
	Maintenance     LicenseDetailAttributeSettingAssetName `bson:"maintenance" json:"maintenance"`
	Tracking        LicenseDetailAttributeSettingAssetName `bson:"tracking" json:"tracking"`
	Audit           LicenseDetailAttributeSettingAssetName `bson:"audit" json:"audit"`
	RepairTicketing LicenseDetailAttributeSettingAssetName `bson:"repairTicketing" json:"repairTicketing"`
}

type LicenseDetailAttributeSettingAssetName struct {
	Status  *string    `bson:"status,omitempty" json:"status,omitempty"`
	AddedAt *time.Time `bson:"addedAt,omitempty" json:"addedAt,omitempty"`
}

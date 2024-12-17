package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asset struct {
	ID      primitive.ObjectID `bson:"_id"                          json:"_id"`
	AssetID string             `bson:"assetId"                      json:"assetId"`
	Name    struct {
		ID               primitive.ObjectID `bson:"_id" json:"_id"`
		NameWithSequence string             `bson:"nameWithSequence" json:"nameWithSequence"`
		Name             string             `bson:"name" json:"name"`
		Key              int                `bson:"key" json:"key"`
		AliasCode        *string            `bson:"aliasCode,omitempty" json:"aliasCode,omitempty"`
	} `bson:"name"                         json:"name"`
	DepreciationGroup struct {
		ID         *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
		UsefulLife *int                `bson:"usefulLife,omitempty" json:"usefulLife,omitempty"`
	} `bson:"depreciationGroup"            json:"depreciationGroup"`
	AliasName          string              `bson:"aliasName,omitempty"          json:"aliasName,omitempty"`
	Measurement        *string             `bson:"measurement,omitempty"        json:"measurement,omitempty"`
	TagType            *string             `bson:"tagType,omitempty"            json:"tagType,omitempty"`
	QR                 AssetTag            `bson:"qr,omitempty"                 json:"qr,omitempty"`
	RFID               AssetTag            `bson:"rfid,omitempty"               json:"rfid,omitempty"`
	MaintenanceStatus  *string             `bson:"maintenanceStatus,omitempty"  json:"maintenanceStatus,omitempty"`
	AuditStatus        *string             `bson:"auditStatus,omitempty"        json:"auditStatus,omitempty"`
	Tag                AssetTag            `bson:"tag"                          json:"tag"`
	FirstImageMedium   *string             `bson:"firstImageMedium,omitempty"   json:"firstImageMedium,omitempty"`
	FirstImageBig      *string             `bson:"firstImageBig,omitempty"      json:"firstImageBig,omitempty"`
	FirstImageSmall    *string             `bson:"firstImageSmall,omitempty"    json:"firstImageSmall,omitempty"`
	SecondImageMedium  *string             `bson:"secondImageMedium,omitempty"  json:"secondImageMedium,omitempty"`
	SecondImageBig     *string             `bson:"secondImageBig,omitempty"     json:"secondImageBig,omitempty"`
	SecondImageSmall   *string             `bson:"secondImageSmall,omitempty"   json:"secondImageSmall,omitempty"`
	SetDefault         *string             `bson:"setDefault,omitempty"         json:"setDefault,omitempty"`
	Status             string              `bson:"status"                       json:"status"`
	PreviousStatus     *string             `bson:"previousStatus,omitempty"     json:"previousStatus,omitempty"`
	Category           AssetTreeAttribute  `bson:"category"                     json:"category"`
	Group              AssetTreeAttribute  `bson:"group"                        json:"group"`
	Brand              AttributeEmbedded   `bson:"brand"                        json:"brand"`
	Model              AttributeEmbedded   `bson:"model"                        json:"model"`
	AssignedTo         *UserInfo           `bson:"assignedTo,omitempty"         json:"assignedTo,omitempty"`
	Borrower           *UserInfo           `bson:"borrower,omitempty"           json:"borrower,omitempty"`
	Policy             *string             `bson:"policy,omitempty"             json:"policy,omitempty"`
	IsTransactionable  bool                `bson:"isTransactionable"            json:"isTransactionable"`
	IsParent           *bool               `bson:"isParent,omitempty"           json:"isParent,omitempty"`
	License            AssetLicense        `bson:"license"                      json:"license"`
	Transactions       AssetTransactions   `bson:"transactions"                 json:"transactions"`
	AddOn              AssetAddOn          `bson:"addOn"                        json:"addOn"`
	Purchase           AssetPurchase       `bson:"purchase"                     json:"purchase"`
	LastModifier       *UserInfo           `bson:"lastModifier,omitempty"       json:"lastModifier,omitempty"`
	Parent             *primitive.ObjectID `bson:"parent,omitempty"             json:"parent,omitempty"`
	NameNumber         *int                `bson:"nameNumber,omitempty"         json:"nameNumber,omitempty"`
	DeviceSerialNumber *string             `bson:"deviceSerialNumber,omitempty" json:"deviceSerialNumber,omitempty"`
	TagReported        *TagReported        `bson:"tagReported,omitempty"        json:"tagReported,omitempty"`
	TagReportedAt      *time.Time          `bson:"tagReportedAt,omitempty"      json:"tagReportedAt,omitempty"`
	TagReportedBy      *string             `bson:"tagReportedBy,omitempty"      json:"tagReportedBy,omitempty"`
	CreatedAt          time.Time           `bson:"createdAt"                    json:"createdAt"`
	UpdatedAt          time.Time           `bson:"updatedAt"                    json:"updatedAt"`
}

type AssetTag struct {
	Code      *string `bson:"code,omitempty"      json:"code,omitempty"`
	ShortCode *string `bson:"shortCode,omitempty" json:"shortCode,omitempty"`
}

type AssetTreeAttribute struct {
	ID       primitive.ObjectID `bson:"_id"                json:"_id"`
	Name     string             `bson:"name"               json:"name"`
	FullPath *string            `bson:"fullPath,omitempty" json:"fullPath,omitempty"`
	Key      int                `bson:"key"                json:"key"`
}

type UserInfo struct {
	ID       *primitive.ObjectID `bson:"_id,omitempty"      json:"_id,omitempty"`
	FullName *string             `bson:"fullName,omitempty" json:"fullName,omitempty"`
	Key      *int                `bson:"key,omitempty"      json:"key,omitempty"`
	Parent   *primitive.ObjectID `bson:"parent,omitempty"   json:"parent,omitempty"`
}

type AssetLicense struct {
	FixedAsset      LicenseDetails `bson:"fixedAsset"      json:"fixedAsset"`
	Maintenance     LicenseDetails `bson:"maintenance"     json:"maintenance"`
	Tracking        LicenseDetails `bson:"tracking"        json:"tracking"`
	Audit           LicenseDetails `bson:"audit"           json:"audit"`
	RepairTicketing LicenseDetails `bson:"repairTicketing" json:"repairTicketing"`
}

type LicenseDetails struct {
	Status  *string    `bson:"status,omitempty"  json:"status,omitempty"`
	AddedAt *time.Time `bson:"addedAt,omitempty" json:"addedAt,omitempty"`
}

type AssetTransactions struct {
	Borrowing  bool `bson:"borrowing"  json:"borrowing"`
	Assignment bool `bson:"assignment" json:"assignment"`
	Transfer   bool `bson:"transfer"   json:"transfer"`
	Disposal   bool `bson:"disposal"   json:"disposal"`
}

type AssetAddOn struct {
	RepairTicketing bool `bson:"repairTicketing" json:"repairTicketing"`
	Maintenance     bool `bson:"maintenance"     json:"maintenance"`
	Tracking        bool `bson:"tracking"        json:"tracking"`
	Audit           bool `bson:"audit"           json:"audit"`
}

type AssetPurchase struct {
	Date             *time.Time `bson:"date,omitempty"             json:"date,omitempty"`
	Cost             *float64   `bson:"cost,omitempty"             json:"cost,omitempty"`
	Currency         *string    `bson:"currency,omitempty"         json:"currency,omitempty"`
	WarrantyDuration *int       `bson:"warrantyDuration,omitempty" json:"warrantyDuration,omitempty"`
	WarrantyEndDate  *time.Time `bson:"warrantyEndDate,omitempty"  json:"warrantyEndDate,omitempty"`
}

type TagReported struct {
	RFID bool `bson:"rfid" json:"rfid"`
	QR   bool `bson:"qr"   json:"qr"`
}

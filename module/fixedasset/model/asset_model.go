package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/module/common/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Asset struct {
	ID                 primitive.ObjectID          `bson:"_id"                          json:"_id"`
	AssetID            string                      `bson:"assetId"                      json:"assetId"`
	AssetIDFull        string                      `bson:"assetIdFull"                      json:"assetIdFull"`
	Name               AssetNameAttribute          `bson:"name"                         json:"name"`
	DepreciationGroup  *DepreciationGroupAttribute `bson:"depreciationGroup"            json:"depreciationGroup"`
	AliasName          string                      `bson:"aliasName,omitempty"          json:"aliasName,omitempty"`
	Measurement        *string                     `bson:"measurement,omitempty"        json:"measurement,omitempty"`
	TagType            *string                     `bson:"tagType,omitempty"            json:"tagType,omitempty"`
	QR                 AssetTag                    `bson:"qr,omitempty"                 json:"qr,omitempty"`
	RFID               AssetTag                    `bson:"rfid,omitempty"               json:"rfid,omitempty"`
	MaintenanceStatus  *string                     `bson:"maintenanceStatus,omitempty"  json:"maintenanceStatus,omitempty"`
	AuditStatus        *string                     `bson:"auditStatus,omitempty"        json:"auditStatus,omitempty"`
	Tag                AssetTag                    `bson:"tag"                          json:"tag"`
	FirstImageMedium   *string                     `bson:"firstImageMedium"   json:"firstImageMedium"`
	FirstImageBig      *string                     `bson:"firstImageBig"      json:"firstImageBig"`
	FirstImageSmall    *string                     `bson:"firstImageSmall"    json:"firstImageSmall"`
	SecondImageMedium  *string                     `bson:"secondImageMedium"  json:"secondImageMedium"`
	SecondImageBig     *string                     `bson:"secondImageBig"     json:"secondImageBig"`
	SecondImageSmall   *string                     `bson:"secondImageSmall"   json:"secondImageSmall"`
	DefaultImageMedium *string                     `bson:"defaultImageMedium"  json:"defaultImageMedium"`
	DefaultImageBig    *string                     `bson:"defaultImageBig"     json:"defaultImageBig"`
	DefaultImageSmall  *string                     `bson:"defaultImageSmall"   json:"defaultImageSmall"`
	SetDefault         *string                     `bson:"setDefault"         json:"setDefault"`
	Status             string                      `bson:"status"                       json:"status"`
	PreviousStatus     *string                     `bson:"previousStatus,omitempty"     json:"previousStatus,omitempty"`
	Category           AssetTreeAttribute          `bson:"category"                     json:"category"`
	Group              AssetTreeAttribute          `bson:"group"                        json:"group"`
	Brand              model.AttributeEmbedded     `bson:"brand"                        json:"brand"`
	Model              model.AttributeEmbedded     `bson:"model"                        json:"model"`
	AssignedTo         *UserInfo                   `bson:"assignedTo,omitempty"         json:"assignedTo,omitempty"`
	Borrower           *UserInfo                   `bson:"borrower,omitempty"           json:"borrower,omitempty"`
	Policy             *string                     `bson:"policy,omitempty"             json:"policy,omitempty"`
	IsTransactionable  bool                        `bson:"isTransactionable"            json:"isTransactionable"`
	IsParent           bool                        `bson:"isParent"           json:"isParent"`
	License            AssetLicense                `bson:"license"                      json:"license"`
	Transactions       AssetTransactions           `bson:"transactions"                 json:"transactions"`
	AddOn              AssetAddOn                  `bson:"addOn"                        json:"addOn"`
	Purchase           AssetPurchase               `bson:"purchase"                     json:"purchase"`
	LastModifier       *UserInfo                   `bson:"lastModifier,omitempty"       json:"lastModifier,omitempty"`
	Parent             *primitive.ObjectID         `bson:"parent"             json:"parent"`
	NameNumber         *int                        `bson:"nameNumber,omitempty"         json:"nameNumber,omitempty"`
	DeviceSerialNumber *string                     `bson:"deviceSerialNumber,omitempty" json:"deviceSerialNumber,omitempty"`
	TagReported        *TagReported                `bson:"tagReported,omitempty"        json:"tagReported,omitempty"`
	TagReportedAt      *time.Time                  `bson:"tagReportedAt,omitempty"      json:"tagReportedAt,omitempty"`
	TagReportedBy      *string                     `bson:"tagReportedBy,omitempty"      json:"tagReportedBy,omitempty"`
	CustomFields       []CustomField               `bson:"customFields"                    json:"customFields"`
	CreatedAt          time.Time                   `bson:"createdAt"                    json:"createdAt"`
	UpdatedAt          time.Time                   `bson:"updatedAt"                    json:"updatedAt"`
}

type AssetTag struct {
	Code      *string `bson:"code"      json:"code"`
	ShortCode *string `bson:"shortCode" json:"shortCode"`
}

type AssetTreeAttribute struct {
	ID       primitive.ObjectID `bson:"_id"                json:"_id"`
	Name     string             `bson:"name"               json:"name"`
	FullPath *string            `bson:"fullPath,omitempty" json:"fullPath,omitempty"`
	Key      int                `bson:"key"                json:"key"`
}
type DepreciationGroupAttribute struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	UsefulLife int                `bson:"usefulLife" json:"usefulLife"`
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
	Date             *time.Time `bson:"date"             json:"date"`
	Cost             *float64   `bson:"cost"             json:"cost"`
	Currency         *string    `bson:"currency"         json:"currency"`
	WarrantyDuration *int       `bson:"warrantyDuration" json:"warrantyDuration"`
	WarrantyEndDate  *time.Time `bson:"warrantyEndDate"  json:"warrantyEndDate"`
}

type TagReported struct {
	RFID bool `bson:"rfid" json:"rfid"`
	QR   bool `bson:"qr"   json:"qr"`
}

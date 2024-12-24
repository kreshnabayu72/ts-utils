package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserEmbedded struct {
	ID        *primitive.ObjectID `json:"_id,omitempty"       bson:"_id,omitempty"`
	FullName  *string             `json:"fullName"            bson:"fullName"`
	Key       *int32              `json:"key"                 bson:"key"`
	Position  *string             `json:"position,omitempty"  bson:"position,omitempty"`
	Division  *string             `json:"division,omitempty"  bson:"division,omitempty"`
	IsDefault *bool               `json:"isDefault,omitempty" bson:"isDefault,omitempty"`
}
type AttributeEmbedded struct {
	ID   *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name *string             `json:"name"          bson:"name"`
	Key  *int32              `json:"key"           bson:"key"`
}

type NameEmbedded struct {
	ID               primitive.ObjectID `bson:"_id"              json:"_id"`
	NameWithSequence string             `bson:"nameWithSequence" json:"nameWithSequence"`
	Name             string             `bson:"name"             json:"name"`
	Key              int                `bson:"key"              json:"key"`
	NameNumber       *int               `bson:"nameNumber"       json:"nameNumber"`
}

type RoleType string

const (
	RTImportData          RoleType = "importData"
	RTAssetAttribute      RoleType = "assetAttributes"
	RTAssetPolicies       RoleType = "assetPolicies"
	RTDepreciationGroup   RoleType = "depreciationGroup"
	RTDepreciationMethod  RoleType = "depreciationMethod"
	RTDocumentDeletion    RoleType = "documentDeletion"
	RTGroup               RoleType = "group"
	RTUser                RoleType = "user"
	RTIotReader           RoleType = "iotReader"
	RTLicense             RoleType = "license"
	RTPurchaseDocument    RoleType = "purchaseDocument"
	RTPurchaseInformation RoleType = "purchaseInformation"
	RTRole                RoleType = "role"
	RTTag                 RoleType = "tag"
	RTRegisterAsset       RoleType = "registerAsset"
	RTTagType             RoleType = "tagType"
	RTOpenApi             RoleType = "openApi"
)

type RoleName string

var RoleNameMap = map[RoleType]RoleName{
	RTImportData:          "Import Data",
	RTAssetAttribute:      "Manage asset names, brands, categories, and custom fields",
	RTAssetPolicies:       "Manage asset policy",
	RTDepreciationGroup:   "Manage depreciation group",
	RTDepreciationMethod:  "Manage depreciation method",
	RTDocumentDeletion:    "Manage document deletion",
	RTGroup:               "Manage group",
	RTUser:                "Manage user",
	RTIotReader:           "Manage IOT and reader",
	RTLicense:             "Manage license",
	RTPurchaseDocument:    "Manage purchase document",
	RTPurchaseInformation: "Manage purchase information",
	RTRole:                "Manage role",
	RTTag:                 "Manage TAG",
	RTRegisterAsset:       "Register assets",
	RTTagType:             "Manage TAG type",
	RTOpenApi:             "Open API",
}

type Permission struct {
	Create bool `json:"create" bson:"create"`
	View   bool `json:"view"   bson:"view"`
	Update bool `json:"update" bson:"update"`
	Delete bool `json:"delete" bson:"delete"`
}

type ManagePermission struct {
	ImportData          Permission `json:"importData"          bson:"importData"`
	AssetAttribute      Permission `json:"assetAttribute"      bson:"assetAttribute"`
	AssetPolicies       Permission `json:"assetPolicies"       bson:"assetPolicies"`
	DepreciationGroup   Permission `json:"depreciationGroup"   bson:"depreciationGroup"`
	DepreciationMethod  Permission `json:"depreciationMethod"  bson:"depreciationMethod"`
	DocumentDeletion    Permission `json:"documentDeletion"    bson:"documentDeletion"`
	Group               Permission `json:"group"               bson:"group"`
	User                Permission `json:"user"                bson:"user"`
	IotReader           Permission `json:"iotReader"           bson:"iotReader"`
	License             Permission `json:"license"             bson:"license"`
	PurchaseDocument    Permission `json:"purchaseDocument"    bson:"purchaseDocument"`
	PurchaseInformation Permission `json:"purchaseInformation" bson:"purchaseInformation"`
	Role                Permission `json:"role"                bson:"role"`
	Tag                 Permission `json:"tag"                 bson:"tag"`
	RegisterAsset       Permission `json:"registerAsset"       bson:"registerAsset"`
	TagType             Permission `json:"tagType"             bson:"tagType"`
	OpenApi             Permission `json:"openApi"             bson:"openApi"`
}

type TransactionType string

const (
	TTBorrowingAssignment TransactionType = "Borrowing/Assignment"
	TTDisposal            TransactionType = "Disposal"
	TTTransfer            TransactionType = "Tranfer"
	TTAudit               TransactionType = "Audit"
	TTMaintenanceRoutine  TransactionType = "Maintenance Routine"
	TTRepairTicketing     TransactionType = "Repair Ticketing"
	TTTracking            TransactionType = "Tracking"
	TTDamaged             TransactionType = "Damaged"
	TTMissing             TransactionType = "Missing"
)

type TransactionGroupAttribute string

const (
	TGABorrowingRole   TransactionGroupAttribute = "borrowingRole"
	TGADisposalRole    TransactionGroupAttribute = "disposalRole"
	TGATransferRole    TransactionGroupAttribute = "transferRole"
	TGAAuditRole       TransactionGroupAttribute = "auditRole"
	TGAMaintenanceRole TransactionGroupAttribute = "maintenanceRole"
	TGARepairRole      TransactionGroupAttribute = "repairRole"
	TGATrackingRole    TransactionGroupAttribute = "trackingRole"
	TGADamagedRole     TransactionGroupAttribute = "damagedRole"
	TGAMissingRole     TransactionGroupAttribute = "missingRole"
)

type TransactionRoleType string

const (
	TRTManager    TransactionRoleType = "Manager"
	TRTMonitoring TransactionRoleType = "Monitoring"
	TRTApproval   TransactionRoleType = "Approval"
	TRTStaff      TransactionRoleType = "Staff"
)

type ApprovalType string

const (
	ApprovalTypeOr  ApprovalType = "Or"
	ApprovalTypeAnd ApprovalType = "And"
)

type TagType string

const (
	TagTypeNonTag TagType = "Non TAG"
	TagTypeQR     TagType = "QR"
	TagTypeRFID   TagType = "RFID"
	TagTypeRFIDQR TagType = "RFID & QR"
)

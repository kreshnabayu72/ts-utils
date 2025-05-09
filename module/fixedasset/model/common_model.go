package model

type RoleType string

const (
	RTImportData          RoleType = "importData"
	RTAssetAttribute      RoleType = "assetAttribute"
	RTAssetPolicies       RoleType = "assetPolicies"
	RTDepreciationGroup   RoleType = "depreciationGroup"
	RTDepreciationMethod  RoleType = "depreciationMethod"
	RTDocumentDeletion    RoleType = "documentDeletion"
	RTGroup               RoleType = "group"
	RTIotReader           RoleType = "iotReader"
	RTLicense             RoleType = "license"
	RTPurchaseDocument    RoleType = "purchaseDocument"
	RTPurchaseInformation RoleType = "purchaseInformation"
	RTUserAndRole         RoleType = "userAndRole"
	RTTag                 RoleType = "tag"
	RTRegisterAsset       RoleType = "registerAsset"
	RTTagType             RoleType = "tagType"
	RTOpenApi             RoleType = "openApi"
)

var RoleTypeMap = map[string]RoleType{
	"importData":          RTImportData,
	"assetAttribute":      RTAssetAttribute,
	"assetPolicies":       RTAssetPolicies,
	"depreciationGroup":   RTDepreciationGroup,
	"depreciationMethod":  RTDepreciationMethod,
	"documentDeletion":    RTDocumentDeletion,
	"group":               RTGroup,
	"iotReader":           RTIotReader,
	"purchaseDocument":    RTPurchaseDocument,
	"purchaseInformation": RTPurchaseInformation,
	"userAndRole":         RTUserAndRole,
	"tag":                 RTTag,
	"registerAsset":       RTRegisterAsset,
	"tagType":             RTTagType,
	"openApi":             RTOpenApi,
}

type RoleName string

var RoleNameMap = map[RoleType]RoleName{
	RTImportData:          "Import Data",
	RTAssetAttribute:      "Manage asset names, brands, categories, and custom fields",
	RTAssetPolicies:       "Manage asset policy",
	RTDepreciationGroup:   "Manage depreciation group",
	RTDepreciationMethod:  "Manage depreciation method",
	RTDocumentDeletion:    "Manage document deletion",
	RTGroup:               "Manage group",
	RTUserAndRole:         "Manage user and role",
	RTIotReader:           "Manage IOT and reader",
	RTLicense:             "Manage license",
	RTPurchaseDocument:    "Manage purchase document",
	RTPurchaseInformation: "Manage purchase information",
	RTTag:                 "Manage TAG",
	RTRegisterAsset:       "Register assets",
	RTTagType:             "Manage TAG type",
	RTOpenApi:             "Open API",
}

// Group/Global
func GetRoleType(role RoleType) string {
	switch role {
	case RTImportData, RTAssetAttribute, RTAssetPolicies, RTTagType, RTDepreciationGroup, RTDepreciationMethod, RTOpenApi:
		return "Global"
	default:
		return "Group"
	}
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

// TransactionManageType represents the transaction role's manage type
type TransactionManageType string

const (
	TransactionManageTypeManager    TransactionManageType = "Manager"
	TransactionManageTypeMonitoring TransactionManageType = "Monitoring"
	TransactionManageTypeStaff      TransactionManageType = "Staff"
	TransactionManageTypeApprover   TransactionManageType = "Approval"
)

type PermissionType string

const (
	PermissionCreate PermissionType = "create"
	PermissionUpdate PermissionType = "update"
	PermissionDelete PermissionType = "delete"
	PermissionView   PermissionType = "view"
)

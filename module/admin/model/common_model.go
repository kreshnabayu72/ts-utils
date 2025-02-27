package model

type RoleType string

const (
	// Global
	RTDeviceAttributes         RoleType = "deviceAttributes"
	RTAddTAGStock              RoleType = "addTagStock"
	RTAddHardwareStock         RoleType = "addHardwareStock"
	RTManageAllocationApproval RoleType = "manageAllocationApproval"
	RTVerifyTAGAllocation      RoleType = "verifyTAGAllocation"
	RTVerifyHardwareAllocation RoleType = "verifyHardwareAllocation"
	RTManageLanguage           RoleType = "manageLanguage"
	RTManagePreload            RoleType = "managePreload"

	// Group
	RTManagePartner                  RoleType = "managePartner"
	RTManageGroup                    RoleType = "manageGroup"
	RTManageUserAndRole              RoleType = "manageUserAndRole"
	RTManageActiveClient             RoleType = "manageActiveClient"
	RTManageExpiredAndArchivedClient RoleType = "manageExpiredAndArchivedClient"
	RTManageAllocation               RoleType = "manageAllocation"
)

type RoleName string

var RoleNameMap = map[RoleType]RoleName{
	// Global
	RTDeviceAttributes:         "Manage category, device name, brand, & SKU",
	RTAddTAGStock:              "Add TAG stock",
	RTAddHardwareStock:         "Add Hardware stock",
	RTManageAllocationApproval: "Manage allocation approval",
	RTVerifyTAGAllocation:      "Verify TAG Allocation",
	RTVerifyHardwareAllocation: "Verify Hardware Allocation",
	RTManageLanguage:           "Manage language",
	RTManagePreload:            "Manage preload",

	// Group
	RTManagePartner:                  "Manage partner",
	RTManageGroup:                    "Manage group",
	RTManageUserAndRole:              "Manage user and role",
	RTManageActiveClient:             "Manage active client",
	RTManageExpiredAndArchivedClient: "Manage expired & archived client",
	RTManageAllocation:               "Manage allocation",
}

// Group/Global
func GetRoleType(role RoleType) string {
	switch role {
	case RTDeviceAttributes, RTAddTAGStock, RTAddHardwareStock, RTManageAllocationApproval,
		RTVerifyTAGAllocation, RTVerifyHardwareAllocation, RTManageLanguage, RTManagePreload:
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

package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SystemRoleAccessType represents the type of the role that has access to
type SystemRoleAccessType string

const (
	SystemRoleAccessTypeGlobal SystemRoleAccessType = "Global"
	SystemRoleAccessTypeGroup  SystemRoleAccessType = "Group"
)

// SystemRoleType represents the possible types of system role
type SystemRoleType string

// Global System Role
const (
	SystemRoleTypeImportData    SystemRoleType = "importData"
	SystemRoleTypeItemAttribute SystemRoleType = "itemAttribute"
	SystemRoleTypeItemPolicy    SystemRoleType = "itemPolicy"
	SystemRoleTypeOpenAPI       SystemRoleType = "openApi"
)

// Group System Role
const (
	SystemRoleTypeDocumentDeletion SystemRoleType = "documentDeletion"
	SystemRoleTypeGroup            SystemRoleType = "group"
	SystemRoleUserRole             SystemRoleType = "userRole"
	SystemRoleIOTReader            SystemRoleType = "iotReader"
	SystemRoleLicense              SystemRoleType = "license"
	SystemRoleTAG                  SystemRoleType = "tag"
	SystemRoleItemStock            SystemRoleType = "itemStock"
)

// GetAccessType gets the SystemRoleAccessType of the SystemRoleType
func (roleType SystemRoleType) GetAccessType() SystemRoleAccessType {
	switch roleType {
	case SystemRoleTypeImportData, SystemRoleTypeItemAttribute, SystemRoleTypeItemPolicy, SystemRoleTypeOpenAPI:
		return SystemRoleAccessTypeGlobal
	default:
		return SystemRoleAccessTypeGroup
	}
}

// SystemRoleName represents the name of the system role's SystemRoleType
type SystemRoleName string

var SystemRoleNameMap = map[SystemRoleType]SystemRoleName{
	SystemRoleTypeImportData:       "Import data",
	SystemRoleTypeItemAttribute:    "Manage item names, item SKUs, brands, categories and custom fields",
	SystemRoleTypeItemPolicy:       "Manage item policy",
	SystemRoleTypeOpenAPI:          "Open API",
	SystemRoleTypeDocumentDeletion: "Manage document deletion",
	SystemRoleTypeGroup:            "Manage group",
	SystemRoleUserRole:             "Manage user and role",
	SystemRoleIOTReader:            "Manage IOT and reader",
	SystemRoleLicense:              "Manage license",
	SystemRoleTAG:                  "Manage TAG",
	SystemRoleItemStock:            "Manage items and stock",
}

// SystemRolePermission represents the CRUD permission of the system role
type SystemRolePermission struct {
	Create bool `json:"create" bson:"create"`
	View   bool `json:"view"   bson:"view"`
	Update bool `json:"update" bson:"update"`
	Delete bool `json:"delete" bson:"delete"`
}

// UserSystemRole represents the user's individual system role data
type UserSystemRole struct {
	ID                  primitive.ObjectID   `json:"_id,omitempty"       bson:"_id,omitempty"`
	User                primitive.ObjectID   `json:"user"                bson:"user"`
	IsActive            bool                 `json:"isActive"            bson:"isActive"`
	IsTotalControl      bool                 `json:"isTotalControl"      bson:"isTotalControl"`
	IsReadOnly          bool                 `json:"isReadOnly"          bson:"isReadOnly"`
	AdministrativeGroup primitive.ObjectID   `json:"administrativeGroup" bson:"administrativeGroup"`
	RoleName            SystemRoleName       `json:"roleName"            bson:"roleName"`
	RoleType            SystemRoleType       `json:"roleType"            bson:"roleType"`
	AccessType          SystemRoleAccessType `json:"accessType"          bson:"accessType"`
	Permissions         SystemRolePermission `json:"permissions"         bson:"permissions"`
	IsActivated         bool                 `json:"isActivated"         bson:"isActivated"`
	UpdatedAt           time.Time            `json:"updatedAt"           bson:"updatedAt"`
	CreatedAt           time.Time            `json:"createdAt"           bson:"createdAt"`
}

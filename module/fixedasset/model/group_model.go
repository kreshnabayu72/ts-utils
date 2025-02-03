package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Approval represents the approval schema in MongoDB.
type Approval struct {
	Level        int    `bson:"level"        json:"level"`
	ApprovalType string `bson:"approvalType" json:"approvalType"` // "Or" or "And"
}

// GroupRole represents the group role schema.
type GroupRole struct {
	IsManagedByParent bool       `bson:"isManagedByParent" json:"isManagedByParent"`
	Approvals         []Approval `bson:"approvals"         json:"approvals"`
}

// AddOnData represents the add-on data schema.
type AddOnData struct {
	Total                *int `bson:"total"                json:"total,omitempty"`
	Used                 int  `bson:"used"                 json:"used"`
	UsedByThisGroup      *int `bson:"usedByThisGroup"      json:"usedByThisGroup,omitempty"`
	AllocatedToSubGroups int  `bson:"allocatedToSubGroups" json:"allocatedToSubGroups"`
	Available            *int `bson:"available"            json:"available,omitempty"`
}

// GroupLicense represents the license schema for a group.
type GroupLicense struct {
	AvailableNew     *int `bson:"availableNew"     json:"availableNew,omitempty"`
	AvailableExpired *int `bson:"availableExpired" json:"availableExpired,omitempty"`
	UsedActive       *int `bson:"usedActive"       json:"usedActive,omitempty"`
	UsedExpired      *int `bson:"usedExpired"      json:"usedExpired,omitempty"`
	Total            *int `bson:"total"            json:"total,omitempty"`
}

// Quota represents the quota information for various categories.
type Quota struct {
	RFID            AddOnData `bson:"rfid"            json:"rfid"`
	FixedAsset      AddOnData `bson:"fixedAsset"      json:"fixedAsset"`
	RepairTicketing AddOnData `bson:"repairTicketing" json:"repairTicketing"`
	Maintenance     AddOnData `bson:"maintenance"     json:"maintenance"`
	Audit           AddOnData `bson:"audit"           json:"audit"`
	Tracking        AddOnData `bson:"tracking"        json:"tracking"`
}

// License represents the license information for various categories.
type License struct {
	RFID            GroupLicense `bson:"rfid"            json:"rfid"`
	FixedAsset      GroupLicense `bson:"fixedAsset"      json:"fixedAsset"`
	Audit           GroupLicense `bson:"audit"           json:"audit"`
	Tracking        GroupLicense `bson:"tracking"        json:"tracking"`
	Maintenance     GroupLicense `bson:"maintenance"     json:"maintenance"`
	RepairTicketing GroupLicense `bson:"repairTicketing" json:"repairTicketing"`
}

// IGroup represents the main group schema.
type Group struct {
	ID                     primitive.ObjectID  `bson:"_id,omitempty"                    json:"_id,omitempty"`
	Name                   string              `bson:"name"                             json:"name"`
	FullPath               *string             `bson:"fullPath,omitempty"               json:"fullPath,omitempty"`
	TotalAssets            int                 `bson:"totalAssets"                      json:"totalAssets"`
	TotalUsers             int                 `bson:"totalUsers"                       json:"totalUsers"`
	Desc                   *string             `bson:"desc,omitempty"                   json:"desc,omitempty"`
	Country                *string             `bson:"country,omitempty"                json:"country,omitempty"`
	State                  *string             `bson:"state,omitempty"                  json:"state,omitempty"`
	City                   *string             `bson:"city,omitempty"                   json:"city,omitempty"`
	Zip                    *string             `bson:"zip,omitempty"                    json:"zip,omitempty"`
	Address                *string             `bson:"address,omitempty"                json:"address,omitempty"`
	Quota                  Quota               `bson:"quota"                            json:"quota"`
	ParentQuota            Quota               `bson:"parentQuota"                      json:"parentQuota"`
	ParentQuotaSourceID    *primitive.ObjectID `bson:"parentQuotaSourceId,omitempty"    json:"parentQuotaSourceId,omitempty"`
	ParentQuotaSourceLevel *int                `bson:"parentQuotaSourceLevel,omitempty" json:"parentQuotaSourceLevel,omitempty"`
	Level                  *int                `bson:"level,omitempty"                  json:"level,omitempty"`
	License                License             `bson:"license"                          json:"license"`
	IsDisposable           bool                `bson:"isDisposable"                     json:"isDisposable"`
	Parent                 *primitive.ObjectID `bson:"parent,omitempty"                 json:"parent,omitempty"`
	BorrowingRole          GroupRole           `bson:"borrowingRole"                    json:"borrowingRole"`
	TransferRole           GroupRole           `bson:"transferRole"                     json:"transferRole"`
	DisposalRole           GroupRole           `bson:"disposalRole"                     json:"disposalRole"`
	MaintenanceRole        GroupRole           `bson:"maintenanceRole"                  json:"maintenanceRole"`
	RepairRole             GroupRole           `bson:"repairRole"                       json:"repairRole"`
	AuditRole              GroupRole           `bson:"auditRole"                        json:"auditRole"`
	TrackingRole           GroupRole           `bson:"trackingRole"                     json:"trackingRole"`
	DamagedRole            GroupRole           `bson:"damagedRole"                      json:"damagedRole"`
	MissingRole            GroupRole           `bson:"missingRole"                      json:"missingRole"`
	IsDeleted              bool                `bson:"isDeleted"                        json:"isDeleted"`
	DeleteFromDatabaseAt   *time.Time          `bson:"deleteFromDatabaseAt,omitempty"   json:"deleteFromDatabaseAt,omitempty"`
	Key                    int                 `bson:"key"                              json:"key"`
	CreatedAt              time.Time           `bson:"createdAt"                        json:"createdAt"`
	UpdatedAt              time.Time           `bson:"updatedAt"                        json:"updatedAt"`
}

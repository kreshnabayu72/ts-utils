// Package model defines the analytics-reporting-service database row types.
package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// model name = "change_logs" — shared by GS Changelog/System Log/Activity Log
// and AC Activity Log > User Action (prd-conventions.md §8.3: three surfaces,
// one collection, one transaction). Product tells the two apart.
//
// Object stays a plain string: the PRD's module list omits Division and Position,
// both of which GS flows already write, so a closed enum would be wrong today.
//
// ObjectId is the id of the record the action changed. Not in the GS PRD, which
// leaves User Detail > Changelog filtering on Name — a snapshot, so a rename
// splits the trail and two users sharing a name collide. Admin Console's copy of
// this collection carries entity for the same reason. Nullable: some
// rows are settings-level (General Settings) and have no target record.
type Changelog struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Action   ChangeLogAction    `json:"action" bson:"action"`
	Field    *string            `json:"field" bson:"field"`
	OldValue *string            `json:"oldValue" bson:"oldValue"`
	NewValue *string            `json:"newValue" bson:"newValue"`

	// Used for filtering which changelog went to which log table in FE
	// Cant use 'object' since sometimes >1 object is shown in 1 same log table
	LogTable string `json:"logTable" bson:"logTable"`

	// changelog's creator
	CreatedById       string `json:"createdById" bson:"createdById"`
	CreatedBySnapshot string `json:"createdBySnapshot" bson:"createdBySnapshot"`

	// changelog's entity (client/partner/distributor)
	EntityID   string     `json:"entityId" bson:"entityId"`
	EntityType EntityType `json:"entityType" bson:"entityType"`

	// Product disambiguates rows written by GS vs AC into this shared collection
	// (nil = GS, same convention as SessionLog.Product). Without it, an AC row
	// ("entity_type='admin_console'" in the AC PRD) and a GS row using EntityType
	// for the tenant would be indistinguishable by that field alone.
	Product *Product `json:"product" bson:"product"`

	// changelog's referenced object
	Object     string      `json:"object" bson:"object"`         // shown on FE
	ObjectType *ObjectType `json:"objectType" bson:"objectType"` // AC's target_entity_type; nil for GS rows, which don't scope by it
	ObjectId   *string     `json:"objectId" bson:"objectId"`     // for filtering
	ObjectName *string     `json:"objectName" bson:"objectName"` // nil for settings-level rows (no named entity) — prd-conventions.md §9.10

	FileManagerType string `json:"fileManagerType"  bson:"fileManagerType"`

	GroupId string `json:"groupId" bson:"groupId"` // for role filtering

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// EntityType is the tenant kind that owns a log row. Shared by change_logs and
// session_logs — it is a tenant concept, not a per-collection one, so it is
// declared once rather than model-prefixed like the other log enums.
//
// It decides the retention window: Client plans purge at 7d/30d/90d/1y,
// Principal/Distributor/Partner never do. Snapshotted at write time like the
// identity fields — resolving it from the tenant at purge time would mean a
// lookup per row.
type EntityType string

const (
	EntityTypeClient      EntityType = "client"
	EntityTypePrincipal   EntityType = "principal"
	EntityTypeDistributor EntityType = "distributor"
	EntityTypePartner     EntityType = "partner"
)

// ObjectType is AC's target_entity_type (Admin-Console/PRD/activity-log/04-data.md
// §1.1) — what kind of record ObjectId points at, needed to scope AC's entity
// filter. GS rows leave this nil: the GS PRD never specs a separate type field,
// only the display Object path, so Object alone is scope enough there.
type ObjectType string

const (
	ObjectTypeDistributor   ObjectType = "distributor"
	ObjectTypePartner       ObjectType = "partner"
	ObjectTypeClient        ObjectType = "client"
	ObjectTypeGroup         ObjectType = "group"
	ObjectTypeUser          ObjectType = "user"
	ObjectTypeRole          ObjectType = "role"
	ObjectTypeDeviceName    ObjectType = "device_name"
	ObjectTypeBrand         ObjectType = "brand"
	ObjectTypeSKU           ObjectType = "sku"
	ObjectTypeCategory      ObjectType = "category"
	ObjectTypeMeasurement   ObjectType = "measurement"
	ObjectTypeLanguage      ObjectType = "language"
	ObjectTypeDataDeletion  ObjectType = "data_deletion"
	ObjectTypeLicenseTiming ObjectType = "license_timing"
	ObjectTypeSubscription  ObjectType = "subscription"
	ObjectTypeAllocation    ObjectType = "allocation"
)

// ChangeLogAction is what was done to the record. GS activity-log/04-data.md
// §4.1 names Create/Edit/Delete explicitly and then says "aksi lain (Activate,
// Deactivate, Assign, dll.)" — the set is open by design, so these constants
// don't claim to be exhaustive. Title Case throughout: this is the shared
// collection's canonical casing (prd-conventions.md §6.1 plain-text Action
// column) — AC's own PRD spec lists lower_snake_case values (`'top_up'`,
// `'resend_activation_email'`), which must be normalized to Title Case at
// write time so one filter dropdown doesn't render two casings of "Create".
type ChangeLogAction string

const (
	ChangeLogActionCreate                ChangeLogAction = "Create"
	ChangeLogActionEdit                  ChangeLogAction = "Edit"
	ChangeLogActionDelete                ChangeLogAction = "Delete"
	ChangeLogActionActivate              ChangeLogAction = "Activate"
	ChangeLogActionDeactivate            ChangeLogAction = "Deactivate"
	ChangeLogActionAssign                ChangeLogAction = "Assign"
	ChangeLogActionRemove                ChangeLogAction = "Remove"
	ChangeLogActionReassign              ChangeLogAction = "Reassign"
	ChangeLogActionImport                ChangeLogAction = "Import"
	ChangeLogActionCombine               ChangeLogAction = "Combine"
	ChangeLogActionSeparate              ChangeLogAction = "Separate"
	ChangeLogActionAudit                 ChangeLogAction = "Audit"
	ChangeLogActionMarkAsAvailable       ChangeLogAction = "Mark as Available"
	ChangeLogActionMarkAsDamaged         ChangeLogAction = "Mark as Damaged"
	ChangeLogActionMarkAsMissing         ChangeLogAction = "Mark as Missing"
	ChangeLogActionAssignUser            ChangeLogAction = "Assign User"
	ChangeLogActionRemoveUser            ChangeLogAction = "Remove User"
	ChangeLogActionEditPermission        ChangeLogAction = "Edit Permission"
	ChangeLogActionPingReader            ChangeLogAction = "Ping Reader"
	ChangeLogActionConnectTablet         ChangeLogAction = "Connect Tablet"
	ChangeLogActionDisconnectTablet      ChangeLogAction = "Disconnect Tablet"
	ChangeLogActionPairTag               ChangeLogAction = "Pair TAG"
	ChangeLogActionUnpairTag             ChangeLogAction = "Unpair TAG"
	ChangeLogActionResendActivationEmail ChangeLogAction = "Resend Activation Email"
	ChangeLogActionCancel                ChangeLogAction = "Cancel"
	ChangeLogActionTopUp                 ChangeLogAction = "Top Up"
	ChangeLogActionUpgrade               ChangeLogAction = "Upgrade"
	ChangeLogActionDowngrade             ChangeLogAction = "Downgrade"
	ChangeLogActionRenew                 ChangeLogAction = "Renew"
)

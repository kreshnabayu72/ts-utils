package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportType string

const (
	ReportTypeCustomReport                ReportType = "Custom Report"
	ReportTypeSummaryOverview             ReportType = "Summary Overview"
	ReportTypeDepreciationReport          ReportType = "Depreciation Report"
	ReportTypeMaintenanceRoutineReport    ReportType = "Maintenance Routine Report"
	ReportTypeRepairmentReport            ReportType = "Repairment Report"
	ReportTypeBorrowingAvailabilityReport ReportType = "Borrowing Availability Report"
)

type ReportFrequencyLabel string

const (
	ReportFrequencyOnce     ReportFrequencyLabel = "Once"
	ReportFrequencyDaily    ReportFrequencyLabel = "Daily"
	ReportFrequencyWeekly   ReportFrequencyLabel = "Weekly"
	ReportFrequencyMonthly  ReportFrequencyLabel = "Monthly"
	ReportFrequencyAnnually ReportFrequencyLabel = "Annually"
	ReportFrequencyCustom   ReportFrequencyLabel = "Custom"
)

type ReportFrequencyUnit string

const (
	ReportFrequencyUnitDays   ReportFrequencyUnit = "Days"
	ReportFrequencyUnitWeeks  ReportFrequencyUnit = "Weeks"
	ReportFrequencyUnitMonths ReportFrequencyUnit = "Months"
	ReportFrequencyUnitYears  ReportFrequencyUnit = "Years"
)

type ReportAssetInformationType string

const (
	ReportAssetInfoBasicInformation   ReportAssetInformationType = "Basic Information"
	ReportAssetInfoBorrow             ReportAssetInformationType = "Borrow"
	ReportAssetInfoTransfer           ReportAssetInformationType = "Transfer"
	ReportAssetInfoAssignment         ReportAssetInformationType = "Assignment"
	ReportAssetInfoTracking           ReportAssetInformationType = "Tracking"
	ReportAssetInfoAudit              ReportAssetInformationType = "Audit"
	ReportAssetInfoDisposal           ReportAssetInformationType = "Disposal"
	ReportAssetInfoMaintenanceRoutine ReportAssetInformationType = "Maintenance Routine"
	ReportAssetInfoRepairTicketing    ReportAssetInformationType = "Repair Ticketing"
	ReportAssetInfoTagging            ReportAssetInformationType = "Tagging"
	ReportAssetInfoLicense            ReportAssetInformationType = "License"
	ReportAssetInfoAssetPolicy        ReportAssetInformationType = "Asset Policy"
	ReportAssetInfoDamage             ReportAssetInformationType = "Damage"
	ReportAssetInfoMissing            ReportAssetInformationType = "Missing"
)

type ListedBy string

const (
	ListedByGroup    ListedBy = "Group"
	ListedByCategory ListedBy = "Category"
)

type ReportInterval struct {
	Label         *string    `json:"label"         bson:"label"`
	CustomDay     []int      `json:"customDay"     bson:"customDay"`
	CustomDate    []int      `json:"customDate"    bson:"customDate"`
	StartInterval *time.Time `json:"startInterval" bson:"startInterval"`
	EndInterval   *time.Time `json:"endInterval"   bson:"endInterval"`
	Month         *time.Time `json:"month"         bson:"month"`
	Year          *time.Time `json:"year"          bson:"year"`
}

type ReportEmail struct {
	Recipient []string `json:"recipient" bson:"recipient"`
	CC        []string `json:"cc"        bson:"cc"`
	BCC       []string `json:"bcc"       bson:"bcc"`
	Subject   string   `json:"subject"   bson:"subject"`
}

type Frequency struct {
	Label       ReportFrequencyLabel `json:"label"       bson:"label"`
	CustomValue *int                 `json:"customValue" bson:"customValue"`
	CustomUnit  *ReportFrequencyUnit `json:"customUnit"  bson:"customUnit"`
}

type Filters struct {
	Group     []int `json:"group"     bson:"group"`
	Category  []int `json:"category"  bson:"category"`
	AssetName []int `json:"assetName" bson:"assetName"`
	Brand     []int `json:"brand"     bson:"brand"`
}

type ModifiedBy struct {
	ID       primitive.ObjectID `json:"_id"      bson:"_id"`
	FullName string             `json:"fullName" bson:"fullName"`
	Key      int                `json:"key"      bson:"key"`
}

type Report struct {
	ID               primitive.ObjectID         `json:"_id"              bson:"_id"`
	IsActive         bool                       `json:"isActive"         bson:"isActive"`
	IsDeleted        bool                       `json:"isDeleted"        bson:"isDeleted"`
	Type             ReportType                 `json:"type"             bson:"type"`
	Name             string                     `json:"name"             bson:"name"`
	Interval         ReportInterval             `json:"interval"         bson:"interval"`
	Email            ReportEmail                `json:"email"            bson:"email"`
	Frequency        Frequency                  `json:"frequency"        bson:"frequency"`
	StartFrom        *time.Time                 `json:"startFrom"        bson:"startFrom"`
	StartDate        *time.Time                 `json:"startDate"        bson:"startDate"` // TODO(refan): current design is not needed, need confirmation
	EndDate          *time.Time                 `json:"endDate"          bson:"endDate"`   // TODO(refan): current design is not needed, need confirmation
	ListedBy         *ListedBy                  `json:"listedBy"         bson:"listedBy"`  // TODO(refan): current design is not needed, need confirmation
	Filters          Filters                    `json:"filters"          bson:"filters"`
	Content          []string                   `json:"content"          bson:"content"`
	Columns          []string                   `json:"columns"          bson:"columns"`
	AssetInformation ReportAssetInformationType `json:"assetInformation" bson:"assetInformation"`
	NextSchedule     *time.Time                 `json:"nextSchedule"     bson:"nextSchedule"`
	ModifiedBy       ModifiedBy                 `json:"modifiedBy"       bson:"modifiedBy"`
	CreatedAt        time.Time                  `json:"createdAt"        bson:"createdAt"`
	UpdatedAt        time.Time                  `json:"updatedAt"        bson:"updatedAt"`
}

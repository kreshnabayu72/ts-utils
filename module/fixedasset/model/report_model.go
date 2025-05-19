package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportType string

const (
	ReportTypeCustomReport             ReportType = "Custom Report"
	ReportTypeSummaryOverview          ReportType = "Summary Overview"
	ReportTypeDepreciationReport       ReportType = "Depreciation Report"
	ReportTypeMaintenanceRoutineReport ReportType = "Maintenance Routine Report"
	ReportTypeRepairmentReport         ReportType = "Repairment Report"
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
	Label         *string    `json:"label,omitempty" bson:"label,omitempty"`
	CustomDay     []int      `json:"customDay,omitempty" bson:"customDay,omitempty"`
	CustomDate    []int      `json:"customDate,omitempty" bson:"customDate,omitempty"`
	StartInterval *time.Time `json:"startInterval,omitempty" bson:"startInterval,omitempty"`
	EndInterval   *time.Time `json:"endInterval,omitempty" bson:"endInterval,omitempty"`
	Month         *time.Time `json:"month,omitempty" bson:"month,omitempty"`
	Year          *time.Time `json:"year,omitempty" bson:"year,omitempty"`
}

type ReportEmail struct {
	Recipient []string `json:"recipient" bson:"recipient"`
	CC        []string `json:"cc,omitempty" bson:"cc,omitempty"`
	BCC       []string `json:"bcc,omitempty" bson:"bcc,omitempty"`
	Subject   string   `json:"subject" bson:"subject"`
}

type Frequency struct {
	Label       ReportFrequencyLabel `json:"label" bson:"label"`
	CustomValue *int                 `json:"customValue,omitempty" bson:"customValue,omitempty"`
	CustomUnit  *ReportFrequencyUnit `json:"customUnit,omitempty" bson:"customUnit,omitempty"`
}

type Filters struct {
	Group     []int `json:"group,omitempty" bson:"group,omitempty"`
	Category  []int `json:"category,omitempty" bson:"category,omitempty"`
	AssetName []int `json:"assetName,omitempty" bson:"assetName,omitempty"`
	Brand     []int `json:"brand,omitempty" bson:"brand,omitempty"`
}

type ModifiedBy struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	FullName string             `json:"fullName" bson:"fullName"`
	Key      int                `json:"key" bson:"key"`
}

type Report struct {
	ID               primitive.ObjectID         `json:"id,omitempty" bson:"_id,omitempty"`
	IsActive         bool                       `json:"isActive" bson:"isActive"`
	IsDeleted        bool                       `json:"isDeleted" bson:"isDeleted"`
	Type             ReportType                 `json:"type" bson:"type"`
	Name             string                     `json:"name" bson:"name"`
	Interval         ReportInterval             `json:"interval" bson:"interval"`
	Email            ReportEmail                `json:"email" bson:"email"`
	Frequency        Frequency                  `json:"frequency" bson:"frequency"`
	StartFrom        *time.Time                 `json:"startFrom,omitempty" bson:"startFrom,omitempty"`
	StartDate        *time.Time                 `json:"startDate,omitempty" bson:"startDate,omitempty"`
	EndDate          *time.Time                 `json:"endDate,omitempty" bson:"endDate,omitempty"`
	ListedBy         *ListedBy                  `json:"listedBy,omitempty" bson:"listedBy,omitempty"`
	Filters          Filters                    `json:"filters" bson:"filters"`
	Content          []string                   `json:"content,omitempty" bson:"content,omitempty"`
	Columns          []string                   `json:"columns,omitempty" bson:"columns,omitempty"`
	AssetInformation ReportAssetInformationType `json:"assetInformation" bson:"assetInformation"`
	NextSchedule     *time.Time                 `json:"nextSchedule,omitempty" bson:"nextSchedule,omitempty"`
	ModifiedBy       ModifiedBy                 `json:"modifiedBy" bson:"modifiedBy"`
	CreatedAt        time.Time                  `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time                  `json:"updatedAt" bson:"updatedAt"`
}

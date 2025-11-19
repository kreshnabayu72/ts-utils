package model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/v2/module/common/model"
)

type Report struct {
	ID        int             `json:"_id"                  column:"reports.id"`
	Name      string          `json:"name"                 column:"reports.name"`
	Type      ReportType      `json:"type"                 column:"reports.type"`
	Frequency ReportFrequency `json:"frequency"            column:"reports.frequency"` //A.K.A Schedule
	StartFrom *time.Time      `json:"startFrom"            column:"reports.start_from"`

	//Not in depreciation report
	TimeInterval           *ReportTimeInterval `json:"timeInterval"           column:"reports.time_interval"`
	TimeIntervalDayOfWeek  *string             `json:"timeIntervalDayOfWeek"  column:"reports.time_interval_day_of_week"`
	TimeIntervalDayOfMonth *string             `json:"timeIntervalDayOfMonth" column:"reports.time_interval_day_of_week"`
	TimeIntervalDayOfYear  *string             `json:"timeIntervalDayOfYear"  column:"reports.time_interval_day_of_week"`

	//Summary Overview
	Content []ReportSummaryOverviewContent `json:"content" column:"reports.content"`

	// Custom Report
	ListedBy             *ReportListedBy             `json:"listedBy"             column:"reports.listed_by"`
	AssetInformationType *ReportAssetInformationType `json:"assetInformationType" column:"reports.asset_information_type"`
	Columns              model.JSONBArrayString      `json:"columns"              column:"reports.columns"`

	//Email
	Subject    string                 `json:"subject"    column:"reports.subject"`
	Recipients model.JSONBArrayString `json:"recipients" column:"reports.recipients"`
	Cc         model.JSONBArrayString `json:"cc"         column:"reports.cc"`
	Bcc        model.JSONBArrayString `json:"bcc"        column:"reports.bcc"`

	//Misc
	NextSchedule *time.Time `json:"nextSchedule" column:"reports.next_schedule"`
	IsActive     bool       `json:"isActive"     column:"reports.is_active"`
	IsDeleted    bool       `json:"isDeleted"    column:"reports.is_deleted"`
	ModifiedById int        `json:"modifiedById" column:"reports.modified_by_id"`
	CreatedAt    time.Time  `json:"createdAt"    column:"reports.created_at"`
	UpdatedAt    time.Time  `json:"updatedAt"    column:"reports.updated_at"`
}

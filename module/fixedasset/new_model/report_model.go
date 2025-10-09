package new_model

import "time"

// TODO(refan): Report - finish commented fields.

type Report struct {
	//Interval
	//Email
	//Frequency
	//Filters
	//Content
	//Columns
	ID                   int                        `json:"_id"                  column:"reports.id"`
	Name                 string                     `json:"name"                 column:"reports.name"`
	IsActive             bool                       `json:"isActive"             column:"reports.is_active"`
	IsDeleted            bool                       `json:"isDeleted"            column:"reports.is_deleted"`
	StartFrom            *time.Time                 `json:"startFrom"            column:"reports.start_from"`
	StartDate            *time.Time                 `json:"startDate"            column:"reports.start_date"` // TODO(refan): current design is not needed, need confirmation
	EndDate              *time.Time                 `json:"endDate"              column:"reports.end_date"`   // TODO(refan): current design is not needed, need confirmation
	ListedBy             *ReportListedBy            `json:"listedBy"             column:"reports.listed_by"`  // TODO(refan): current design is not needed, need confirmation
	AssetInformationType ReportAssetInformationType `json:"assetInformationType" column:"reports.asset_information_type"`
	NextSchedule         *time.Time                 `json:"nextSchedule"         column:"reports.next_schedule"`
	ModifiedById         int                        `json:"modifiedById"         column:"reports.modified_by_id"`
	CreatedAt            time.Time                  `json:"createdAt"            column:"reports.created_at"`
	UpdatedAt            time.Time                  `json:"updatedAt"            column:"reports.updated_at"`
}

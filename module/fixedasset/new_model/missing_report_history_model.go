package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type MissingReportHistory struct {
	ID               int                 `json:"_id"              column:"missing_report_histories.id"                 sql:"BIGINT PRIMARY KEY"`
	Asset            int                 `json:"asset"            column:"missing_report_histories.asset"              sql:"BIGINT"`
	AssetID          string              `json:"assetId"          column:"missing_report_histories.asset_id"           sql:"VARCHAR"`
	TagType          *TagType            `json:"tagType"          column:"missing_report_histories.tag_type"           sql:"tag_type_enum"`
	Name             string              `json:"name"             column:"missing_report_histories.name"               sql:"VARCHAR"`
	Brand            string              `json:"brand"            column:"missing_report_histories.brand"              sql:"VARCHAR"`
	Model            string              `json:"model"            column:"missing_report_histories.model"              sql:"VARCHAR"`
	NameWithSequence string              `json:"nameWithSequence" column:"missing_report_histories.name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                 `json:"sequenceNumber"   column:"missing_report_histories.sequence_number"    sql:"BIGINT"`
	Image            *string             `json:"image"            column:"missing_report_histories.image"              sql:"VARCHAR"`
	Group            string              `json:"group"            column:"missing_report_histories.group"              sql:"VARCHAR"`
	GroupID          int                 `json:"groupId"          column:"missing_report_histories.group_id"           sql:"BIGINT"`
	ReportedBy       int                 `json:"reportedBy"       column:"missing_report_histories.reported_by"        sql:"VARCHAR"`
	Status           MissingReportStatus `json:"status"           column:"missing_report_histories.status"             sql:"missing_report_status_enum NOT NULL"`
	Notes            string              `json:"notes"            column:"missing_report_histories.notes"              sql:"VARCHAR"`
	FoundAt          time.Time           `json:"foundAt"          column:"missing_report_histories.created_at"         sql:"TIMESTAMPTZ"`
	CreatedAt        time.Time           `json:"createdAt"        column:"missing_report_histories.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time           `json:"updatedAt"        column:"missing_report_histories.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *MissingReportHistory) TableName() string {
	return "missing_report_histories"
}

func (m *MissingReportHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}

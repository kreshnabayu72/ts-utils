package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type MissingReportHistory struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	// Foreign key to assets
	Asset int `json:"asset"            column:"asset"           sql:"BIGINT"`
	// Asset ID based on alias code
	AssetID          string   `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	TagType          *TagType `json:"tagType" column:"tag_type" sql:"tag_type_enum"`
	Name             string   `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string   `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string   `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string   `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int      `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Image            *string  `json:"image"         column:"image"        sql:"VARCHAR"`
	Group            string   `json:"group"   column:"group" sql:"VARCHAR"`
	GroupID          int      `json:"groupId"   column:"group_id" sql:"BIGINT"`

	ReportedBy int                 `json:"reportedBy" column:"reported_by" sql:"VARCHAR"`
	Status     MissingReportStatus `json:"status" column:"status" sql:"missing_report_status_enum NOT NULL"`
	Notes      string              `json:"notes" column:"notes" sql:"VARCHAR"`

	FoundAt   time.Time `json:"foundAt" column:"created_at" sql:"TIMESTAMPTZ"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *MissingReportHistory) TableName() string {
	return "missing_report_histories"
}

func (m *MissingReportHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)
}

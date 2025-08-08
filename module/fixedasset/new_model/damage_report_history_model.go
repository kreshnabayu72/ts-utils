package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type DamageReportHistory struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Asset            string  `json:"asset"            column:"asset"           sql:"VARCHAR"`
	AssetID          string  `json:"assetId"            column:"asset_id"           sql:"VARCHAR"`
	TagType          TagType `json:"tagType" column:"tag_type" sql:"tag_type_enum"`
	Name             string  `json:"name"             column:"name"            sql:"VARCHAR"`
	Brand            string  `json:"brand"             column:"brand"            sql:"VARCHAR"`
	Model            string  `json:"model"             column:"model"            sql:"VARCHAR"`
	NameWithSequence string  `json:"nameWithSequence"   column:"name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int     `json:"sequenceNumber"     column:"sequence_number"    sql:"BIGINT"`
	Group            string  `json:"group"   column:"group" sql:"VARCHAR"`
	GroupID          int     `json:"groupId"   column:"group_id" sql:"BIGINT"`

	ReportedBy  string             `json:"reportedBy" column:"reported_by" sql:"VARCHAR"`
	DamageImage string             `json:"damageImage" column:"damage_image" sql:"VARCHAR NOT NULL CHECK (char_length(damage_image) > 0)"`
	Status      DamageReportStatus `json:"status" column:"status" sql:"damage_report_status_enum NOT NULL"`
	Notes       string             `json:"notes" column:"notes" sql:"VARCHAR"`

	RepairedAt time.Time `json:"repairedAt" column:"repaired_at" sql:"TIMESTAMPTZ"`
	CreatedAt  time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt  time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (d *DamageReportHistory) TableName() string {
	return "damage_report_histories"
}

func (d *DamageReportHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(d)
}

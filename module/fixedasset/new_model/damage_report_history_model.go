package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type DamageReportHistory struct {
	ID               int                `json:"_id"                column:"damage_report_histories.id"                 sql:"BIGINT PRIMARY KEY"`
	Asset            int                `json:"asset"              column:"damage_report_histories.asset"              sql:"BIGINT"`
	AssetID          string             `json:"assetId"            column:"damage_report_histories.asset_id"           sql:"VARCHAR"`
	TagType          *TagType           `json:"tagType"            column:"damage_report_histories.tag_type"           sql:"tag_type_enum"`
	Name             string             `json:"name"               column:"damage_report_histories.name"               sql:"VARCHAR"`
	Brand            string             `json:"brand"              column:"damage_report_histories.brand"              sql:"VARCHAR"`
	Model            string             `json:"model"              column:"damage_report_histories.model"              sql:"VARCHAR"`
	NameWithSequence string             `json:"nameWithSequence"   column:"damage_report_histories.name_with_sequence" sql:"VARCHAR"`
	SequenceNumber   int                `json:"sequenceNumber"     column:"damage_report_histories.sequence_number"    sql:"BIGINT"`
	Group            string             `json:"group"              column:"damage_report_histories.group"              sql:"VARCHAR"`
	GroupID          int                `json:"groupId"            column:"damage_report_histories.group_id"           sql:"BIGINT"`
	ReportedBy       string             `json:"reportedBy"         column:"damage_report_histories.reported_by"        sql:"VARCHAR"`
	DamageImage      string             `json:"damageImage"        column:"damage_report_histories.damage_image"       sql:"VARCHAR NOT NULL CHECK (char_length(damage_image) > 0)"`
	Status           DamageReportStatus `json:"status"             column:"damage_report_histories.status"             sql:"damage_report_status_enum NOT NULL"`
	Notes            string             `json:"notes"              column:"damage_report_histories.notes"              sql:"VARCHAR"`
	RepairedAt       time.Time          `json:"repairedAt"         column:"damage_report_histories.repaired_at"        sql:"TIMESTAMPTZ"`
	CreatedAt        time.Time          `json:"createdAt"          column:"damage_report_histories.created_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt        time.Time          `json:"updatedAt"          column:"damage_report_histories.updated_at"         sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (d *DamageReportHistory) TableName() string {
	return "damage_report_histories"
}

func (d *DamageReportHistory) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(d)

	// Add foreign key constraints as defined in the SQL migration file
	builder.AddForeignKey(
		"fk_damage_history_asset",
		"asset",
		"assets",
		"id",
		"",
	)
}

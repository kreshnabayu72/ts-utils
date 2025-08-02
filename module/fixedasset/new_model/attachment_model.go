package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Attachment struct {
	ID        int            `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	AssetID   int            `json:"assetId" column:"asset_id" sql:"BIGINT"`
	Type      AttachmentType `json:"type" column:"type" sql:"attachment_type_enum NOT NULL"`
	Name      string         `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	File      string         `json:"file" column:"file" sql:"VARCHAR NOT NULL"`
	CreatedBy int            `json:"createdBy" column:"created_by" sql:"BIGINT"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (a *Attachment) TableName() string {
	return "attachments"
}

func (a *Attachment) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddForeignKey(
		"fk_attachments_asset",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)

	builder.AddForeignKey(
		"fk_attachments_created_by",
		"created_by",
		"users",
		"id",
		"SET NULL",
	)
}

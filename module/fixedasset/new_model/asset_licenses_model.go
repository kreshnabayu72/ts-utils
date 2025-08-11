package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AssetLicense struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	AssetId   int                `json:"assetId" column:"asset_id" sql:"BIGINT"`
	Type      AssetLicenseType   `json:"type" column:"type" sql:"asset_license_type NOT NULL"`
	IsExpired any                `json:"isExpired" column:"is_expired" sql:"BOOLEAN DEFAULT FALSE"`
	Status    AssetLicenseStatus `json:"status" column:"status" sql:"asset_license_status NOT NULL"`
	AddedAt   time.Time          `json:"addedAt" column:"added_at" sql:"TIMESTAMPTZ NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (a *AssetLicense) TableName() string {
	return "asset_licenses"
}

func (a *AssetLicense) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(a)

	builder.AddForeignKey(
		"fk_asset_licenses_asset_id",
		"asset_id",
		"assets",
		"id",
		"CASCADE",
	)
}

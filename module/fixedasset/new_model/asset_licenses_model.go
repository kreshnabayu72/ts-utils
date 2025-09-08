package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AssetLicense struct {
	ID        int                `json:"_id"       column:"asset_licenses.id"         sql:"BIGINT PRIMARY KEY"`
	AssetId   int                `json:"assetId"   column:"asset_licenses.asset_id"   sql:"BIGINT"`
	Type      AssetLicenseType   `json:"type"      column:"asset_licenses.type"       sql:"asset_license_type NOT NULL"`
	IsExpired any                `json:"isExpired" column:"asset_licenses.is_expired" sql:"BOOLEAN DEFAULT FALSE"`
	Status    AssetLicenseStatus `json:"status"    column:"asset_licenses.status"     sql:"asset_license_status NOT NULL"`
	AddedAt   time.Time          `json:"addedAt"   column:"asset_licenses.added_at"   sql:"TIMESTAMPTZ NOT NULL"`
	CreatedAt time.Time          `json:"createdAt" column:"asset_licenses.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time          `json:"updatedAt" column:"asset_licenses.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

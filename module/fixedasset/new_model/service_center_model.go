package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type ServiceCenter struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	IsActive       bool       `json:"isActive" column:"is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsDefault      bool       `json:"isDefault" column:"is_default" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Name           string     `json:"name" column:"name" sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	City           string     `json:"city" column:"city" sql:"VARCHAR"`
	Country        string     `json:"country" column:"country" sql:"VARCHAR"`
	State          string     `json:"state" column:"state" sql:"VARCHAR"`
	ZipCode        string     `json:"zipCode" column:"zip_code" sql:"VARCHAR"`
	Address        string     `json:"address" column:"address" sql:"VARCHAR"`
	PhoneNumber    string     `json:"phoneNumber" column:"phone_number" sql:"VARCHAR"`
	PicPhoneNumber string     `json:"picPhoneNumber" column:"pic_phone_number" sql:"VARCHAR"`
	PicName        string     `json:"picName" column:"pic_name" sql:"VARCHAR"`
	ModifiedById   *int       `json:"modifiedById" column:"modified_by_id" sql:"BIGINT"`
	IsDeleted      bool       `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt      *time.Time `json:"deletedAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt      time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *ServiceCenter) TableName() string {
	return "service_centers"
}

func (m *ServiceCenter) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("idx_service_centers_name", "name", false, "")
	builder.AddForeignKey(
		"fk_service_centers_modified_by",
		"modified_by_id",
		"users",
		"id",
		"SET NULL",
	)

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.ServiceCenterPostQuery)
}

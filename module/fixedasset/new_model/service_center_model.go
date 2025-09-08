package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type ServiceCenter struct {
	ID             int        `json:"_id"            column:"service_centers.id"               sql:"BIGINT PRIMARY KEY"`
	IsActive       bool       `json:"isActive"       column:"service_centers.is_active"        sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsDefault      bool       `json:"isDefault"      column:"service_centers.is_default"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Name           string     `json:"name"           column:"service_centers.name"             sql:"VARCHAR NOT NULL CHECK (CHAR_LENGTH(name) > 0)"`
	City           string     `json:"city"           column:"service_centers.city"             sql:"VARCHAR"`
	Country        string     `json:"country"        column:"service_centers.country"          sql:"VARCHAR"`
	State          string     `json:"state"          column:"service_centers.state"            sql:"VARCHAR"`
	ZipCode        string     `json:"zipCode"        column:"service_centers.zip_code"         sql:"VARCHAR"`
	Address        string     `json:"address"        column:"service_centers.address"          sql:"VARCHAR"`
	PhoneNumber    string     `json:"phoneNumber"    column:"service_centers.phone_number"     sql:"VARCHAR"`
	PicPhoneNumber string     `json:"picPhoneNumber" column:"service_centers.pic_phone_number" sql:"VARCHAR"`
	PicName        string     `json:"picName"        column:"service_centers.pic_name"         sql:"VARCHAR"`
	ModifiedById   *int       `json:"modifiedById"   column:"service_centers.modified_by_id"   sql:"BIGINT"`
	IsDeleted      bool       `json:"isDeleted"      column:"service_centers.is_deleted"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeletedAt      *time.Time `json:"deletedAt"      column:"service_centers.deleted_at"       sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt      time.Time  `json:"createdAt"      column:"service_centers.created_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time  `json:"updatedAt"      column:"service_centers.updated_at"       sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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

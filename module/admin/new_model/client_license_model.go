package new_model

import (
	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	admin_migration "gitea.qwertysystem.net/BETS/ts-utils/module/admin/migration"
)

type ClientLicense struct {
	ID           int                 `json:"_id"          column:"client_licenses.id"            sql:"BIGINT PRIMARY KEY"`
	ClientID     int                 `json:"clientId"     column:"client_licenses.client_id"     sql:"BIGINT"`
	Type         LicenseType         `json:"type"         column:"client_licenses.type"` // TODO: sql
	Managed      int                 `json:"managed"      column:"client_licenses.managed"       sql:"INTEGER"`
	Unmanaged    int                 `json:"unmanaged"    column:"client_licenses.unmanaged"     sql:"INTEGER"`
	Available    int                 `json:"available"    column:"client_licenses.available"     sql:"INTEGER"`
	Total        int                 `json:"total"        column:"client_licenses.total"         sql:"INTEGER"`
	PurchaseType LicensePurchaseType `json:"purchaseType" column:"client_licenses.purchase_type"` // TODO: sql
}

func (m *ClientLicense) TableName() string {
	return "client_licenses"
}

func (m *ClientLicense) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey(
		"client_licenses_client_id_foreign_key",
		"client_id",
		"clients",
		"id",
		"CASCADE",
	)

	// ─────── Additional Queries ───────
	builder.SetPostQuery(admin_migration.ClientLicensePostQuery)
}

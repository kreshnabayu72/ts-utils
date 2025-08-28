package new_model

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ClientLicense struct {
	ID           int                 `json:"_id"          column:"client_licenses.id"            sql:"BIGINT PRIMARY KEY"`
	ClientID     int                 `json:"clientId"     column:"client_licenses.client_id"     sql:"BIGINT"`
	Type         LicenseType         `json:"type"         column:"client_licenses.type" sql:"license_type_enum NOT NULL"`
	Managed      int                 `json:"managed"      column:"client_licenses.managed"       sql:"INTEGER"`
	Unmanaged    int                 `json:"unmanaged"    column:"client_licenses.unmanaged"     sql:"INTEGER"`
	Available    int                 `json:"available"    column:"client_licenses.available"     sql:"INTEGER"`
	Total        int                 `json:"total"        column:"client_licenses.total"         sql:"INTEGER"`
	PurchaseType LicensePurchaseType `json:"purchaseType" column:"client_licenses.purchase_type" sql:"license_purchase_type_enum NOT NULL"`
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
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	path := filepath.Join(baseDir, "migration/client_license_post_query.sql")

	postQuery, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error path: " + path)
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}

package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type LicenseTransaction struct {
	ID            int                 `json:"_id"           column:"license_transactions.id"             sql:"BIGINT PRIMARY KEY"`
	Type          LicenseModuleType   `json:"type"          column:"license_transactions.type"           sql:"license_module_enum"`
	LicenseType   LicensePurchaseType `json:"licenseType"   column:"license_transactions.license_type"   sql:"license_transaction_type_enum"`
	Amount        int                 `json:"amount"        column:"license_transactions.amount"         sql:"INTEGER"`
	DurationMonth int                 `json:"durationMonth" column:"license_transactions.duration_month" sql:"INTEGER"`
	DurationText  int                 `json:"durationText"  column:"license_transactions.duration_text"  sql:"VARCHAR"`
	PurchaseType  LicensePurchaseType `json:"purchaseType"  column:"license_transactions.purchase_type"  sql:"license_transaction_purchase_type_enum"`
	ActivatedAt   time.Time           `json:"activatedAt"   column:"license_transactions.activated_at"   sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	ExpiredAt     time.Time           `json:"expiredAt"     column:"license_transactions.expired_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	CreatedAt     time.Time           `json:"createdAt"     column:"license_transactions.created_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt     time.Time           `json:"updatedAt"     column:"license_transactions.updated_at"     sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (l *LicenseTransaction) TableName() string {
	return "license_transactions"
}

func (l *LicenseTransaction) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(l)
}

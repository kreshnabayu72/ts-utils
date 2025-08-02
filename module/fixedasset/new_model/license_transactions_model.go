package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type LicenseTransaction struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Type          LicenseModuleType   `json:"type" column:"type" sql:"license_module_enum"`
	LicenseType   LicensePurchaseType `json:"licenseType" column:"license_type" sql:"license_transaction_type_enum"`
	Amount        int                 `json:"amount" column:"amount" sql:"INTEGER"`
	DurationMonth int                 `json:"durationMonth" column:"duration_month" sql:"INTEGER"`
	DurationText  int                 `json:"durationText" column:"duration_text" sql:"VARCHAR"`
	PurchaseType  LicensePurchaseType `json:"purchaseType" column:"purchase_type" sql:"license_transaction_purchase_type_enum"`

	ActivatedAt time.Time `json:"activatedAt" column:"activated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	ExpiredAt   time.Time `json:"expiredAt" column:"expired_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	CreatedAt   time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (l *LicenseTransaction) TableName() string {
	return "license_transactions"
}

func (l *LicenseTransaction) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(l)
}

package new_model

import (
	"time"
)

type LicenseTransaction struct {
	ID            int                 `json:"_id"           column:"license_transactions.id"`
	Type          LicenseModuleType   `json:"type"          column:"license_transactions.type"`
	LicenseType   LicensePurchaseType `json:"licenseType"   column:"license_transactions.license_type"`
	Amount        int                 `json:"amount"        column:"license_transactions.amount"`
	DurationMonth int                 `json:"durationMonth" column:"license_transactions.duration_month"`
	DurationText  int                 `json:"durationText"  column:"license_transactions.duration_text"`
	PurchaseType  LicensePurchaseType `json:"purchaseType"  column:"license_transactions.purchase_type"`
	ActivatedAt   time.Time           `json:"activatedAt"   column:"license_transactions.activated_at"`
	ExpiredAt     time.Time           `json:"expiredAt"     column:"license_transactions.expired_at"`
	CreatedAt     time.Time           `json:"createdAt"     column:"license_transactions.created_at"`
	UpdatedAt     time.Time           `json:"updatedAt"     column:"license_transactions.updated_at"`
}

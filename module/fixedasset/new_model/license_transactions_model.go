package new_model

import "time"

type LicenseTransaction struct {
	ID int `json:"_id"                  column:"id"`

	Type          LicenseModuleType   `json:"type"                 column:"type"`
	LicenseType   LicensePurchaseType `json:"licenseType"                 column:"license_type"`
	Amount        int                 `json:"amount"                 column:"amount"`
	DurationMonth int                 `json:"durationMonth"                 column:"durationMonth"`
	DurationText  int                 `json:"durationText"                 column:"durationText"`
	PurchaseType  LicensePurchaseType `json:"purchaseType"                 column:"purchase_type"`

	ActivatedAt time.Time `json:"activatedAt"            column:"activated_at"`
	ExpiredAt   time.Time `json:"expiredAt"            column:"expired_at"`
	CreatedAt   time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt"            column:"updated_at"`
}

package new_model

import "time"

type Company struct {
	ID int `json:"_id" column:"id"`

	// ─────── Company Information ───────
	Name               string        `json:"name" column:"name"`
	ClientID           string        `json:"clientId"  column:"client_id"`
	Status             CompanyStatus `json:"status" column:"status"`
	Email              string        `json:"email" column:"email"`
	IsEmailConfirmed   bool          `json:"isEmailConfirmed" column:"is_email_confirmed"`
	Country            string        `json:"country" column:"country"`
	State              string        `json:"state" column:"state"`
	City               string        `json:"city" column:"city"`
	ZIP                string        `json:"zip" column:"zip"`
	PhoneNumber        string        `json:"phoneNumber" column:"phone_number"`
	Fax                string        `json:"fax" column:"fax"`
	Address            string        `json:"address" column:"address"`
	ContactPersonName  string        `json:"contactPersonName" column:"contact_person_name"`
	ContactPersonEmail string        `json:"contactPersonEmail" column:"contact_person_email"`
	PendingEmailChange *string       `json:"pendingEmailChange" column:"pending_email_change"`
	Image              *string       `json:"image" column:"image"`
	GroupID            int           `json:"groupId" column:"group_id"`
	ModifiedByID       int           `json:"modifiedById" column:"modified_by_id"`
	PartnerID          int           `json:"partnerId" column:"partner_id"`
	ExpiredAt          time.Time     `json:"expiredAt" column:"expired_at"`
	CreatedAt          time.Time     `json:"createdAt" column:"created_at"`
	UpdatedAt          time.Time     `json:"updatedAt" column:"updated_at"`

	// ─────── Open API Settings ───────
	OpenAPIStatus      *CompanyOpenAPIStatus `json:"openApiStatus" column:"open_api_status"`
	OpenAPIRequestedAt *time.Time            `json:"openApiRequestedAt" column:"open_api_requested_at"`
	OpenAPIApprovedAt  *time.Time            `json:"openApiApprovedAt" column:"open_api_approved_at"`

	// ─────── License Plan ───────
	PlanFixedAsset  PlanTypes `json:"planFixedAsset" column:"plan_fixed_asset"`
	PlanSupplyAsset PlanTypes `json:"planSupplyAsset" column:"plan_supply_asset"`
	PlanInventory   PlanTypes `json:"planInventory" column:"plan_inventory"`

	// ─────── Deprec Method ───────
	DepreciationMethod        CompanyDepreciationMethod  `json:"depreciationMethod" column:"depreciation_method"`
	PendingDepreciationMethod *CompanyDepreciationMethod `json:"pendingDepreciationMethod" column:"pending_depreciation_method"`

	// ─────── Add On Status ───────
	AddOnRepairTicketing bool `json:"addOnRepairTicketing" column:"add_on_repair_ticketing"`
	AddOnMaintenance     bool `json:"addOnMaintenance" column:"add_on_maintenance"`
	AddOnTracking        bool `json:"addOnTracking" column:"add_on_tracking"`
	AddOnAssetControl    bool `json:"addOnAssetControl" column:"add_on_asset_control"`
	AddOnAudit           bool `json:"addOnAudit" column:"add_on_audit"`

	// ─────── Transaction Settings ───────
	AuditAttrsMaxGracePeriod           int `json:"auditAttrsMaxGracePeriod" column:"audit_attrs_max_grace_period"`
	TransferAttrsTransactionTimeout    int `json:"transferAttrsTransactionTimeout" column:"transfer_attrs_transaction_timeout"`
	TransferAttrsAssetControlToMissing int `json:"transferAttrsAssetControlToMissing" column:"transfer_attrs_asset_control_to_missing"`
	BorrowingAttrsMaxBorrowingTime     int `json:"borrowingAttrsMaxBorrowingTime" column:"borrowing_attrs_max_borrowing_time"`
	BorrowingAttrsDefaultBorrowingTime int `json:"borrowingAttrsDefaultBorrowingTime" column:"borrowing_attrs_default_borrowing_time"`
	BorrowingAttrsConfirmationTimeout  int `json:"borrowingAttrsConfirmationTimeout" column:"borrowing_attrs_confirmation_timeout"`
	BorrowingAttrsTransactionTimeout   int `json:"borrowingAttrsTransactionTimeout" column:"borrowing_attrs_transaction_timeout"`
	DisposalAttrsTransactionTimeout    int `json:"disposalAttrsTransactionTimeout" column:"disposal_attrs_transaction_timeout"`

	// ─────── General Settings ───────
	LoginSessionMinutes int    `json:"loginSessionMinutes" column:"login_session_minutes"`
	DateFormat          string `json:"dateFormat" column:"date_format"`
	TimeFormat          string `json:"timeFormat" column:"time_format"`
	Geolocation         bool   `json:"geolocation" column:"geolocation"`
	TimeZone            string `json:"timeZone" column:"time_zone"`
	PIN                 string `json:"pin" column:"pin"`
	CurrencySymbol      string `json:"currencySymbol" column:"currency_symbol"`
	CurrencyCurrency    string `json:"currencyCurrency" column:"currency_currency"`
	CurrencyLocale      string `json:"currencyLocale" column:"currency_locale"`
	CurrencyLabel       string `json:"currencyLabel" column:"currency_label"`

	// ─────── File Storage ───────
	FileStorageTotalKB           float64 `json:"fileStorageTotalKB" column:"file_storage_total_kb"`
	FileStorageAvailableKB       float64 `json:"fileStorageAvailableKB" column:"file_storage_available_kb"`
	FileStorageUsedKB            float64 `json:"fileStorageUsedKB" column:"file_storage_used_kb"`
	FileStorageAttachmentKB      float64 `json:"fileStorageAttachmentKB" column:"file_storage_attachment_kb"`
	FileStorageDamagedKB         float64 `json:"fileStorageDamagedKB" column:"file_storage_damaged_kb"`
	FileStorageDisposalKB        float64 `json:"fileStorageDisposalKB" column:"file_storage_disposal_kb"`
	FileStorageRepairTicketingKB float64 `json:"fileStorageRepairTicketingKB" column:"file_storage_repair_ticketing_kb"`
	FileStorageRoutineKB         float64 `json:"fileStorageRoutineKB" column:"file_storage_routine_kb"`
}

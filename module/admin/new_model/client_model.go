package new_model

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
)

type Client struct {
	ID int `json:"id" column:"id" sql:"BIGINT PRIMARY KEY"`
	// ─────── Company Information ───────
	Name               string        `json:"name" column:"name" sql:"BIGINT PRIMARY KEY"`
	ClientID           string        `json:"clientId"  column:"client_id" sql:"VARCHAR NOT NULL"`
	Status             CompanyStatus `json:"status" column:"status" sql:"client_status_enum NOT NULL"`
	Email              *string       `json:"email" column:"email" sql:"VARCHAR NULL"`
	IsEmailConfirmed   bool          `json:"isEmailConfirmed" column:"is_email_confirmed" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Country            string        `json:"country" column:"country" sql:"VARCHAR NOT NULL"`
	State              string        `json:"state" column:"state" sql:"VARCHAR NOT NULL"`
	City               string        `json:"city" column:"city" sql:"VARCHAR NOT NULL"`
	ZIP                string        `json:"zip" column:"zip" sql:"VARCHAR NOT NULL"`
	PhoneNumber        string        `json:"phoneNumber" column:"phone_number" sql:"VARCHAR NOT NULL"`
	Fax                *string       `json:"fax" column:"fax" sql:"VARCHAR NULL"`
	Address            string        `json:"address" column:"address" sql:"VARCHAR NOT NULL"`
	ContactPersonName  string        `json:"contactPersonName" column:"contact_person_name" sql:"VARCHAR NOT NULL"`
	ContactPersonEmail string        `json:"contactPersonEmail" column:"contact_person_email" sql:"VARCHAR NOT NULL"`
	PendingEmailChange *string       `json:"pendingEmailChange" column:"pending_email_change" sql:"VARCHAR NULL"`
	Image              *string       `json:"image" column:"image" sql:"VARCHAR NULL"`
	GroupID            int           `json:"groupId" column:"group_id"` // TODO: sql
	ModifiedByID       int           `json:"modifiedById" column:"modified_by_id" sql:"BIGINT"`
	PartnerID          int           `json:"partnerId" column:"partner_id"` // TODO: sql
	ExpiredAt          time.Time     `json:"expiredAt" column:"expired_at" sql:"TIMESTAMP NULL"`
	CreatedAt          time.Time     `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time     `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`

	// ─────── Open API Settings ───────
	OpenAPIStatus      *CompanyOpenAPIStatus `json:"openApiStatus" column:"open_api_status" sql:"open_api_status_enum NULL"`
	OpenAPIRequestedAt *time.Time            `json:"openApiRequestedAt" column:"open_api_requested_at" sql:"TIMESTAMP NULL"`
	OpenAPIApprovedAt  *time.Time            `json:"openApiApprovedAt" column:"open_api_approved_at" sql:"TIMESTAMP NULL"`

	// ─────── License Plan ───────
	PlanFixedAsset       fixedasset_model.PlanTypes `json:"planFixedAsset" column:"plan_fixed_asset" sql:"license_plan_enum NULL"`
	PlanSupplyAsset      fixedasset_model.PlanTypes `json:"planSupplyAsset" column:"plan_supply_asset" sql:"license_plan_enum NULL"`
	PlanInventory        fixedasset_model.PlanTypes `json:"planInventory" column:"plan_inventory" sql:"license_plan_enum NULL"`
	PlanProcurementAsset fixedasset_model.PlanTypes `json:"planProcurementAsset" column:"plan_procurement_asset" sql:"license_plan_enum NULL"`

	// ─────── Deprec Method ───────
	DepreciationMethod        CompanyDepreciationMethod  `json:"depreciationMethod" column:"depreciation_method" sql:"depreciation_method_enum NOT NULL"`
	PendingDepreciationMethod *CompanyDepreciationMethod `json:"pendingDepreciationMethod" column:"pending_depreciation_method" sql:"depreciation_method_enum NULL"`

	// ─────── Add On Status ───────
	AddOnRepairTicketing bool `json:"addOnRepairTicketing" column:"add_on_repair_ticketing" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnMaintenance     bool `json:"addOnMaintenance" column:"add_on_maintenance" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnTracking        bool `json:"addOnTracking" column:"add_on_tracking" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAssetControl    bool `json:"addOnAssetControl" column:"add_on_asset_control" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAudit           bool `json:"addOnAudit" column:"add_on_audit" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`

	// ─────── Transaction Settings ───────
	AuditAttrsMaxGracePeriod           int `json:"auditAttrsMaxGracePeriod" column:"audit_attrs_max_grace_period" sql:"INTEGER"`
	TransferAttrsTransactionTimeout    int `json:"transferAttrsTransactionTimeout" column:"transfer_attrs_transaction_timeout" sql:"INTEGER"`
	TransferAttrsAssetControlToMissing int `json:"transferAttrsAssetControlToMissing" column:"transfer_attrs_asset_control_to_missing" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	BorrowingAttrsMaxBorrowingTime     int `json:"borrowingAttrsMaxBorrowingTime" column:"borrowing_attrs_max_borrowing_time" sql:"INTEGER"`
	BorrowingAttrsDefaultBorrowingTime int `json:"borrowingAttrsDefaultBorrowingTime" column:"borrowing_attrs_default_borrowing_time" sql:"INTEGER"`
	BorrowingAttrsConfirmationTimeout  int `json:"borrowingAttrsConfirmationTimeout" column:"borrowing_attrs_confirmation_timeout" sql:"INTEGER"`
	BorrowingAttrsTransactionTimeout   int `json:"borrowingAttrsTransactionTimeout" column:"borrowing_attrs_transaction_timeout" sql:"INTEGER"`
	DisposalAttrsTransactionTimeout    int `json:"disposalAttrsTransactionTimeout" column:"disposal_attrs_transaction_timeout" sql:"INTEGER"`

	// ─────── General Settings ───────
	LoginSessionMinutes int    `json:"loginSessionMinutes" column:"login_session_minutes" sql:"INTEGER NULL"`
	DateFormat          string `json:"dateFormat" column:"date_format" sql:"VARCHAR NOT NULL"`
	TimeFormat          string `json:"timeFormat" column:"time_format" sql:"VARCHAR NOT NULL DEFAULT FALSE"`
	Geolocation         bool   `json:"geolocation" column:"geolocation" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TimeZone            string `json:"timeZone" column:"time_zone" sql:"VARCHAR NOT NULL"`
	PIN                 string `json:"pin" column:"pin" sql:"VARCHAR(6) NOT NULL"`
	CurrencySymbol      string `json:"currencySymbol" column:"currency_symbol" sql:"VARCHAR NOT NULL"`
	CurrencyCurrency    string `json:"currencyCurrency" column:"currency_currency" sql:"VARCHAR NOT NULL"`
	CurrencyLocale      string `json:"currencyLocale" column:"currency_locale" sql:"VARCHAR NOT NULL"`
	CurrencyLabel       string `json:"currencyLabel" column:"currency_label" sql:"VARCHAR NOT NULL"`

	// ─────── File Storage ───────
	FileStorageTotalKB           float64 `json:"fileStorageTotalKB" column:"file_storage_total_kb" sql:"FLOAT"`
	FileStorageAvailableKB       float64 `json:"fileStorageAvailableKB" column:"file_storage_available_kb" sql:"FLOAT"`
	FileStorageUsedKB            float64 `json:"fileStorageUsedKB" column:"file_storage_used_kb" sql:"FLOAT"`
	FileStorageAttachmentKB      float64 `json:"fileStorageAttachmentKB" column:"file_storage_attachment_kb" sql:"FLOAT"`
	FileStorageDamagedKB         float64 `json:"fileStorageDamagedKB" column:"file_storage_damaged_kb" sql:"FLOAT"`
	FileStorageDisposalKB        float64 `json:"fileStorageDisposalKB" column:"file_storage_disposal_kb" sql:"FLOAT"`
	FileStorageRepairTicketingKB float64 `json:"fileStorageRepairTicketingKB" column:"file_storage_repair_ticketing_kb" sql:"FLOAT"`
	FileStorageRoutineKB         float64 `json:"fileStorageRoutineKB" column:"file_storage_routine_kb" sql:"FLOAT"`
}

func (m *Client) TableName() string {
	return "clients"
}

func (m *Client) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	// ─────── Enums ───────
	builder.AddEnum(
		"client_status_enum",
		[]string{
			string(CompanyStatusActive),
			string(CompanyStatusExpired),
			string(CompanyStatusArchived),
		},
	)
	builder.AddEnum(
		"open_api_status_enum",
		[]string{
			string(CompanyOpenAPIStatusApproved),
			string(CompanyOpenAPIStatusWaiting),
		},
	)
	builder.AddEnum(
		"depreciation_method_enum",
		[]string{
			string(CompanyDepreciationMethodStraightLine),
			string(CompanyDepreciationMethodDoubleDecliningBalance),
		},
	)
	builder.AddEnum(
		"license_plan_enum",
		[]string{
			string(fixedasset_model.PlanBasic),
			string(fixedasset_model.PlanStandard),
			string(fixedasset_model.PlanEnterprise),
		},
	)
	builder.AddEnum("purchase_type_enum", []string{"Purchase", "Trial"})
	builder.AddEnum(
		"license_type_enum",
		[]string{
			"concurrent_user",
			"fixed_asset",
			"fixed_asset_audit",
			"fixed_asset_asset_control",
			"fixed_asset_maintenance",
			"fixed_asset_repair_ticketing",
			"fixed_asset_tracking",
			"supply_asset",
			"supply_asset_asset_control",
			"supply_asset_stock_opname",
			"supply_asset_tracking",
			"inventory",
		},
	)

	// ─────── Foreign Keys ───────
	builder.AddIndex("clients_name_uindex", "name", true, "")
	builder.AddIndex("clients_client_id_uindex", "client_id", true, "")

	// ─────── Additional Queries ───────
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	postQuery, err := os.ReadFile(
		filepath.Join(baseDir, "..", "migration/client_post_query.sql"),
	)
	if err != nil {
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}

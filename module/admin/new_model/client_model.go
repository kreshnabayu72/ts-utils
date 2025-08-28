package new_model

import (
	"fmt"
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
	Name               string        `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	ClientID           string        `json:"clientId"  column:"client_id" sql:"VARCHAR NOT NULL"`
	Status             CompanyStatus `json:"status" column:"status" sql:"client_status_enum NOT NULL DEFAULT 'Active'"`
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
	PendingEmailChange *string       `json:"pendingEmailChange" column:"pending_email_change" sql:"VARCHAR NULL DEFAULT NULL"`
	Image              *string       `json:"image" column:"image" sql:"VARCHAR NULL DEFAULT NULL"`
	GroupID            int           `json:"groupId" column:"group_id" sql:"BIGINT NULL"`
	ModifiedByID       int           `json:"modifiedById" column:"modified_by_id" sql:"BIGINT NULL"`
	PartnerID          int           `json:"partnerId" column:"partner_id" sql:"BIGINT NULL"`
	ExpiredAt          time.Time     `json:"expiredAt" column:"expired_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt          time.Time     `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt          time.Time     `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`

	// ─────── Open API Settings ───────
	OpenAPIStatus      *CompanyOpenAPIStatus `json:"openApiStatus" column:"open_api_status" sql:"open_api_status_enum NULL DEFAULT NULL"`
	OpenAPIRequestedAt *time.Time            `json:"openApiRequestedAt" column:"open_api_requested_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	OpenAPIApprovedAt  *time.Time            `json:"openApiApprovedAt" column:"open_api_approved_at" sql:"TIMESTAMPTZ NULL DEFAULT NULL"`

	// ─────── License Plan ───────
	PlanFixedAsset       fixedasset_model.PlanTypes `json:"planFixedAsset" column:"plan_fixed_asset" sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanSupplyAsset      fixedasset_model.PlanTypes `json:"planSupplyAsset" column:"plan_supply_asset" sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanInventory        fixedasset_model.PlanTypes `json:"planInventory" column:"plan_inventory_asset" sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanProcurementAsset fixedasset_model.PlanTypes `json:"planProcurementAsset" column:"plan_procurement_asset" sql:"license_plan_enum DEFAULT 'Basic'"`

	// ─────── Deprec Method ───────
	DepreciationMethod        CompanyDepreciationMethod  `json:"depreciationMethod" column:"depreciation_method" sql:"depreciation_method_enum NOT NULL DEFAULT 'Straight-Line'"`
	PendingDepreciationMethod *CompanyDepreciationMethod `json:"pendingDepreciationMethod" column:"pending_depreciation_method" sql:"depreciation_method_enum NULL DEFAULT NULL"`

	// ─────── Add On Status ───────
	AddOnRepairTicketing bool `json:"addOnRepairTicketing" column:"addon_repair_ticketing" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnMaintenance     bool `json:"addOnMaintenance" column:"addon_maintenance" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnTracking        bool `json:"addOnTracking" column:"addon_tracking" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAssetControl    bool `json:"addOnAssetControl" column:"addon_asset_control" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAudit           bool `json:"addOnAudit" column:"addon_audit" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`

	// ─────── Transaction Settings ───────
	AuditAttrsMaxGracePeriod           int `json:"auditAttrsMaxGracePeriod" column:"audit_attrs_max_grace_period" sql:"INTEGER DEFAULT 3"`
	TransferAttrsTransactionTimeout    int `json:"transferAttrsTransactionTimeout" column:"transfer_attrs_transaction_timeout" sql:"INTEGER DEFAULT 5"`
	TransferAttrsAssetControlToMissing int `json:"transferAttrsAssetControlToMissing" column:"transfer_attrs_asset_control_to_missing" sql:"INTEGER DEFAULT 1"`
	BorrowingAttrsMaxBorrowingTime     int `json:"borrowingAttrsMaxBorrowingTime" column:"borrowing_attrs_max_borrowing_time" sql:"INTEGER DEFAULT 15"`
	BorrowingAttrsDefaultBorrowingTime int `json:"borrowingAttrsDefaultBorrowingTime" column:"borrowing_attrs_default_borrowing_time" sql:"INTEGER DEFAULT 7"`
	BorrowingAttrsConfirmationTimeout  int `json:"borrowingAttrsConfirmationTimeout" column:"borrowing_attrs_confirmation_timeout" sql:"INTEGER DEFAULT 24"`
	BorrowingAttrsTransactionTimeout   int `json:"borrowingAttrsTransactionTimeout" column:"borrowing_attrs_transaction_timeout" sql:"INTEGER DEFAULT 5"`
	DisposalAttrsTransactionTimeout    int `json:"disposalAttrsTransactionTimeout" column:"disposal_attrs_transaction_timeout" sql:"INTEGER DEFAULT 5"`

	// ─────── General Settings ───────
	LoginSessionMinutes int    `json:"loginSessionMinutes" column:"login_session_minutes" sql:"INTEGER DEFAULT 30"`
	DateFormat          string `json:"dateFormat" column:"date_format" sql:"VARCHAR DEFAULT '02/01/06'"`
	TimeFormat          string `json:"timeFormat" column:"time_format" sql:"VARCHAR NOT NULL DEFAULT '15:04:05'"`
	Geolocation         bool   `json:"geolocation" column:"geolocation" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TimeZone            string `json:"timeZone" column:"time_zone" sql:"VARCHAR default 'Asia/Jakarta'"`
	PIN                 string `json:"pin" column:"pin" sql:"VARCHAR(6) DEFAULT '123456'"`
	CurrencySymbol      string `json:"currencySymbol" column:"currency_symbol" sql:"VARCHAR DEFAULT 'IDR'"`
	CurrencyCurrency    string `json:"currencyCurrency" column:"currency_currency" sql:"VARCHAR DEFAULT 'Rp'"`
	CurrencyLocale      string `json:"currencyLocale" column:"currency_locale" sql:"VARCHAR DEFAULT 'id-ID'"`
	CurrencyLabel       string `json:"currencyLabel" column:"currency_label" sql:"VARCHAR DEFAULT 'Rp - Indonesian Rupiah'"`

	// ─────── File Storage ───────
	FileStorageTotalKB           float64 `json:"fileStorageTotalKB" column:"file_storage_total_kb" sql:"FLOAT DEFAULT 5242880"`         //5 gb
	FileStorageAvailableKB       float64 `json:"fileStorageAvailableKB" column:"file_storage_available_kb" sql:"FLOAT DEFAULT 5242880"` //5gb
	FileStorageUsedKB            float64 `json:"fileStorageUsedKB" column:"file_storage_used_kb" sql:"FLOAT DEFAULT 0"`
	FileStorageAttachmentKB      float64 `json:"fileStorageAttachmentKB" column:"file_storage_attachment_kb" sql:"FLOAT DEFAULT 0"`
	FileStorageDamagedKB         float64 `json:"fileStorageDamagedKB" column:"file_storage_damage_kb" sql:"FLOAT DEFAULT 0"`
	FileStorageDisposalKB        float64 `json:"fileStorageDisposalKB" column:"file_storage_disposal_kb" sql:"FLOAT DEFAULT 0"`
	FileStorageRepairTicketingKB float64 `json:"fileStorageRepairTicketingKB" column:"file_storage_repair_ticketing_kb" sql:"FLOAT DEFAULT 0"`
	FileStorageRoutineKB         float64 `json:"fileStorageRoutineKB" column:"file_storage_routine_kb" sql:"FLOAT DEFAULT 0"`
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

	// ─────── Index ───────
	builder.AddIndex("clients_name_uindex", "name", true, "")
	builder.AddIndex("clients_client_id_uindex", "client_id", true, "")

	// ─────── Foreign Keys ───────
	// builder.AddForeignKey("fk_client_group", "group_id", "groups", "id", "SET NULL")
	// builder.AddForeignKey("fk_client_modified_by_id", "modified_by_id", "users", "id", "SET NULL")
	// builder.AddForeignKey("fk_client_partner_id", "partner_id", "partners", "id", "SET NULL")

	// ─────── Additional Queries ───────
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	path := filepath.Join(baseDir, "migration/client_post_query.sql")

	postQuery, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error path: " + path)
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}

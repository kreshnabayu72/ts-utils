package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	admin_migration "gitea.qwertysystem.net/BETS/ts-utils/module/admin/new_model/migration"
	fixedasset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
)

type Client struct {
	ID                                 int                        `json:"id"                                 column:"clients.id"                                      sql:"BIGINT PRIMARY KEY"`
	Name                               string                     `json:"name"                               column:"clients.name"                                    sql:"VARCHAR NOT NULL"`
	ClientID                           string                     `json:"clientId"                           column:"clients.client_id"                               sql:"VARCHAR NOT NULL"`
	Status                             CompanyStatus              `json:"status"                             column:"clients.status"                                  sql:"client_status_enum NOT NULL DEFAULT 'Active'"`
	Email                              *string                    `json:"email"                              column:"clients.email"                                   sql:"VARCHAR NULL"`
	IsEmailConfirmed                   bool                       `json:"isEmailConfirmed"                   column:"clients.is_email_confirmed"                      sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Country                            string                     `json:"country"                            column:"clients.country"                                 sql:"VARCHAR NOT NULL"`
	State                              string                     `json:"state"                              column:"clients.state"                                   sql:"VARCHAR NOT NULL"`
	City                               string                     `json:"city"                               column:"clients.city"                                    sql:"VARCHAR NOT NULL"`
	ZIP                                string                     `json:"zip"                                column:"clients.zip"                                     sql:"VARCHAR NOT NULL"`
	PhoneNumber                        string                     `json:"phoneNumber"                        column:"clients.phone_number"                            sql:"VARCHAR NOT NULL"`
	Fax                                *string                    `json:"fax"                                column:"clients.fax"                                     sql:"VARCHAR NULL"`
	Address                            string                     `json:"address"                            column:"clients.address"                                 sql:"VARCHAR NOT NULL"`
	ContactPersonName                  string                     `json:"contactPersonName"                  column:"clients.contact_person_name"                     sql:"VARCHAR NOT NULL"`
	ContactPersonEmail                 string                     `json:"contactPersonEmail"                 column:"clients.contact_person_email"                    sql:"VARCHAR NOT NULL"`
	PendingEmailChange                 *string                    `json:"pendingEmailChange"                 column:"clients.pending_email_change"                    sql:"VARCHAR NULL DEFAULT NULL"`
	Image                              *string                    `json:"image"                              column:"clients.image"                                   sql:"VARCHAR NULL DEFAULT NULL"`
	GroupID                            int                        `json:"groupId"                            column:"clients.group_id"                                sql:"BIGINT NULL"`
	ModifiedByID                       int                        `json:"modifiedById"                       column:"clients.modified_by_id"                          sql:"BIGINT NULL"`
	PartnerID                          int                        `json:"partnerId"                          column:"clients.partner_id"                              sql:"BIGINT NULL"`
	ExpiredAt                          time.Time                  `json:"expiredAt"                          column:"clients.expired_at"                              sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	CreatedAt                          time.Time                  `json:"createdAt"                          column:"clients.created_at"                              sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt                          time.Time                  `json:"updatedAt"                          column:"clients.updated_at"                              sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	OpenAPIStatus                      *CompanyOpenAPIStatus      `json:"openApiStatus"                      column:"clients.open_api_status"                         sql:"open_api_status_enum NULL DEFAULT NULL"`
	OpenAPIRequestedAt                 *time.Time                 `json:"openApiRequestedAt"                 column:"clients.open_api_requested_at"                   sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	OpenAPIApprovedAt                  *time.Time                 `json:"openApiApprovedAt"                  column:"clients.open_api_approved_at"                    sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	PlanFixedAsset                     fixedasset_model.PlanTypes `json:"planFixedAsset"                     column:"clients.plan_fixed_asset"                        sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanSupplyAsset                    fixedasset_model.PlanTypes `json:"planSupplyAsset"                    column:"clients.plan_supply_asset"                       sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanInventory                      fixedasset_model.PlanTypes `json:"planInventory"                      column:"clients.plan_inventory_asset"                    sql:"license_plan_enum DEFAULT 'Basic'"`
	PlanProcurementAsset               fixedasset_model.PlanTypes `json:"planProcurementAsset"               column:"clients.plan_procurement_asset"                  sql:"license_plan_enum DEFAULT 'Basic'"`
	DepreciationMethod                 CompanyDepreciationMethod  `json:"depreciationMethod"                 column:"clients.depreciation_method"                     sql:"depreciation_method_enum NOT NULL DEFAULT 'Straight-Line'"`
	PendingDepreciationMethod          *CompanyDepreciationMethod `json:"pendingDepreciationMethod"          column:"clients.pending_depreciation_method"             sql:"depreciation_method_enum NULL DEFAULT NULL"`
	DepreciationMethodUpdatedAt        *time.Time                 `json:"depreciationMethodUpdatedAt"        column:"clients.depreciation_method_updated_at"          sql:"TIMESTAMPTZ NULL DEFAULT NULL"`
	AddOnRepairTicketing               bool                       `json:"addOnRepairTicketing"               column:"clients.addon_repair_ticketing"                  sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnMaintenance                   bool                       `json:"addOnMaintenance"                   column:"clients.addon_maintenance"                       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnTracking                      bool                       `json:"addOnTracking"                      column:"clients.addon_tracking"                          sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAssetControl                  bool                       `json:"addOnAssetControl"                  column:"clients.addon_asset_control"                     sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AddOnAudit                         bool                       `json:"addOnAudit"                         column:"clients.addon_audit"                             sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	AuditAttrsMaxGracePeriod           int                        `json:"auditAttrsMaxGracePeriod"           column:"clients.audit_attrs_max_grace_period"            sql:"INTEGER DEFAULT 3"`
	TransferAttrsTransactionTimeout    int                        `json:"transferAttrsTransactionTimeout"    column:"clients.transfer_attrs_transaction_timeout"      sql:"INTEGER DEFAULT 5"`
	TransferAttrsAssetControlToMissing int                        `json:"transferAttrsAssetControlToMissing" column:"clients.transfer_attrs_asset_control_to_missing" sql:"INTEGER DEFAULT 1"`
	BorrowingAttrsMaxBorrowingTime     int                        `json:"borrowingAttrsMaxBorrowingTime"     column:"clients.borrowing_attrs_max_borrowing_time"      sql:"INTEGER DEFAULT 15"`
	BorrowingAttrsDefaultBorrowingTime int                        `json:"borrowingAttrsDefaultBorrowingTime" column:"clients.borrowing_attrs_default_borrowing_time"  sql:"INTEGER DEFAULT 7"`
	BorrowingAttrsConfirmationTimeout  int                        `json:"borrowingAttrsConfirmationTimeout"  column:"clients.borrowing_attrs_confirmation_timeout"    sql:"INTEGER DEFAULT 24"`
	BorrowingAttrsTransactionTimeout   int                        `json:"borrowingAttrsTransactionTimeout"   column:"clients.borrowing_attrs_transaction_timeout"     sql:"INTEGER DEFAULT 5"`
	DisposalAttrsTransactionTimeout    int                        `json:"disposalAttrsTransactionTimeout"    column:"clients.disposal_attrs_transaction_timeout"      sql:"INTEGER DEFAULT 5"`
	LoginSessionMinutes                int                        `json:"loginSessionMinutes"                column:"clients.login_session_minutes"                   sql:"INTEGER DEFAULT 30"`
	DateFormat                         string                     `json:"dateFormat"                         column:"clients.date_format"                             sql:"VARCHAR DEFAULT '02/01/06'"`
	TimeFormat                         string                     `json:"timeFormat"                         column:"clients.time_format"                             sql:"VARCHAR NOT NULL DEFAULT '15:04:05'"`
	Geolocation                        bool                       `json:"geolocation"                        column:"clients.geolocation"                             sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TimeZone                           string                     `json:"timeZone"                           column:"clients.time_zone"                               sql:"VARCHAR default 'Asia/Jakarta'"`
	PIN                                string                     `json:"pin"                                column:"clients.pin"                                     sql:"VARCHAR(6) DEFAULT '123456'"`
	CurrencySymbol                     string                     `json:"currencySymbol"                     column:"clients.currency_symbol"                         sql:"VARCHAR DEFAULT 'Rp'"`
	CurrencyCurrency                   string                     `json:"currencyCurrency"                   column:"clients.currency_currency"                       sql:"VARCHAR DEFAULT 'IDR'"`
	CurrencyLocale                     string                     `json:"currencyLocale"                     column:"clients.currency_locale"                         sql:"VARCHAR DEFAULT 'id-ID'"`
	CurrencyLabel                      string                     `json:"currencyLabel"                      column:"clients.currency_label"                          sql:"VARCHAR DEFAULT 'Rp - Indonesian Rupiah'"`
	FileStorageTotalKB                 float64                    `json:"fileStorageTotalKB"                 column:"clients.file_storage_total_kb"                   sql:"FLOAT DEFAULT 5242880"` // 5 gb
	FileStorageAvailableKB             float64                    `json:"fileStorageAvailableKB"             column:"clients.file_storage_available_kb"               sql:"FLOAT DEFAULT 5242880"` // 5gb
	FileStorageUsedKB                  float64                    `json:"fileStorageUsedKB"                  column:"clients.file_storage_used_kb"                    sql:"FLOAT DEFAULT 0"`
	FileStorageAttachmentKB            float64                    `json:"fileStorageAttachmentKB"            column:"clients.file_storage_attachment_kb"              sql:"FLOAT DEFAULT 0"`
	FileStorageDamagedKB               float64                    `json:"fileStorageDamagedKB"               column:"clients.file_storage_damage_kb"                  sql:"FLOAT DEFAULT 0"`
	FileStorageDisposalKB              float64                    `json:"fileStorageDisposalKB"              column:"clients.file_storage_disposal_kb"                sql:"FLOAT DEFAULT 0"`
	FileStorageRepairTicketingKB       float64                    `json:"fileStorageRepairTicketingKB"       column:"clients.file_storage_repair_ticketing_kb"        sql:"FLOAT DEFAULT 0"`
	FileStorageRoutineKB               float64                    `json:"fileStorageRoutineKB"               column:"clients.file_storage_routine_kb"                 sql:"FLOAT DEFAULT 0"`
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

	// ─────── Foreign Keys ─────── (Not priority for now)
	// builder.AddForeignKey("fk_client_group", "group_id", "groups", "id", "SET NULL")
	// builder.AddForeignKey("fk_client_modified_by_id", "modified_by_id", "users", "id", "SET NULL")
	// builder.AddForeignKey("fk_client_partner_id", "partner_id", "partners", "id", "SET NULL")

	// ─────── Additional Queries ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(admin_migration.ClientPostQuery)
}

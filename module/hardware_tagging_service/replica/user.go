package replica

import "time"

type User struct {
	ID                         int        `json:"_id" column:"users.id" sqlType:"BIGINT"`
	IsActiveGlobalSettings     bool       `json:"isActiveGlobalSettings" column:"users.is_active_global_settings" sqlType:"BOOL"`
	IsActiveFixedAsset         bool       `json:"isActiveFixedAsset" column:"users.is_active_fixed_asset" sqlType:"BOOL"`
	IsActiveSupplyAsset        bool       `json:"isActiveSupplyAsset" column:"users.is_active_supply_asset" sqlType:"BOOL"`
	IsActiveAdminConsole       bool       `json:"isActiveAdminConsole" column:"users.is_active_admin_console" sqlType:"BOOL"`
	IsActiveInventory          bool       `json:"isActiveInventory" column:"users.is_active_inventory" sqlType:"BOOL"`
	IsActiveProcurement        bool       `json:"isActiveProcurement" column:"users.is_active_procurement" sqlType:"BOOL"`
	IsActive                   bool       `json:"isActive" column:"users.is_active" sqlType:"BOOL"`
	IsTotalControlFixedAsset   bool       `json:"isTotalControlFixedAsset" column:"users.is_total_control_fixed_asset" sqlType:"BOOL"`
	IsTotalControlSupplyAsset  bool       `json:"isTotalControlSupplyAsset" column:"users.is_total_control_supply_asset" sqlType:"BOOL"`
	IsTotalControlAdminConsole bool       `json:"isTotalControlAdminConsole" column:"users.is_total_control_admin_console" sqlType:"BOOL"`
	IsTotalControlInventory    bool       `json:"isTotalControlInventory" column:"users.is_total_control_inventory" sqlType:"BOOL"`
	IsTotalControlProcurement  bool       `json:"isTotalControlProcurement" column:"users.is_total_control_procurement" sqlType:"BOOL"`
	IsReadOnlyFixedAsset       bool       `json:"isReadOnlyFixedAsset" column:"users.is_read_only_fixed_asset" sqlType:"BOOL"`
	IsReadOnlySupplyAsset      bool       `json:"isReadOnlySupplyAsset" column:"users.is_read_only_supply_asset" sqlType:"BOOL"`
	IsReadOnlyAdminConsole     bool       `json:"isReadOnlyAdminConsole" column:"users.is_read_only_admin_console" sqlType:"BOOL"`
	IsReadOnlyInventory        bool       `json:"isReadOnlyInventory" column:"users.is_read_only_inventory" sqlType:"BOOL"`
	IsReadOnlyProcurement      bool       `json:"isReadOnlyProcurement" column:"users.is_read_only_procurement" sqlType:"BOOL"`
	IsDefault                  bool       `json:"isDefault" column:"users.is_default" sqlType:"BOOL"`
	IsEmailConfirmed           bool       `json:"isEmailConfirmed" column:"users.is_email_confirmed" sqlType:"BOOL"`
	IsSubUser                  bool       `json:"isSubUser" column:"users.is_sub_user" sqlType:"BOOL"`
	FirstName                  string     `json:"firstName" column:"users.first_name" sqlType:"VARCHAR"`
	LastName                   string     `json:"lastName" column:"users.last_name" sqlType:"VARCHAR"`
	FullName                   string     `json:"fullName" column:"users.full_name" sqlType:"VARCHAR"`
	Email                      string     `json:"email" column:"users.email" sqlType:"VARCHAR"`
	EmployeeID                 *string    `json:"employeeId" column:"users.employee_id" sqlType:"VARCHAR"`
	ProfilePicture             *string    `json:"profilePicture" column:"users.profile_picture" sqlType:"VARCHAR"`
	FcmToken                   *string    `json:"fcmToken" column:"users.fcm_token" sqlType:"VARCHAR"`
	Access                     []string   `json:"access" column:"users.access" sqlType:"JSONB"`
	PositionID                 *int       `json:"positionId" column:"users.position_id" sqlType:"BIGINT"`
	DivisionID                 *int       `json:"divisionId" column:"users.division_id" sqlType:"BIGINT"`
	PendingEmailChange         *string    `json:"pendingEmailChange" column:"users.pending_email_change" sqlType:"VARCHAR"`
	ExpiryDate                 *time.Time `json:"expiryDate" column:"users.expiry_date" sqlType:"TIMESTAMPTZ"`
	RfidID                     *int       `json:"rfidId" column:"users.rfid_id" sqlType:"BIGINT"`
	ParentID                   *int       `json:"parentId" column:"users.parent_id" sqlType:"BIGINT"`
	ParentName                 *string    `json:"parentName" column:"users.parent_name" sqlType:"VARCHAR"`
	ModifiedByID               *int       `json:"modifiedById" column:"users.modified_by_id" sqlType:"BIGINT"`
	IsDeleted                  bool       `json:"isDeleted" column:"users.is_deleted" sqlType:"BOOL"`
	DeletedAt                  *time.Time `json:"deletedAt" column:"users.deleted_at" sqlType:"TIMESTAMPTZ"`
	CreatedAt                  time.Time  `json:"createdAt" column:"users.created_at" sqlType:"TIMESTAMPTZ"`
	UpdatedAt                  time.Time  `json:"updatedAt" column:"users.updated_at" sqlType:"TIMESTAMPTZ"`
	IsConsented                bool       `json:"isConsented" column:"users.is_consented" sqlType:"BOOL"`
	ConsentDate                *time.Time `json:"consentDate" column:"users.consent_date" sqlType:"TIMESTAMPTZ"`
}

func (r *User) GetTableName() string {
	return "users"
}

func (r *User) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                             "BIGINT",
		"is_active_global_settings":      "BOOL",
		"is_active_fixed_asset":          "BOOL",
		"is_active_supply_asset":         "BOOL",
		"is_active_admin_console":        "BOOL",
		"is_active_inventory":            "BOOL",
		"is_active_procurement":          "BOOL",
		"is_active":                      "BOOL",
		"is_total_control_fixed_asset":   "BOOL",
		"is_total_control_supply_asset":  "BOOL",
		"is_total_control_admin_console": "BOOL",
		"is_total_control_inventory":     "BOOL",
		"is_total_control_procurement":   "BOOL",
		"is_read_only_fixed_asset":       "BOOL",
		"is_read_only_supply_asset":      "BOOL",
		"is_read_only_admin_console":     "BOOL",
		"is_read_only_inventory":         "BOOL",
		"is_read_only_procurement":       "BOOL",
		"is_default":                     "BOOL",
		"is_email_confirmed":             "BOOL",
		"is_sub_user":                    "BOOL",
		"first_name":                     "VARCHAR",
		"last_name":                      "VARCHAR",
		"full_name":                      "VARCHAR",
		"email":                          "VARCHAR",
		"employee_id":                    "VARCHAR",
		"profile_picture":                "VARCHAR",
		"fcm_token":                      "VARCHAR",
		"access":                         "JSONB",
		"position_id":                    "BIGINT",
		"division_id":                    "BIGINT",
		"pending_email_change":           "VARCHAR",
		"expiry_date":                    "TIMESTAMPTZ",
		"rfid_id":                        "BIGINT",
		"parent_id":                      "BIGINT",
		"parent_name":                    "VARCHAR",
		"modified_by_id":                 "BIGINT",
		"is_deleted":                     "BOOL",
		"deleted_at":                     "TIMESTAMPTZ",
		"created_at":                     "TIMESTAMPTZ",
		"updated_at":                     "TIMESTAMPTZ",
		"is_consented":                   "BOOL",
		"consent_date":                   "TIMESTAMPTZ",
	}
}

func (r *User) GetColumns() []string {
	return []string{
		"id",
		"is_active_global_settings",
		"is_active_fixed_asset",
		"is_active_supply_asset",
		"is_active_admin_console",
		"is_active_inventory",
		"is_active_procurement",
		"is_active",
		"is_total_control_fixed_asset",
		"is_total_control_supply_asset",
		"is_total_control_admin_console",
		"is_total_control_inventory",
		"is_total_control_procurement",
		"is_read_only_fixed_asset",
		"is_read_only_supply_asset",
		"is_read_only_admin_console",
		"is_read_only_inventory",
		"is_read_only_procurement",
		"is_default",
		"is_email_confirmed",
		"is_sub_user",
		"first_name",
		"last_name",
		"full_name",
		"email",
		"employee_id",
		"profile_picture",
		"fcm_token",
		"access",
		"position_id",
		"division_id",
		"pending_email_change",
		"expiry_date",
		"rfid_id",
		"parent_id",
		"parent_name",
		"modified_by_id",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
		"is_consented",
		"consent_date",
	}
}

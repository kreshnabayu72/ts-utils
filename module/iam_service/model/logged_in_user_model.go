package model

import (
	"time"
)

type LoggedInUser struct {
	ID                         int        `json:"_id"                        column:"users.id"`
	IsActiveGlobalSettings     bool       `json:"isActiveGlobalSettings"     column:"users.is_active_global_settings"`
	IsActiveFixedAsset         bool       `json:"isActiveFixedAsset"         column:"users.is_active_fixed_asset"`
	IsActiveSupplyAsset        bool       `json:"isActiveSupplyAsset"        column:"users.is_active_supply_asset"`
	IsActiveAdminConsole       bool       `json:"isActiveAdminConsole"       column:"users.is_active_admin_console"`
	IsActiveInventory          bool       `json:"isActiveInventory"          column:"users.is_active_inventory"`
	IsActive                   bool       `json:"isActive"                   column:"users.is_active"                      special:"generated"`
	IsTotalControlFixedAsset   bool       `json:"isTotalControlFixedAsset"   column:"users.is_total_control_fixed_asset"`
	IsReadOnlyFixedAsset       bool       `json:"isReadOnlyFixedAsset"       column:"users.is_read_only_fixed_asset"`
	IsTotalControlSupplyAsset  bool       `json:"isTotalControlSupplyAsset"  column:"users.is_total_control_supply_asset"`
	IsReadOnlySupplyAsset      bool       `json:"isReadOnlySupplyAsset"      column:"users.is_read_only_supply_asset"`
	IsTotalControlAdminConsole bool       `json:"isTotalControlAdminConsole" column:"users.is_total_control_admin_console"`
	IsReadOnlyAdminConsole     bool       `json:"isReadOnlyAdminConsole"     column:"users.is_read_only_admin_console"`
	IsTotalControlInventory    bool       `json:"isTotalControlInventory"    column:"users.is_total_control_inventory"`
	IsReadOnlyInventory        bool       `json:"isReadOnlyInventory"        column:"users.is_read_only_inventory"`
	IsDefault                  bool       `json:"isDefault"                  column:"users.is_default"`
	IsEmailConfirmed           bool       `json:"isEmailConfirmed"           column:"users.is_email_confirmed"`
	IsSubUser                  bool       `json:"isSubUser"                  column:"users.is_sub_user"                    special:"generated"`
	FirstName                  string     `json:"firstName"                  column:"users.first_name"`
	LastName                   string     `json:"lastName"                   column:"users.last_name"`
	Email                      string     `json:"email"                      column:"users.email"`
	EmployeeId                 *string    `json:"employeeId"                 column:"users.employee_id"`
	FullName                   string     `json:"fullName"                   column:"users.full_name"                      special:"generated"`
	ProfilePicture             *string    `json:"profilePicture"             column:"users.profile_picture"`
	FCMToken                   *string    `json:"fcmToken"                   column:"users.fcm_token"`
	PositionId                 *int       `json:"positionId"                 column:"users.position_id"`
	DivisionId                 *int       `json:"divisionId"                 column:"users.division_id"`
	PendingEmailChange         *string    `json:"pendingEmailChange"         column:"users.pending_email_change"`
	ExpiryDate                 *time.Time `json:"expiryDate"                 column:"users.expiry_date"`
	RfidId                     *int       `json:"rfidId"                     column:"users.rfid_id"`
	ParentId                   *int       `json:"parentId"                   column:"users.parent_id"`
	ParentName                 *string    `json:"parentName"                 column:"users.parent_name"`
	IsConsented                bool       `json:"isConsented"                column:"users.is_consented"`
	ConsentDate                *time.Time `json:"consentDate"                column:"users.consent_date"`
	ModifiedById               *int       `json:"modifiedById"               column:"users.modified_by_id"`
	Access                     []string   `json:"access"                     column:"users.access"`
	IsDeleted                  bool       `json:"isDeleted"                  column:"users.is_deleted"`
	CreatedAt                  time.Time  `json:"createdAt"                  column:"users.created_at"`
	UpdatedAt                  time.Time  `json:"updatedAt"                  column:"users.updated_at"`
}

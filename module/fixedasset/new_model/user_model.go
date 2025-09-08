package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_migration "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model/migration"
)

type User struct {
	ID                     int        `json:"_id"                    column:"users.id"                        sql:"BIGINT PRIMARY KEY"`
	FCMToken               string     `json:"fcmToken"               column:"users.fcm_token"                 sql:"VARCHAR NULL"`
	FirstName              string     `json:"firstName"              column:"users.first_name"                sql:"VARCHAR NOT NULL"`
	LastName               string     `json:"lastName"               column:"users.last_name"                 sql:"VARCHAR NOT NULL"`
	FullName               string     `json:"fullName"               column:"users.full_name"                 sql:"VARCHAR NOT NULL GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED"`
	Email                  string     `json:"email"                  column:"users.email"                     sql:"VARCHAR NOT NULL"`
	PhoneNumber            string     `json:"phoneNumber"            column:"users.phone_number"              sql:"VARCHAR NOT NULL"`
	PendingEmailChange     string     `json:"pendingEmailChange"     column:"users.pending_email_change"      sql:"VARCHAR NULL DEFAULT NULL"`
	IsActive               bool       `json:"isActive"               column:"users.is_active"                 sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsTemporary            bool       `json:"isTemporary"            column:"users.is_temporary"              sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsDefault              bool       `json:"isDefault"              column:"users.is_default"                sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsEmailConfirmed       bool       `json:"isEmailConfirmed"       column:"users.is_email_confirmed"        sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ExpiryDate             *time.Time `json:"expiryDate"             column:"users.expiry_date"               sql:"TIMESTAMPTZ NULL"`
	IsExpired              *bool      `json:"isExpired"              column:"CASE WHEN expiry_date IS NULL THEN NULL ELSE expiry_date < NOW() END"`
	Position               *string    `json:"position"               column:"users.position"                  sql:"VARCHAR NULL"`
	Division               string     `json:"division"               column:"users.division"                  sql:"VARCHAR NOT NULL"`
	RfidCode               *string    `json:"rfidCode"               column:"users.rfid_code"                 sql:"VARCHAR DEFAULT NULL"`
	RfidId                 *int       `json:"rfidId"                 column:"users.rfid_id"                   sql:"BIGINT DEFAULT NULL"`
	ProfilePicture         *string    `json:"profilePicture"         column:"users.profile_picture"           sql:"VARCHAR DEFAULT NULL"`
	EmployeeId             *string    `json:"employeeId"             column:"users.employee_id"               sql:"VARCHAR DEFAULT NULL"`
	IsSubUser              bool       `json:"isSubUser"              column:"users.is_sub_user"               sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ParentId               *int       `json:"parentId"               column:"users.parent_id"                 sql:"BIGINT REFERENCES users (id) ON DELETE CASCADE NULL"`
	ReadOnlyIsActive       bool       `json:"readOnlyIsActive"       column:"users.read_only_is_active"       sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ReadOnlyIsAssigned     bool       `json:"readOnlyIsAssigned"     column:"users.read_only_is_assigned"     sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TotalControlIsActive   bool       `json:"totalControlIsActive"   column:"users.total_control_is_active"   sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TotalControlIsAssigned bool       `json:"totalControlIsAssigned" column:"users.total_control_is_assigned" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ModifiedById           int        `json:"modifiedById"           column:"users.modified_by_id"            sql:"BIGINT NOT NULL"`
	ModifiedByFirstName    string     `json:"modifiedByFirstName"    column:"users.modified_by_first_name"    sql:"VARCHAR NOT NULL"`
	ModifiedByLastName     string     `json:"modifiedByLastName"     column:"users.modified_by_last_name"     sql:"VARCHAR NOT NULL"`
	ModifiedByFullName     string     `json:"modifiedByFullName"     column:"users.modified_by_full_name"     sql:"VARCHAR NOT NULL GENERATED ALWAYS AS (modified_by_first_name || ' ' || modified_by_last_name) STORED"`
	HasModuleAccess        bool       `json:"hasModuleAccess"        column:"users.has_module_access"         sql:"BOOLEAN NOT NULL DEFAULT FALSE"` //For showing/hiding user that doesnt have fixed asset access from global settings
	IsDeleted              bool       `json:"isDeleted"              column:"users.is_deleted"                sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeleteFromDatabaseAt   *time.Time `json:"deleteFromDatabaseAt"   column:"users.deleted_at"                sql:"TIMESTAMPTZ NULL"`
	CreatedAt              time.Time  `json:"createdAt"              column:"users.created_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt              time.Time  `json:"updatedAt"              column:"users.updated_at"                sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

type UserAttribute struct {
	ID        string `json:"_id,omitempty" column:"id::text"`
	FullName  string `json:"fullName"      column:"concat_ws(' ', first_name, last_name)"`
	IsDefault bool   `json:"isDefault"     column:"is_default"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("users_email_uindex", "email", true, "")
	builder.AddIndex("users_rfid_code_uindex", "rfid_code", true, "")

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(fixed_asset_migration.UserPostQuery)
}

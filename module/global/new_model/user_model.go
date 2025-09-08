package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
	fixed_asset_model "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	global_migration "gitea.qwertysystem.net/BETS/ts-utils/module/global/new_model/migration"
)

type User struct {
	ID int `json:"_id" column:"users.id" sql:"BIGINT PRIMARY KEY"`

	FirstName          string  `json:"firstName" column:"users.first_name" sql:"VARCHAR NOT NULL"`
	LastName           string  `json:"lastName" column:"users.last_name" sql:"VARCHAR NOT NULL"`
	FullName           string  `json:"fullName" column:"users.full_name" sql:"VARCHAR GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED"`
	PhoneNumber        string  `json:"phoneNumber" column:"users.phone_number" sql:"VARCHAR NOT NULL"`
	Email              string  `json:"email" column:"users.email" sql:"VARCHAR NOT NULL"`
	PendingEmailChange *string `json:"pendingEmailChange" column:"users.pending_email_change" sql:"VARCHAR NULL"`
	IsActive           bool    `json:"isActive" column:"users.is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsDefault          bool    `json:"isDefault" column:"users.is_default" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsDeleted          bool    `json:"isDeleted" column:"users.is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsEmailConfirmed   bool    `json:"isEmailConfirmed" column:"users.is_email_confirmed" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`

	PositionID      *int                             `json:"positionId" column:"users.position_id" sql:"BIGINT NULL"`
	DivisionID      int                              `json:"divisionId" column:"users.division_id" sql:"BIGINT"`
	RfidCode        *string                          `json:"rfidCode" column:"users.rfid_code" sql:"VARCHAR NULL"`
	ProfilePicture  *string                          `json:"profilePicture" column:"users.profile_picture" sql:"VARCHAR NULL"`
	EmployeeID      *string                          `json:"employeeId" column:"users.employee_id" sql:"VARCHAR NULL"`
	Access          []fixed_asset_model.ModuleAccess `json:"access" column:"users.access" sql:"TEXT[]"`
	ModifiedBy      int                              `json:"modifiedById" column:"users.modified_by_id" sql:"BIGINT REFERENCES users(id)"`
	IsTemporary     bool                             `json:"isTemporary" column:"users.is_temporary" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ExpiryDate      *time.Time                       `json:"expiryDate" column:"users.expiry_date" sql:"TIMESTAMPTZ NULL"`
	HasModuleAccess bool                             `json:"hasModuleAccess" column:"users.has_module_access" sql:"BOOLEAN NOT NULL DEFAULT FALSE"` //For showing/hiding user that doesnt have fixed asset access from global settings

	FCMToken string `json:"fcmToken"    column:"users.fcm_token" sql:"VARCHAR NULL"`

	CreatedAt *time.Time `json:"createdAt" column:"users.created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt *time.Time `json:"updatedAt" column:"users.updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddForeignKey("fk_global_user_position", "position_id", "positions", "id", "SET NULL")
	builder.AddForeignKey("fk_global_user_division", "division_id", "divisions", "id", "CASCADE")
	builder.AddForeignKey("fk_global_user_modified_by", "modified_by_id", "users", "id", "SET NULL")

	builder.AddIndex("users_email_uindex", "email", true, "")

	// ─────── Insert default data ───────
	// NOTE: reading from sql file has been tried and doesnt work in server, so use string instead
	builder.SetPostQuery(global_migration.UserPostQuery)
}

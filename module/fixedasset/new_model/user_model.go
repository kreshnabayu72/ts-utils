package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type User struct {
	ID                     int        `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	FCMToken               string     `json:"fcmToken"    column:"fcm_token" sql:"VARCHAR NULL"`
	FirstName              string     `json:"firstName" column:"first_name" sql:"VARCHAR NOT NULL"`
	LastName               string     `json:"lastName" column:"last_name" sql:"VARCHAR NOT NULL"`
	FullName               string     `json:"fullName" column:"full_name" sql:"VARCHAR NOT NULL GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED"`
	Email                  string     `json:"email" column:"email" sql:"VARCHAR NOT NULL"`
	PhoneNumber            string     `json:"phoneNumber" column:"phone_number" sql:"VARCHAR NOT NULL"`
	PendingEmailChange     string     `json:"pendingEmailChange" column:"pending_email_change" sql:"VARCHAR NULL DEFAULT NULL"`
	IsActive               bool       `json:"isActive" column:"is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsTemporary            bool       `json:"isTemporary" column:"is_temporary" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsDefault              bool       `json:"isDefault" column:"is_default" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsEmailConfirmed       bool       `json:"isEmailConfirmed" column:"is_email_confirmed" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ExpiryDate             *time.Time `json:"expiryDate" column:"expiry_date" sql:"TIMESTAMPTZ NULL"`
	IsExpired              *bool      `json:"isExpired" column:"CASE WHEN expiry_date IS NULL THEN NULL ELSE expiry_date < NOW() END"`
	Position               *string    `json:"position" column:"position" sql:"VARCHAR NULL"`
	Division               string     `json:"division" column:"division" sql:"VARCHAR NOT NULL"`
	RfidCode               *string    `json:"rfidCode" column:"rfid_code" sql:"VARCHAR DEFAULT NULL"`
	RfidId                 *int       `json:"rfidId" column:"rfid_id" sql:"BIGINT DEFAULT NULL"`
	ProfilePicture         *string    `json:"profilePicture" column:"profile_picture" sql:"VARCHAR DEFAULT NULL"`
	EmployeeId             *string    `json:"employeeId" column:"employee_id" sql:"VARCHAR DEFAULT NULL"`
	IsSubUser              bool       `json:"isSubUser" column:"is_sub_user" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ParentId               *int       `json:"parentId" column:"parent_id" sql:"BIGINT REFERENCES users (id) ON DELETE CASCADE NULL"`
	ReadOnlyIsActive       bool       `json:"readOnlyIsActive" column:"read_only_is_active" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ReadOnlyIsAssigned     bool       `json:"readOnlyIsAssigned" column:"read_only_is_assigned" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TotalControlIsActive   bool       `json:"totalControlIsActive" column:"total_control_is_active" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	TotalControlIsAssigned bool       `json:"totalControlIsAssigned" column:"total_control_is_assigned" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	ModifiedById           int        `json:"modifiedById" column:"modified_by_id" sql:"BIGINT NOT NULL"`
	ModifiedByFirstName    string     `json:"modifiedByFirstName" column:"modified_by_first_name" sql:"VARCHAR NOT NULL"`
	ModifiedByLastName     string     `json:"modifiedByLastName" column:"modified_by_last_name" sql:"VARCHAR NOT NULL"`
	ModifiedByFullName     string     `json:"modifiedByFullName" column:"modified_by_full_name" sql:"VARCHAR NOT NULL GENERATED ALWAYS AS (modified_by_first_name || ' ' || modified_by_last_name) STORED"`
	IsDeleted              bool       `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	DeleteFromDatabaseAt   *time.Time `json:"deleteFromDatabaseAt" column:"deleted_at" sql:"TIMESTAMPTZ NULL"`
	CreatedAt              time.Time  `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt              time.Time  `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
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
}

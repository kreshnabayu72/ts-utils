package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type User struct {
	ID             int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	FirstName      string    `json:"firstName" column:"first_name" sql:"VARCHAR NOT NULL"`
	LastName       string    `json:"lastName" column:"last_name" sql:"VARCHAR NOT NULL"`
	FullName       string    `json:"fullName" column:"full_name" sql:"VARCHAR NOT NULL GENERATED ALWAYS AS (first_name || ' ' || last_name) STORED"`
	Email          string    `json:"email" column:"email" sql:"VARCHAR NOT NULL"`
	IsActive       bool      `json:"isActive" column:"is_active" sql:"BOOLEAN NOT NULL DEFAULT TRUE"`
	IsDefault      bool      `json:"isDefault" column:"is_default" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	IsDeleted      bool      `json:"isDeleted" column:"is_deleted" sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	Position       *string   `json:"position" column:"position" sql:"VARCHAR NULL"`
	Division       string    `json:"division" column:"division" sql:"VARCHAR NOT NULL"`
	RFIDCode       *string   `json:"rfidCode" column:"rfid_code" sql:"VARCHAR NULL"`
	ProfilePicture *string   `json:"profilePicture" column:"profile_picture" sql:"VARCHAR NULL"`
	EmployeeID     *string   `json:"employeeId" column:"employee_id" sql:"VARCHAR NULL"`
	ModifiedByID   int       `json:"modifiedById" column:"modified_by_id" sql:"BIGINT REFERENCES users(id)"`
	CreatedAt      time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
	UpdatedAt      time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMP NOT NULL DEFAULT NOW()"`
}

func (m *User) TableName() string {
	return "users"
}

func (m *User) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	builder.AddIndex("users_email_uindex", "email", true, "is_deleted = FALSE")
}

package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportUser struct {
	ID int `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`

	Key         int            `json:"key" column:"key" sql:"INTEGER NOT NULL"`
	FirstName   string         `json:"firstName" column:"first_name" sql:"VARCHAR NOT NULL"`
	LastName    string         `json:"lastName" column:"last_name" sql:"VARCHAR NOT NULL"`
	Position    string         `json:"position" column:"position" sql:"VARCHAR NOT NULL"`
	Division    string         `json:"division" column:"division" sql:"VARCHAR NOT NULL"`
	PhoneNumber string         `json:"phoneNumber" column:"phone_number" sql:"VARCHAR NOT NULL"`
	Access      []ModuleAccess `json:"access" column:"access" sql:"module_access_enum[]"`
	Email       string         `json:"email" column:"email" sql:"VARCHAR NOT NULL"`
	EmployeeId  string         `json:"employeeId" column:"employee_id" sql:"VARCHAR NOT NULL"`

	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (u *ImportUser) TableName() string {
	return "import_users"
}

func (u *ImportUser) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(u)

	builder.AddIndex("import_users_email", "LOWER(email)", true, "")
}

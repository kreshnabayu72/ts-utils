package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type ImportUser struct {
	ID          int            `json:"_id"         column:"import_users.id"           sql:"BIGINT PRIMARY KEY"`
	Key         int            `json:"key"         column:"import_users.key"          sql:"INTEGER NOT NULL"`
	FirstName   string         `json:"firstName"   column:"import_users.first_name"   sql:"VARCHAR NOT NULL"`
	LastName    string         `json:"lastName"    column:"import_users.last_name"    sql:"VARCHAR NOT NULL"`
	Position    string         `json:"position"    column:"import_users.position"     sql:"VARCHAR NOT NULL"`
	Division    string         `json:"division"    column:"import_users.division"     sql:"VARCHAR NOT NULL"`
	PhoneNumber string         `json:"phoneNumber" column:"import_users.phone_number" sql:"VARCHAR NOT NULL"`
	Access      []ModuleAccess `json:"access"      column:"import_users.access"       sql:"module_access_enum[]"`
	Email       string         `json:"email"       column:"import_users.email"        sql:"VARCHAR NOT NULL"`
	EmployeeId  string         `json:"employeeId"  column:"import_users.employee_id"  sql:"VARCHAR NOT NULL"`
	CreatedAt   time.Time      `json:"createdAt"   column:"import_users.created_at"   sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt   time.Time      `json:"updatedAt"   column:"import_users.updated_at"   sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (u *ImportUser) TableName() string {
	return "import_users"
}

func (u *ImportUser) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(u)

	builder.AddIndex("import_users_email", "LOWER(email)", true, "")
}

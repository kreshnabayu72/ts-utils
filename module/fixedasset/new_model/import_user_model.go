package new_model

import (
	"time"
)

type ImportUser struct {
	ID          int       `json:"_id"         column:"import_users.id"`
	Key         int       `json:"key"         column:"import_users.key"`
	FirstName   string    `json:"firstName"   column:"import_users.first_name"`
	LastName    string    `json:"lastName"    column:"import_users.last_name"`
	Position    string    `json:"position"    column:"import_users.position"`
	Division    string    `json:"division"    column:"import_users.division"`
	PhoneNumber string    `json:"phoneNumber" column:"import_users.phone_number"`
	Access      string    `json:"access"      column:"import_users.access"`
	Email       string    `json:"email"       column:"import_users.email"`
	EmployeeId  string    `json:"employeeId"  column:"import_users.employee_id"`
	CreatedAt   time.Time `json:"createdAt"   column:"import_users.created_at"`
	UpdatedAt   time.Time `json:"updatedAt"   column:"import_users.updated_at"`
}

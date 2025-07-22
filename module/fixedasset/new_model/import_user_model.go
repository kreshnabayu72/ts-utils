package new_model

import "time"

type ImportUser struct {
	ID int `json:"_id"                  column:"id"`

	Key         int            `json:"key"                 column:"key"`
	FirstName   string         `json:"firstName"                column:"first_name"`
	LastName    string         `json:"lastName"                column:"last_name"`
	Position    string         `json:"position"                column:"position"`
	Division    string         `json:"division"                column:"division"`
	PhoneNumber string         `json:"phoneNumber"                column:"phone_number"`
	Access      []ModuleAccess `json:"access"                column:"access"`
	Email       string         `json:"email"                column:"email"`
	EmployeeId  string         `json:"employeeId"                column:"employeeId"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}

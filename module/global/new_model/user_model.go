package new_model

import (
	"time"
)

type User struct {
	ID int `json:"_id" column:"id"`

	FirstName string `json:"firstName" column:"first_name"`
	LastName  string `json:"lastName" column:"last_name"`
	Email     string `json:"email" column:"email"`
	IsActive  bool   `json:"isActive" column:"is_active"`
	IsDefault bool   `json:"isDefault" column:"is_default"`

	PositionID     *int    `json:"positionId" column:"position_id"`
	DivisionID     int     `json:"divisionId" column:"division_id"`
	RfidCode       *string `json:"rfidCode" column:"rfid_code"`
	ProfilePicture *string `json:"profilePicture" column:"profile_picture"`
	EmployeeID     *string `json:"employeeId" column:"employee_id"`
	ModifiedBy     int     `json:"modifiedById" column:"modified_by_id"`

	CreatedAt *time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" column:"updated_at"`
}

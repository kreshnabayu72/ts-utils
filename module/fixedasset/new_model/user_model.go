package new_model

import "time"

type User struct {
	ID                     int        `json:"_id"                    column:"id"`
	FirstName              string     `json:"firstName"              column:"first_name"`
	LastName               string     `json:"lastName"               column:"last_name"`
	FullName               string     `json:"fullName"               column:"full_name"`
	Email                  string     `json:"email"                  column:"email"`
	IsActive               bool       `json:"isActive"               column:"is_active"`
	IsDefault              bool       `json:"isDefault"              column:"is_default"`
	IsEmailConfirmed       bool       `json:"isEmailConfirmed"       column:"is_email_confirmed"`
	ExpiryDate             *time.Time `json:"expiryDate" column:"expiry_date"`
	IsExpired              bool       `json:"isExpired" column:"is_expired"`
	Position               *string    `json:"position"               column:"position"`
	Division               string     `json:"division"               column:"division"`
	RFIDCode               *string    `json:"rfidCode"               column:"rfid_code"`
	ProfilePicture         *string    `json:"profilePicture"         column:"profile_picture"`
	EmployeeId             *string    `json:"employeeId"             column:"employee_id"`
	IsSubUser              bool       `json:"isSubUser"              column:"is_sub_user"`
	ParentId               *int       `json:"parentId"               column:"parent_id"`
	ReadOnlyIsActive       bool       `json:"readOnlyIsActive"       column:"read_only_is_active"`
	ReadOnlyIsAssigned     bool       `json:"readOnlyIsAssigned"     column:"read_only_is_assigned"`
	TotalControlIsActive   bool       `json:"totalControlIsActive"   column:"total_control_is_active"`
	TotalControlIsAssigned bool       `json:"totalControlIsAssigned" column:"total_control_is_assigned"`
	ModifiedById           int        `json:"modifiedById"           column:"modified_by_id"`
	ModifiedByFirstName    string     `json:"modifiedByFirstName"    column:"modified_by_first_name"`
	ModifiedByLastName     string     `json:"modifiedByLastName"     column:"modified_by_last_name"`
	ModifiedByFullName     string     `json:"modifiedByFullName"     column:"modified_by_full_name"`
	IsDeleted              bool       `json:"isDeleted"              column:"is_deleted"`
	DeleteFromDatabaseAt   *time.Time `json:"deleteFromDatabaseAt"   column:"delete_from_database_at"`
	CreatedAt              time.Time  `json:"createdAt"              column:"created_at"`
	UpdatedAt              time.Time  `json:"updatedAt"              column:"updated_at"`
}

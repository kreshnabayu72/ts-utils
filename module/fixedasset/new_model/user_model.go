package new_model

import (
	"time"
)

type User struct {
	ID                     int        `json:"_id"                    column:"users.id"`
	FCMToken               string     `json:"fcmToken"               column:"users.fcm_token"`
	FirstName              string     `json:"firstName"              column:"users.first_name"`
	LastName               string     `json:"lastName"               column:"users.last_name"`
	FullName               string     `json:"fullName"               column:"users.full_name"                                                      special:"generated"`
	Email                  string     `json:"email"                  column:"users.email"`
	PhoneNumber            string     `json:"phoneNumber"            column:"users.phone_number"`
	PendingEmailChange     string     `json:"pendingEmailChange"     column:"users.pending_email_change"`
	IsActive               bool       `json:"isActive"               column:"users.is_active"`
	IsTemporary            bool       `json:"isTemporary"            column:"users.is_temporary"`
	IsDefault              bool       `json:"isDefault"              column:"users.is_default"`
	IsEmailConfirmed       bool       `json:"isEmailConfirmed"       column:"users.is_email_confirmed"`
	ExpiryDate             *time.Time `json:"expiryDate"             column:"users.expiry_date"`
	IsExpired              *bool      `json:"isExpired"              column:"CASE WHEN expiry_date IS NULL THEN NULL ELSE expiry_date < NOW() END" special:"generated"`
	Position               *string    `json:"position"               column:"users.position"`
	Division               string     `json:"division"               column:"users.division"`
	RfidCode               *string    `json:"rfidCode"               column:"users.rfid_code"`
	RfidId                 *int       `json:"rfidId"                 column:"users.rfid_id"`
	ProfilePicture         *string    `json:"profilePicture"         column:"users.profile_picture"`
	EmployeeId             *string    `json:"employeeId"             column:"users.employee_id"`
	IsSubUser              bool       `json:"isSubUser"              column:"users.is_sub_user"`
	ParentId               *int       `json:"parentId"               column:"users.parent_id"`
	ReadOnlyIsActive       bool       `json:"readOnlyIsActive"       column:"users.read_only_is_active"`
	ReadOnlyIsAssigned     bool       `json:"readOnlyIsAssigned"     column:"users.read_only_is_assigned"`
	TotalControlIsActive   bool       `json:"totalControlIsActive"   column:"users.total_control_is_active"`
	TotalControlIsAssigned bool       `json:"totalControlIsAssigned" column:"users.total_control_is_assigned"`
	ModifiedById           int        `json:"modifiedById"           column:"users.modified_by_id"`
	ModifiedByFirstName    string     `json:"modifiedByFirstName"    column:"users.modified_by_first_name"`
	ModifiedByLastName     string     `json:"modifiedByLastName"     column:"users.modified_by_last_name"`
	ModifiedByFullName     string     `json:"modifiedByFullName"     column:"users.modified_by_full_name"                                          special:"generated"`
	HasModuleAccess        bool       `json:"hasModuleAccess"        column:"users.has_module_access"` //For showing/hiding user that doesnt have fixed asset access from global settings
	IsDeleted              bool       `json:"isDeleted"              column:"users.is_deleted"`
	DeleteFromDatabaseAt   *time.Time `json:"deleteFromDatabaseAt"   column:"users.deleted_at"`
	CreatedAt              time.Time  `json:"createdAt"              column:"users.created_at"`
	UpdatedAt              time.Time  `json:"updatedAt"              column:"users.updated_at"`
}

type UserAttribute struct {
	ID        string `json:"_id,omitempty" column:"id::text"`
	FullName  string `json:"fullName"      column:"concat_ws(' ', first_name, last_name)"`
	IsDefault bool   `json:"isDefault"     column:"is_default"`
}

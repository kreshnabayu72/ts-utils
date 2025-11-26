package model

import (
	"time"

	fixed_asset_model "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/fixedasset/model"
)

type User struct {
	ID                 int                              `json:"_id"                column:"users.id"`
	FirstName          string                           `json:"firstName"          column:"users.first_name"`
	LastName           string                           `json:"lastName"           column:"users.last_name"`
	FullName           string                           `json:"fullName"           column:"users.full_name"             special:"generated"`
	PhoneNumber        string                           `json:"phoneNumber"        column:"users.phone_number"`
	Email              string                           `json:"email"              column:"users.email"`
	PendingEmailChange *string                          `json:"pendingEmailChange" column:"users.pending_email_change"`
	IsActive           bool                             `json:"isActive"           column:"users.is_active"`
	IsDefault          bool                             `json:"isDefault"          column:"users.is_default"`
	IsDeleted          bool                             `json:"isDeleted"          column:"users.is_deleted"`
	IsEmailConfirmed   bool                             `json:"isEmailConfirmed"   column:"users.is_email_confirmed"`
	PositionID         *int                             `json:"positionId"         column:"users.position_id"`
	DivisionID         int                              `json:"divisionId"         column:"users.division_id"`
	RfidCode           *string                          `json:"rfidCode"           column:"users.rfid_code"`
	ProfilePicture     *string                          `json:"profilePicture"     column:"users.profile_picture"`
	EmployeeID         *string                          `json:"employeeId"         column:"users.employee_id"`
	Access             []fixed_asset_model.ModuleAccess `json:"access"             column:"users.access"`
	ModifiedBy         int                              `json:"modifiedById"       column:"users.modified_by_id"`
	IsTemporary        bool                             `json:"isTemporary"        column:"users.is_temporary"`
	ExpiryDate         *time.Time                       `json:"expiryDate"         column:"users.expiry_date"`
	HasModuleAccess    bool                             `json:"hasModuleAccess"    column:"users.has_module_access"` //For showing/hiding user that doesn't have fixed asset access from global settings
	FCMToken           string                           `json:"fcmToken"           column:"users.fcm_token"`
	CreatedAt          *time.Time                       `json:"createdAt"          column:"users.created_at"`
	UpdatedAt          *time.Time                       `json:"updatedAt"          column:"users.updated_at"`
}

package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

// User is the tenant user row. Per-product state that used to live here as one column
// pair per product now lives elsewhere: activity in the active_per_product JSONB, and
// total-control / read-only in user_access_levels (see UserAccessLevel). Adding a
// product no longer means adding columns to this table.
type User struct {
	ID             int     `json:"_id"            column:"users.id"`
	FirstName      string  `json:"firstName"      column:"users.first_name"`
	LastName       string  `json:"lastName"       column:"users.last_name"`
	Email          string  `json:"email"          column:"users.email"`
	EmployeeId     *string `json:"employeeId"     column:"users.employee_id"`
	ProfilePicture *string `json:"profilePicture" column:"users.profile_picture"`
	TagId          *int    `json:"tagId"          column:"users.tag_id"`
	PositionId     *int    `json:"positionId"     column:"users.position_id"`
	DivisionId     int     `json:"divisionId"     column:"users.division_id"`

	ActivationStatus enum.UserActivationStatus `json:"activationStatus" column:"users.activation_status"`
	IsGlobalActive   bool                      `json:"isGlobalActive"   column:"users.is_global_active"`
	IsTemporary      bool                      `json:"isTemporary"      column:"users.is_temporary"`

	PreferredLanguage string `json:"preferredLanguage" column:"users.preferred_language"`
	DarkModeEnabled   bool   `json:"darkModeEnabled"   column:"users.dark_mode_enabled"`

	// ActivePerProduct replaces the is_active_<product> columns: product -> is active.
	// Absent key means the user has no access to that product at all.
	Access           []string                   `json:"access"           column:"users.access"`
	ActivePerProduct map[enum.ModuleAccess]bool `json:"activePerProduct" column:"users.active_per_product"`
	CustomFields     map[string]any             `json:"customFields"     column:"users.custom_fields"`

	ActivePeriodUntil *time.Time `json:"activePeriodUntil" column:"users.active_period_until"`

	DeletedAt  *time.Time `json:"deletedAt"  column:"users.deleted_at"`
	CreatedBy  *int       `json:"createdBy"  column:"users.created_by"`
	ModifiedBy *int       `json:"modifiedBy" column:"users.modified_by"`
	CreatedAt  time.Time  `json:"createdAt"  column:"users.created_at"`
	UpdatedAt  time.Time  `json:"updatedAt"  column:"users.updated_at"`
}

package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

// UserAccessLevel is a user's blanket access level for one product — the row form of
// users.is_total_control_* / is_read_only_*, which encoded the same fact as a pair of
// booleans per product and cost two new columns for every product added.
//
// Unique on (user_id, product): one level per product. The boolean pairs it replaces
// could express "total control AND read only", which was never a meaningful state.
//
// Product carries the same caveat as UserSystemRole.Product — it is typed
// enum.ModuleAccess ("Fixed Asset"), not analytics_reporting_service/model.Product
// ("fams"). Same concept, two value sets, still unconsolidated.
type UserAccessLevel struct {
	ID          int               `json:"_id"         column:"user_access_levels.id"`
	UserId      int               `json:"userId"      column:"user_access_levels.user_id"`
	Product     enum.ModuleAccess `json:"product"     column:"user_access_levels.product"`
	AccessLevel enum.AccessLevel  `json:"accessLevel" column:"user_access_levels.access_level"`
	ModifiedBy  *int              `json:"modifiedBy"  column:"user_access_levels.modified_by"`

	CreatedAt time.Time `json:"createdAt" column:"user_access_levels.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"user_access_levels.updated_at"`
}

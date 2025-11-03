package new_model

import (
	"time"
)

type User struct {
	ID       int    `json:"id"                                 column:"users.id"`
	Name     string `json:"name"                               column:"users.name"`
	FullName string `json:"fullName"                               column:"users.full_name"`

	PositionId int    `json:"positionId"                               column:"users.position_id"`
	Position   string `json:"position"                               column:"users.position"`
	DivisionId int    `json:"divisionId"                               column:"users.division_id"`
	Division   string `json:"division"                               column:"users.division"`

	IsActive  bool `json:"isActive"                          column:"users.is_active"`
	IsDeleted bool `json:"isDeleted"                          column:"users.is_deleted"`

	CreatedAt time.Time `json:"createdAt"                          column:"users.created_at"`
	UpdatedAt time.Time `json:"updatedAt"                          column:"users.updated_at"`
}

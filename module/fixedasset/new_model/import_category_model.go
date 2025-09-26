package new_model

import (
	"time"
)

type ImportCategory struct {
	ID           int       `json:"_id"          column:"import_categories.id"`
	Key          int       `json:"key"          column:"import_categories.key"           special:"generated"`
	Name         string    `json:"name"         column:"import_categories.name"`
	CategoryCode string    `json:"categoryCode" column:"import_categories.category_code"`
	CreatedAt    time.Time `json:"createdAt"    column:"import_categories.created_at"`
	UpdatedAt    time.Time `json:"updatedAt"    column:"import_categories.updated_at"`
}

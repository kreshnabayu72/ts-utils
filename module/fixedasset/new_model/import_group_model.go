package new_model

import (
	"time"
)

type ImportGroup struct {
	ID           int       `json:"_id"          column:"import_groups.id"`
	Key          int       `json:"key"          column:"import_groups.key"           special:"generated"`
	Name         string    `json:"name"         column:"import_groups.name"`
	GroupCode    string    `json:"groupCode"    column:"import_groups.group_code"`
	IsDisposable *string   `json:"isDisposable" column:"import_groups.is_disposable"`
	CreatedAt    time.Time `json:"createdAt"    column:"import_groups.created_at"`
	UpdatedAt    time.Time `json:"updatedAt"    column:"import_groups.updated_at"`
}

package replica

import "time"

type Asset struct {
	ID               int       `json:"_id"             column:"assets.id"                 sqlType:"BIGINT"`
	NameWithSequence string    `json:"nameWithSequence" column:"assets.name_with_sequence" sqlType:"VARCHAR"`
	Status           string    `json:"status"           column:"assets.status"             sqlType:"VARCHAR"`
	IsDeleted        bool      `json:"isDeleted"        column:"assets.is_deleted"         sqlType:"BOOLEAN"`
	DeletedAt        time.Time `json:"deletedAt"        column:"assets.deleted_at"         sqlType:"TIMESTAMPTZ"`
	CreatedAt        time.Time `json:"createdAt"        column:"assets.created_at"         sqlType:"TIMESTAMPTZ"`
	UpdatedAt        time.Time `json:"updatedAt"        column:"assets.updated_at"         sqlType:"TIMESTAMPTZ"`
}

func (a *Asset) GetTableName() string {
	return "assets"
}

func (a *Asset) GetTableNameWithSchema() string {
	return "public.assets"
}

func (a *Asset) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                 "BIGINT",
		"name_with_sequence": "VARCHAR",
		"status":             "VARCHAR",
		"is_deleted":         "BOOLEAN",
		"deleted_at":         "TIMESTAMPTZ",
		"created_at":         "TIMESTAMPTZ",
		"updated_at":         "TIMESTAMPTZ",
	}
}

func (a *Asset) GetColumns() []string {
	return []string{
		"id",
		"name_with_sequence",
		"status",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

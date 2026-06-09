package replica

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type Asset struct {
	ID               int              `json:"_id"             column:"assets.id"`
	NameWithSequence string           `json:"nameWithSequence" column:"assets.name_with_sequence"`
	Status           enum.AssetStatus `json:"status"           column:"assets.status"`
	IsDeleted        bool             `json:"isDeleted"        column:"assets.is_deleted"`
	DeletedAt        *time.Time       `json:"deletedAt"        column:"assets.deleted_at"`
	CreatedAt        time.Time        `json:"createdAt"        column:"assets.created_at"`
	UpdatedAt        time.Time        `json:"updatedAt"        column:"assets.updated_at"`
}

func (a *Asset) GetTableName() string {
	return "assets"
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

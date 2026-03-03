package replica

import "time"

type FixedAssetGroup struct {
	ID        int        `json:"_id" column:"fixed_asset_groups.id"`
	Name      string     `json:"name" column:"fixed_asset_groups.name"`
	IsDeleted bool       `json:"isDeleted" column:"fixed_asset_groups.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"fixed_asset_groups.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"fixed_asset_groups.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"fixed_asset_groups.updated_at"`
}

func (r *FixedAssetGroup) GetTableName() string {
	return "fixed_asset_groups"
}

func (r *FixedAssetGroup) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"is_deleted": "BOOL",
		"deleted_at": "TIMESTAMPTZ",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
}

func (r *FixedAssetGroup) GetColumns() []string {
	return []string{
		"id",
		"name",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

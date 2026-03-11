package replica

import "time"

type SupplyAssetGroup struct {
	ID        int        `json:"_id" column:"supply_asset_groups.id"`
	Name      string     `json:"name" column:"supply_asset_groups.name"`
	IsDeleted bool       `json:"isDeleted" column:"supply_asset_groups.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"supply_asset_groups.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"supply_asset_groups.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"supply_asset_groups.updated_at"`
}

func (r *SupplyAssetGroup) GetTableName() string {
	return "supply_asset_groups"
}

func (r *SupplyAssetGroup) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"is_deleted": "BOOL",
		"deleted_at": "TIMESTAMPTZ",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
}

func (r *SupplyAssetGroup) GetColumns() []string {
	return []string{
		"id",
		"name",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

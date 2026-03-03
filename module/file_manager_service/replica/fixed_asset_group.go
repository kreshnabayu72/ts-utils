package replica

import "time"

type FixedAssetGroup struct {
	ID        int        `json:"_id"        column:"fixed_asset_groups.id"`
	Name      string     `json:"name"       column:"fixed_asset_groups.name"`
	IsDeleted bool       `json:"isDeleted"  column:"fixed_asset_groups.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"  column:"fixed_asset_groups.deleted_at"`
}

func (f *FixedAssetGroup) GetTableName() string {
	return "fixed_asset_groups"
}

func (f *FixedAssetGroup) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"is_deleted": "BOOLEAN",
		"deleted_at": "TIMESTAMPTZ",
	}
}

func (f *FixedAssetGroup) GetColumns() []string {
	return []string{
		"id",
		"name",
		"is_deleted",
		"deleted_at",
	}
}

package replica

import "time"

type FixedAssetGroup struct {
	ID        int       `json:"_id"        column:"fixed_asset_groups.id"         sqlType:"BIGINT"`
	Name      string    `json:"name"       column:"fixed_asset_groups.name"       sqlType:"VARCHAR"`
	Code      string    `json:"code"       column:"fixed_asset_groups.code"       sqlType:"VARCHAR"`
	FullPath  string    `json:"fullPath"   column:"fixed_asset_groups.full_path"  sqlType:"VARCHAR"`
	IsDeleted bool      `json:"isDeleted"  column:"fixed_asset_groups.is_deleted" sqlType:"BOOLEAN"`
	DeletedAt time.Time `json:"deletedAt"  column:"fixed_asset_groups.deleted_at" sqlType:"TIMESTAMPTZ"`
}

func (f *FixedAssetGroup) GetTableName() string {
	return "fixed_asset_groups"
}

func (f *FixedAssetGroup) GetTableNameWithSchema() string {
	return "public.fixed_asset_groups"
}

func (f *FixedAssetGroup) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"code":       "VARCHAR",
		"full_path":  "VARCHAR",
		"is_deleted": "BOOLEAN",
		"deleted_at": "TIMESTAMPTZ",
	}
}

func (f *FixedAssetGroup) GetColumns() []string {
	return []string{
		"id",
		"name",
		"code",
		"full_path",
		"is_deleted",
		"deleted_at",
	}
}

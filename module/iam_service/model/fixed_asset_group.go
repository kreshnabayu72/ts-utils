package model

type FixedAssetGroup struct {
	ID       int    `json:"_id"      column:"fixed_asset_groups.id"`
	Name     string `json:"name"     column:"fixed_asset_groups.name"`
	Code     string `json:"code"     column:"fixed_asset_groups.code"`
	Level    int    `json:"level"    column:"fixed_asset_groups.level"`
	ParentId *int   `json:"parentId" column:"fixed_asset_groups.parent_id"`
}

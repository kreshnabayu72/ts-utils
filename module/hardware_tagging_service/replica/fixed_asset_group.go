package replica

import "time"

type FixedAssetGroup struct {
	ID                  int        `json:"_id" column:"fixed_asset_groups.id"`
	Name                string     `json:"name" column:"fixed_asset_groups.name"`
	Code                string     `json:"code" column:"fixed_asset_groups.code"`
	Level               int        `json:"level" column:"fixed_asset_groups.level"`
	ParentId            *int       `json:"parentId" column:"fixed_asset_groups.parent_id"`
	AncestorId          *int       `json:"ancestorId" column:"fixed_asset_groups.ancestor_id"`
	ParentQuotaSourceId int        `json:"parentQuotaSourceId" column:"fixed_asset_groups.parent_quota_source_id"`
	FullPath            string     `json:"fullPath" column:"fixed_asset_groups.full_path"`
	IsDeleted           bool       `json:"isDeleted" column:"fixed_asset_groups.is_deleted"`
	DeletedAt           *time.Time `json:"deletedAt" column:"fixed_asset_groups.deleted_at"`
}

func (r *FixedAssetGroup) GetTableName() string {
	return "fixed_asset_groups"
}

func (r *FixedAssetGroup) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":                     "BIGINT",
		"name":                   "VARCHAR",
		"code":                   "VARCHAR",
		"level":                  "INT",
		"is_disposable":          "BOOL",
		"parent_id":              "BIGINT",
		"ancestor_id":            "BIGINT",
		"parent_quota_source_id": "BIGINT",
		"full_path":              "VARCHAR",
		"is_deleted":             "BOOL",
		"deleted_at":             "TIMESTAMPTZ",
		"created_at":             "TIMESTAMPTZ",
		"updated_at":             "TIMESTAMPTZ",
	}
}

func (r *FixedAssetGroup) GetColumns() []string {
	return []string{
		"id",
		"name",
		"code",
		"level",
		"is_disposable",
		"parent_id",
		"ancestor_id",
		"parent_quota_source_id",
		"full_path",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

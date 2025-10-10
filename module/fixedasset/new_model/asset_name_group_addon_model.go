package new_model

import "time"

type AssetNameGroupAddOn struct {
	ID        int              `json:"_id"       column:"asset_name_group_addons.id"`
	NameID    int              `json:"nameId"    column:"asset_name_group_addons.name_id"`
	GroupID   int              `json:"groupId"   column:"asset_name_group_addons.group_id"`
	Type      AssetLicenseType `json:"type"      column:"asset_name_group_addons.type"`
	CreatedAt time.Time        `json:"createdAt" column:"asset_name_group_addons.created_at"`
	UpdatedAt time.Time        `json:"updatedAt" column:"asset_name_group_addons.updated_at"`
}

package model

import "time"

type SupplyAssetGroup struct {
	ID        int        `json:"_id" column:"supply_asset_groups.id"`
	Name      string     `json:"name" column:"supply_asset_groups.name"`
	IsDeleted bool       `json:"isDeleted" column:"supply_asset_groups.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"supply_asset_groups.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"supply_asset_groups.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"supply_asset_groups.updated_at"`
}


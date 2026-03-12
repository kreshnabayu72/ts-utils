package model

import "time"

type HandheldReader struct {
	ID                 int        `json:"id" column:"handheld_readers.id"`
	AliasName          string     `json:"aliasName" column:"handheld_readers.alias_name"`
	Status             string     `json:"status" column:"handheld_readers.status"`
	HandheldStockId    int        `json:"handheldStockId" column:"handheld_readers.handheld_stock_id"`
	FixedAssetGroupId  *int       `json:"fixedAssetGroupId" column:"handheld_readers.fixed_asset_group_id"`
	SupplyAssetGroupId *int       `json:"supplyAssetGroupId" column:"handheld_readers.supply_asset_group_id"`
	ReportedById       *int       `json:"reportedById" column:"handheld_readers.reported_by_id"`
	ReportedAt         *time.Time `json:"reportedAt" column:"handheld_readers.reported_at"`
	IsDeleted          bool       `json:"isDeleted" column:"handheld_readers.is_deleted"`
	DeletedAt          *time.Time `json:"deletedAt" column:"handheld_readers.deleted_at"`
	CreatedAt          *time.Time `json:"createdAt" column:"handheld_readers.created_at"`
	UpdatedAt          *time.Time `json:"updatedAt" column:"handheld_readers.updated_at"`
}
func (r *HandheldReader) GetColumns() []string {
	return []string{
		"id",
		"alias_name",
		"status",
		"handheld_stock_id",
		"fixed_asset_group_id",
		"supply_asset_group_id",
		"reported_by_id",
		"reported_at",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

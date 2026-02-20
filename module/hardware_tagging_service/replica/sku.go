package replica

import "time"

type Sku struct {
	ID        int        `json:"_id"       column:"skus.id" sqlType:"BIGINT"`
	Name      string     `json:"shortCode" column:"skus.name" sqlType:"VARCHAR"`
	Code      string     `json:"code"      column:"skus.code" sqlType:"VARCHAR"`
	Image     *string    `json:"image"     column:"skus.image" sqlType:"VARCHAR"`
	IsDeleted bool       `json:"isDeleted" column:"skus.is_deleted" sqlType:"BOOL"`
	DeletedAt *time.Time `json:"deletedAt" column:"skus.deleted_at" sqlType:"TIMESTAMPTZ"`
	CreatedAt time.Time  `json:"createdAt" column:"skus.created_at" sqlType:"TIMESTAMPTZ"`
	UpdatedAt time.Time  `json:"updatedAt" column:"skus.updated_at" sqlType:"TIMESTAMPTZ"`
}

func (r *Sku) GetTableName() string {
	return "skus"
}

func (r *Sku) GetTableNameWithSchema() string {
	return "public.skus"
}

func (r *Sku) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"code":       "VARCHAR",
		"image":      "VARCHAR",
		"is_deleted": "BOOL",
		"deleted_at": "TIMESTAMPTZ",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
}

func (r *Sku) GetColumns() []string {
	return []string{
		"id",
		"name",
		"code",
		"image",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

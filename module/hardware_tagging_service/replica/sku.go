package replica

import "time"

type Sku struct {
	ID           int        `json:"_id"       column:"skus.id"`
	DeviceNameID int        `json:"deviceNameId" column:"skus.device_name_id"`
	Name         string     `json:"shortCode" column:"skus.name"`
	Code         string     `json:"code"      column:"skus.code"`
	FirstImage   *string    `json:"firstImage"     column:"skus.first_image"`
	SecondImage  *string    `json:"secondImage"     column:"skus.second_image"`
	SetDefault   *string    `json:"setDefault"     column:"skus.set_default"`
	Image        *string    `json:"image"     column:"skus.image" special:"generated"`
	IsDeleted    bool       `json:"isDeleted" column:"skus.is_deleted"`
	TotalPorts   int        `json:"totalPorts" column:"skus.total_ports"`
	DeletedAt    *time.Time `json:"deletedAt" column:"skus.deleted_at"`
	CreatedAt    time.Time  `json:"createdAt" column:"skus.created_at"`
	UpdatedAt    time.Time  `json:"updatedAt" column:"skus.updated_at"`
}

func (r *Sku) GetTableName() string {
	return "skus"
}

func (r *Sku) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":             "BIGINT",
		"device_name_id": "BIGINT",
		"name":           "VARCHAR",
		"code":           "VARCHAR",
		"first_image":    "VARCHAR",
		"second_image":   "VARCHAR",
		"set_default":    "VARCHAR",
		"image":          "VARCHAR",
		"is_deleted":     "BOOL",
		"total_ports":    "INTEGER",
		"deleted_at":     "TIMESTAMPTZ",
		"created_at":     "TIMESTAMPTZ",
		"updated_at":     "TIMESTAMPTZ",
	}
}

func (r *Sku) GetColumns() []string {
	return []string{
		"id",
		"device_name_id",
		"name",
		"code",
		"first_image",
		"second_image",
		"set_default",
		"image",
		"is_deleted",
		"total_ports",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

package model

import "time"

type Sku struct {
	ID           int `json:"id" column:"skus.id"`
	DeviceNameID int `json:"deviceNameId" column:"skus.device_name_id"`

	// Either only name or only code or both (depends on final admin figma)
	Name string `json:"name" column:"skus.name"`
	Code string `json:"code" column:"skus.code"`

	FirstImage  string     `json:"firstImage" column:"skus.first_image"`
	SecondImage string     `json:"secondImage" column:"skus.second_image"`
	SetDefault  string     `json:"setDefault" column:"skus.set_default"`
	Image       string     `json:"image" column:"skus.image" special:"generated"`
	TotalPorts  int        `json:"totalPorts" column:"skus.total_ports"`
	IsDeleted   bool       `json:"isDeleted" column:"skus.is_deleted"`
	DeletedAt   *time.Time `json:"deletedAt" column:"skus.deleted_at"`
	CreatedAt   time.Time  `json:"createdAt" column:"skus.created_at"`
	UpdatedAt   time.Time  `json:"updatedAt" column:"skus.updated_at"`
}


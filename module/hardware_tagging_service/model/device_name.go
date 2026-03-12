package model

import "time"

type DeviceName struct {
	ID        int        `json:"id" column:"device_names.id"`
	Name      string     `json:"name" column:"device_names.name"`
	IsDeleted bool       `json:"isDeleted" column:"device_names.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"device_names.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"device_names.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"device_names.updated_at"`
}
func (r *DeviceName) GetColumns() []string {
	return []string{
		"id",
		"name",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

package replica

import "time"

type DeviceName struct {
	ID        int        `json:"_id" column:"device_names.id"`
	Name      string     `json:"name" column:"device_names.name"`
	IsDeleted bool       `json:"isDeleted" column:"device_names.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"device_names.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"device_names.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"device_names.updated_at"`
}

func (r *DeviceName) GetTableName() string {
	return "device_names"
}

func (r *DeviceName) GetTableNameWithSchema() string {
	return "public.device_names"
}

func (r *DeviceName) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":         "BIGINT",
		"name":       "VARCHAR",
		"is_deleted": "BOOL",
		"deleted_at": "TIMESTAMPTZ",
		"created_at": "TIMESTAMPTZ",
		"updated_at": "TIMESTAMPTZ",
	}
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

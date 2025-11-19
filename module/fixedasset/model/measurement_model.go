package model

import (
	"time"
)

type Measurement struct {
	ID        int        `json:"_id"       column:"measurements.id"`
	Name      string     `json:"name"      column:"measurements.name"`
	IsDeleted bool       `json:"isDeleted" column:"measurements.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"measurements.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"measurements.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"measurements.updated_at"`
}

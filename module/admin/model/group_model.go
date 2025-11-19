package model

import (
	"time"
)

type Group struct {
	ID        int        `json:"id"        column:"groups.id"`
	Name      string     `json:"name"      column:"groups.name"`
	Key       int        `json:"key"       column:"groups.key"`
	Level     int        `json:"level"     column:"groups.level"`
	ParentID  int        `json:"parentId"  column:"groups.parent_id"`
	IsDeleted bool       `json:"isDeleted" column:"groups.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt" column:"groups.deleted_at"`
	CreatedAt time.Time  `json:"createdAt" column:"groups.created_at"`
	UpdatedAt time.Time  `json:"updatedAt" column:"groups.updated_at"`
}

package model

import "time"

type Position struct {
	ID        int       `json:"_id"       column:"positions.id"`
	Name      string    `json:"name"      column:"positions.name"`
	CreatedAt time.Time `json:"createdAt" column:"positions.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"positions.updated_at"`
}


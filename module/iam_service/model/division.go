package model

import "time"

type Division struct {
	ID        int       `json:"_id"       column:"divisions.id"`
	Name      string    `json:"name"      column:"divisions.name"`
	CreatedAt time.Time `json:"createdAt" column:"divisions.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"divisions.updated_at"`
}
func (r *Division) GetColumns() []string {
	return []string{
		"id",
		"name",
		"created_at",
		"updated_at",
	}
}

package model

import (
	"time"
)

type Qr struct {
	ID           int        `json:"_id"          column:"qrs.id"`
	Code         string     `json:"code"         column:"qrs.code"`
	ShortCode    string     `json:"shortCode"    column:"qrs.short_code"`
	Status       TagStatus  `json:"status"       column:"qrs.status"`
	ReportedById *int       `json:"reportedById" column:"qrs.reported_by_id"`
	ReportedAt   *time.Time `json:"reportedAt"   column:"qrs.reported_at"`
	CreatedAt    time.Time  `json:"createdAt"    column:"qrs.created_at"`
	UpdatedAt    time.Time  `json:"updatedAt"    column:"qrs.updated_at"`
}

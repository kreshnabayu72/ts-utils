package model

import (
	"time"
)

type AuditScheduledAsset struct {
	ID              int       `json:"_id"             column:"audit_scheduled_assets.id"`
	Asset           int       `json:"asset"           column:"audit_scheduled_assets.asset"`
	AuditScheduleId int       `json:"auditScheduleId" column:"audit_scheduled_assets.audit_schedule_id"`
	IsActive        bool      `json:"isActive"        column:"audit_scheduled_assets.is_active"`
	CreatedAt       time.Time `json:"createdAt"       column:"audit_scheduled_assets.created_at"`
	UpdatedAt       time.Time `json:"updatedAt"       column:"audit_scheduled_assets.updated_at"`
}

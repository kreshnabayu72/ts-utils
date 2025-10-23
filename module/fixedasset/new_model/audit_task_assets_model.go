package new_model

import (
	"time"
)

type AuditTaskAssets struct {
	ID           int                      `json:"_id"          column:"audit_task_assets.id"`
	Asset        int                      `json:"asset"        column:"audit_task_assets.asset"`
	AuditTaskId  int                      `json:"auditTaskId"  column:"audit_task_assets.audit_task_id"`
	FoundById    *int                     `json:"foundById"    column:"audit_task_assets.found_by_id"`
	Status       AuditTaskAssetStatus     `json:"status"       column:"audit_task_assets.status"`
	Condition    *AuditTaskAssetCondition `json:"condition"    column:"audit_task_assets.condition"`
	Notes        string                   `json:"notes"        column:"audit_task_assets.notes"`
	DamageImage  *string                  `json:"damageImage"        column:"audit_task_assets.damage_image"`
	FoundGroupId *int                     `json:"foundGroupId" column:"audit_task_assets.found_group_id"`
	Address      string                   `json:"address"      column:"audit_task_assets.address"`
	Coordinate   string                   `json:"coordinate"   column:"audit_task_assets.coordinate"`
	FoundAt      *time.Time               `json:"foundAt"    column:"audit_task_assets.found_at"`
	AuditedAt    *time.Time               `json:"auditedAt"    column:"audit_task_assets.audited_at"`
	CreatedAt    time.Time                `json:"createdAt"    column:"audit_task_assets.created_at"`
	UpdatedAt    time.Time                `json:"updatedAt"    column:"audit_task_assets.updated_at"`
}

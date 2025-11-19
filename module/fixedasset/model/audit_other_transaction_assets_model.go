package model

import (
	"time"
)

type AuditOtherTransactionAsset struct {
	ID          int                   `json:"_id"         column:"audit_other_transaction_assets.id"`
	Asset       int                   `json:"asset"       column:"audit_other_transaction_assets.asset"`
	AuditTaskId int                   `json:"auditTaskId" column:"audit_other_transaction_assets.audit_task_id"`
	Type        AuditOtherTransaction `json:"type"        column:"audit_other_transaction_assets.type"`
	CreatedAt   time.Time             `json:"createdAt"   column:"audit_other_transaction_assets.created_at"`
	UpdatedAt   time.Time             `json:"updatedAt"   column:"audit_other_transaction_assets.updated_at"`
}

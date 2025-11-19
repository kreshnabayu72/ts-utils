package model

import (
	"time"
)

type AuditTaskCategories struct {
	ID          int       `json:"_id"         column:"audit_task_categories.id"`
	CategoryID  int       `json:"categoryId"  column:"audit_task_categories.category_id"`
	AuditTaskId int       `json:"auditTaskId" column:"audit_task_categories.audit_task_id"`
	CreatedAt   time.Time `json:"createdAt"   column:"audit_task_categories.created_at"`
	UpdatedAt   time.Time `json:"updatedAt"   column:"audit_task_categories.updated_at"`
}

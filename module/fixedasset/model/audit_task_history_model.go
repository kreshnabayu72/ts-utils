package model

import (
	"time"
)

type AuditTaskHistory struct {
	ID            int        `json:"_id"            column:"audit_task_histories.id"`
	TransactionId string     `json:"transaction_id" column:"audit_task_histories.transaction_id"`
	Name          string     `json:"name"           column:"audit_task_histories.name"`
	GroupID       int        `json:"groupId"        column:"audit_task_histories.group_id"`
	CategoryNames []string   `json:"categoryNames"  column:"audit_task_histories.category_names"`
	GroupName     string     `json:"groupName"      column:"audit_task_histories.group_name"`
	TagType       string     `json:"tagType"        column:"audit_task_histories.tag_type"`
	AuditType     string     `json:"auditType"      column:"audit_task_histories.audit_type"`
	Status        string     `json:"status"         column:"audit_task_histories.status"`
	Duration      int        `json:"duration"       column:"audit_task_histories.duration"`
	ModifiedById  *int       `json:"modifiedById"   column:"audit_task_histories.modified_by_id"`
	ApprovalRound int        `json:"approvalRound"  column:"audit_task_histories.approval_round"`
	GracePeriodAt *time.Time `json:"gracePeriodAt"  column:"audit_task_histories.grace_period_at"`
	StartedAt     *time.Time `json:"startedAt"      column:"audit_task_histories.started_at"`
	FinishedAt    *time.Time `json:"finishedAt"     column:"audit_task_histories.finished_at"`
	CreatedAt     time.Time  `json:"createdAt"      column:"audit_task_histories.created_at"`
	UpdatedAt     time.Time  `json:"updatedAt"      column:"audit_task_histories.updated_at"`
}

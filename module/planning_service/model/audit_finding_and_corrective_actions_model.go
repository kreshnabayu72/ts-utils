package new_model

import (
	"time"
)

type AuditFindingAndCorrectiveActions struct {
	ID int `json:"id" column:"audit_finding_and_corrective_actions.id"`

	FindingId        string `json:"findingId" column:"audit_finding_and_corrective_actions.finding_id"`
	AuditFinding     string `json:"auditFinding" column:"audit_finding_and_corrective_actions.audit_finding"`
	CorrectiveAction string `json:"correctiveAction" column:"audit_finding_and_corrective_actions.corrective_action"`

	PicId       int    `json:"picId" column:"audit_finding_and_corrective_actions.pic_id"`
	PicFullName string `json:"picFullName" column:"audit_finding_and_corrective_actions.pic_full_name"`

	TargetDate time.Time `json:"targetDate" column:"audit_finding_and_corrective_actions.target_date"`
	Status     string    `json:"status" column:"audit_finding_and_corrective_actions.status"`

	IsDeleted bool       `json:"isDeleted"        column:"audit_finding_and_corrective_actions.is_deleted"`
	DeletedAt *time.Time `json:"deletedAt"        column:"audit_finding_and_corrective_actions.deleted_at"`

	CreatedAt time.Time `json:"createdAt" column:"audit_finding_and_corrective_actions.created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"audit_finding_and_corrective_actions.updated_at"`
}

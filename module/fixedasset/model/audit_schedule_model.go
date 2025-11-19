package model

import (
	"time"
)

type AuditSchedule struct {
	ID                   int        `json:"_id"                  column:"audit_schedules.id"`
	IsActive             bool       `json:"isActive"             column:"audit_schedules.is_active"`
	Name                 string     `json:"name"                 column:"audit_schedules.name"`
	GroupID              int        `json:"groupId"              column:"audit_schedules.group_id"`
	TagType              TagType    `json:"tagType"              column:"audit_schedules.tag_type"`
	AuditType            AuditType  `json:"auditType"            column:"audit_schedules.audit_type"`
	FrequencyLabel       Frequency  `json:"frequencyLabel"       column:"audit_schedules.frequency_label"`
	FrequencyCustomValue *int       `json:"frequencyCustomValue" column:"audit_schedules.frequency_custom_value"`
	FrequencyCustomLabel *string    `json:"frequencyCustomLabel" column:"audit_schedules.frequency_custom_label"`
	Duration             int        `json:"duration"             column:"audit_schedules.duration"`
	CreatedById          int        `json:"createdById"         column:"audit_schedules.created_by_id"`
	ModifiedById         int        `json:"modifiedById"         column:"audit_schedules.modified_by_id"`
	IsDeleted            bool       `json:"isDeleted"            column:"audit_schedules.is_deleted"`
	NextSchedule         *time.Time `json:"nextSchedule"         column:"audit_schedules.next_schedule"`
	DeletedAt            *time.Time `json:"deletedAt"            column:"audit_schedules.deleted_at"`
	CreatedAt            time.Time  `json:"createdAt"            column:"audit_schedules.created_at"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"audit_schedules.updated_at"`
}

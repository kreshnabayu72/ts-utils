package new_model

import (
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type AuditSchedule struct {
	ID                   int        `json:"_id"                  column:"audit_schedules.id"                     sql:"BIGINT PRIMARY KEY"`
	IsActive             bool       `json:"isActive"             column:"audit_schedules.is_active"              sql:"BOOLEAN DEFAULT TRUE"`
	Name                 string     `json:"name"                 column:"audit_schedules.name"                   sql:"VARCHAR"`
	GroupID              int        `json:"groupId"              column:"audit_schedules.group_id"               sql:"BIGINT"`
	TagType              TagType    `json:"tagType"              column:"audit_schedules.tag_type"               sql:"tag_type_enum"`
	AuditType            AuditType  `json:"auditType"            column:"audit_schedules.audit_type"             sql:"audit_type_enum"`
	FrequencyLabel       Frequency  `json:"frequencyLabel"       column:"audit_schedules.frequency_label"        sql:"VARCHAR"`
	FrequencyCustomValue *int       `json:"frequencyCustomValue" column:"audit_schedules.frequency_custom_value" sql:"INTEGER NULL DEFAULT NULL"`
	Duration             int        `json:"duration"             column:"audit_schedules.duration"               sql:"INTEGER NOT NULL"`
	ModifiedById         int        `json:"modifiedById"         column:"audit_schedules.modified_by_id"         sql:"BIGINT NOT NULL"`
	IsDeleted            bool       `json:"isDeleted"            column:"audit_schedules.is_deleted"             sql:"BOOLEAN NOT NULL DEFAULT FALSE"`
	NextSchedule         *time.Time `json:"nextSchedule"         column:"audit_schedules.next_schedule"          sql:"TIMESTAMPTZ NULL"`
	DeletedAt            *time.Time `json:"deletedAt"            column:"audit_schedules.deleted_at"             sql:"TIMESTAMPTZ NULL"`
	CreatedAt            time.Time  `json:"createdAt"            column:"audit_schedules.created_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt            time.Time  `json:"updatedAt"            column:"audit_schedules.updated_at"             sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (b *AuditSchedule) TableName() string {
	return "audit_schedules"
}

func (b *AuditSchedule) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(b)

	builder.AddForeignKey(
		"fk_audit_schedules_modified_by",
		"modified_by_id",
		"users",
		"id",
		"SET NULL",
	)

	builder.AddForeignKey(
		"fk_audit_schedules_group",
		"group_id",
		"groups",
		"id",
		"",
	)
}

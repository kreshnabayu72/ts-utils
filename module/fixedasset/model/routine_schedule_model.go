package model

import (
	"time"
)

type RoutineSchedule struct {
	ID                   int                      `json:"_id"                  column:"routine_schedules.id"`
	Name                 string                   `json:"name"                 column:"routine_schedules.name"`
	AssetNameID          int                      `json:"assetNameId"          column:"routine_schedules.asset_name_id"`
	GroupID              int                      `json:"groupId"              column:"routine_schedules.group_id"`
	FrequencyLabel       RoutineScheduleFrequency `json:"frequencyLabel"       column:"routine_schedules.frequency_label"`
	FrequencyCustomValue *int                     `json:"frequencyCustomValue" column:"routine_schedules.frequency_custom_value"`
	FrequencyCustomLabel *string                  `json:"frequencyCustomLabel" column:"routine_schedules.frequency_custom_label"`
	Duration             int                      `json:"duration"             column:"routine_schedules.duration"`
	ModifiedByID         int                      `json:"modifiedById"         column:"routine_schedules.modified_by_id"`
	CreatedByID          int                      `json:"createdById"          column:"routine_schedules.created_by_id"`
	IsDeleted            bool                     `json:"isDeleted"            column:"routine_schedules.is_deleted"`
	DeletedAt            *time.Time               `json:"deletedAt"            column:"routine_schedules.deleted_at"`
	NextSchedule         time.Time                `json:"nextSchedule"         column:"routine_schedules.next_schedule"`
	CreatedAt            time.Time                `json:"createdAt"            column:"routine_schedules.created_at"`
	UpdatedAt            time.Time                `json:"updatedAt"            column:"routine_schedules.updated_at"`
}

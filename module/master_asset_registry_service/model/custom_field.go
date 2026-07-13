package model

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
)

type CustomField struct {
	ID          int                      `json:"_id"         column:"custom_fields.id"`
	Name        string                   `json:"name"        column:"custom_fields.name"`
	Type        enum.CustomFieldType     `json:"type"        column:"custom_fields.type"`
	DataType    enum.CustomFieldDataType `json:"dataType"    column:"custom_fields.data_type"`
	IsRequired  bool                     `json:"isRequired"  column:"custom_fields.is_required"`
	OptionValue []string                 `json:"optionValue" column:"custom_fields.option_value"`
	IsActive    bool                     `json:"isActive"    column:"custom_fields.is_active"`
	IsDeleted   bool                     `json:"isDeleted"   column:"custom_fields.is_deleted"`
	DeletedAt   *time.Time               `json:"deletedAt"   column:"custom_fields.deleted_at"`
	CreatedAt   time.Time                `json:"createdAt"   column:"custom_fields.created_at"`
	UpdatedAt   time.Time                `json:"updatedAt"   column:"custom_fields.updated_at"`
}


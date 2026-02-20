package replica

import (
	"time"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/module/fixedasset/model"
)

type CustomField struct {
	ID          int                       `json:"_id"         column:"custom_fields.id"`
	Name        string                    `json:"name"        column:"custom_fields.name"`
	Type        model.CustomFieldType     `json:"type"        column:"custom_fields.type"`
	DataType    model.CustomFieldDataType `json:"dataType"    column:"custom_fields.data_type"`
	IsRequired  bool                      `json:"isRequired"  column:"custom_fields.is_required"`
	OptionValue []string                  `json:"optionValue" column:"custom_fields.option_value"`
	IsActive    bool                      `json:"isActive"    column:"custom_fields.is_active"`
	IsDeleted   bool                      `json:"isDeleted"   column:"custom_fields.is_deleted"`
	DeletedAt   *time.Time                `json:"deletedAt"   column:"custom_fields.deleted_at"`
	CreatedAt   time.Time                 `json:"createdAt"   column:"custom_fields.created_at"`
	UpdatedAt   time.Time                 `json:"updatedAt"   column:"custom_fields.updated_at"`
}

func (c *CustomField) GetTableName() string {
	return "custom_fields"
}

func (c *CustomField) GetTableNameWithSchema() string {
	return "public.custom_fields"
}

func (c *CustomField) GetColumnTypes() map[string]string {
	return map[string]string{
		"id":           "BIGINT",
		"name":         "VARCHAR",
		"type":         "VARCHAR",
		"data_type":    "VARCHAR",
		"is_required":  "BOOLEAN",
		"option_value": "TEXT[]",
		"is_active":    "BOOLEAN",
		"is_deleted":   "BOOLEAN",
		"deleted_at":   "TIMESTAMPTZ",
		"created_at":   "TIMESTAMPTZ",
		"updated_at":   "TIMESTAMPTZ",
	}
}

func (c *CustomField) GetColumns() []string {
	return []string{
		"id",
		"name",
		"type",
		"data_type",
		"is_required",
		"option_value",
		"is_active",
		"is_deleted",
		"deleted_at",
		"created_at",
		"updated_at",
	}
}

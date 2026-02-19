package replica

import "time"

type CustomField struct {
	ID          int       `json:"_id"         column:"custom_fields.id"           sqlType:"BIGINT"`
	Name        string    `json:"name"        column:"custom_fields.name"         sqlType:"VARCHAR"`
	Type        string    `json:"type"        column:"custom_fields.type"         sqlType:"VARCHAR"`
	DataType    string    `json:"dataType"    column:"custom_fields.data_type"    sqlType:"VARCHAR"`
	IsRequired  bool      `json:"isRequired"  column:"custom_fields.is_required"  sqlType:"BOOLEAN"`
	OptionValue string    `json:"optionValue" column:"custom_fields.option_value" sqlType:"TEXT[]"`
	IsActive    bool      `json:"isActive"    column:"custom_fields.is_active"    sqlType:"BOOLEAN"`
	IsDeleted   bool      `json:"isDeleted"   column:"custom_fields.is_deleted"   sqlType:"BOOLEAN"`
	DeletedAt   time.Time `json:"deletedAt"   column:"custom_fields.deleted_at"   sqlType:"TIMESTAMPTZ"`
	CreatedAt   time.Time `json:"createdAt"   column:"custom_fields.created_at"   sqlType:"TIMESTAMPTZ"`
	UpdatedAt   time.Time `json:"updatedAt"   column:"custom_fields.updated_at"   sqlType:"TIMESTAMPTZ"`
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

package new_model

import (
	"time"
)

type ImportCustomField struct {
	ID           int                 `json:"_id"          column:"import_custom_fields.id"`
	Key          int                 `json:"key"          column:"import_custom_fields.key"           special:"generated"`
	IsRequired   string              `json:"isRequired"   column:"import_custom_fields.is_required"`
	OptionValue  string              `json:"optionValue"  column:"import_custom_fields.option_value"`
	DataType     CustomFieldDataType `json:"dataType"     column:"import_custom_fields.data_type"`
	Type         CustomFieldType     `json:"type"         column:"import_custom_fields.type"`
	Name         string              `json:"name"         column:"import_custom_fields.name"`
	Category     string              `json:"category"     column:"import_custom_fields.category"`
	CategoryCode string              `json:"categoryCode" column:"import_custom_fields.category_code"`
	CreatedAt    time.Time           `json:"createdAt"    column:"import_custom_fields.created_at"`
	UpdatedAt    time.Time           `json:"updatedAt"    column:"import_custom_fields.updated_at"`
}

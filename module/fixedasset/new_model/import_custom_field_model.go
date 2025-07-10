package new_model

import "time"

type ImportCustomField struct {
	ID int `json:"id" column:"id"`

	Name         string              `json:"name"          column:"name"`
	Type         CustomFieldType     `json:"type"          column:"type"`
	DataType     CustomFieldDataType `json:"dataType"     column:"data_type"`
	OptionValue  string              `json:"optionValue"  column:"option_value"`
	IsRequired   string              `json:"isRequired"   column:"is_required"`
	Category     string              `json:"category"      column:"category"`
	CategoryCode string              `json:"categoryCode" column:"category_code"` // not used in basic

	Key int `json:"key" column:"key"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

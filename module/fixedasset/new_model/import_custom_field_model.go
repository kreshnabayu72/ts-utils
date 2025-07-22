package new_model

import "time"

type ImportCustomField struct {
	ID int `json:"_id" column:"id"`

	IsRequired   string              `json:"isRequired"   column:"is_required"`
	OptionValue  string              `json:"optionValue"  column:"option_value"`
	DataType     CustomFieldDataType `json:"dataType"     column:"data_type"`
	Type         CustomFieldType     `json:"type"          column:"type"`
	Name         string              `json:"name"          column:"name"`
	Category     string              `json:"category"      column:"category"`
	CategoryCode string              `json:"categoryCode" column:"category_code"` // not used in basic

	Key int `json:"key" column:"key"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

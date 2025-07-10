package new_model

import "time"

type CustomFieldCategory struct {
	ID int `json:"id" column:"id"`

	CategoryId    int `json:"category_id"    column:"category_id"`
	CustomFieldId int `json:"custom_field_id" column:"custom_field_id"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

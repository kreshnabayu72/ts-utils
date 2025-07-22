package new_model

import "time"

type CustomFieldCategory struct {
	ID int `json:"_id" column:"id"`

	CategoryId    int `json:"categoryId"    column:"category_id"`
	CustomFieldId int `json:"customFieldId" column:"custom_field_id"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

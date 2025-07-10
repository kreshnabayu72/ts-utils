package new_model

import "time"

type ImportCategory struct {
	ID           int    `json:"id"            column:"id"`
	Name         string `json:"name"          column:"name"`
	CategoryCode string `json:"categoryCode" column:"category_code"`
	Key          int    `json:"key"           column:"key"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

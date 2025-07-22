package new_model

import "time"

type ImportBrand struct {
	ID  int `json:"id"            column:"id"`
	Key int `json:"key"           column:"key"`

	Name         string `json:"name"          column:"name"`
	Category     string `json:"category"          column:"category"`
	CategoryCode string `json:"categoryCode" column:"category_code"`
	Models       string `json:"models" column:"models"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

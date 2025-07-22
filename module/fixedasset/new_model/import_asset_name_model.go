package new_model

import "time"

type ImportAssetName struct {
	ID  int `json:"_id"            column:"id"`
	Key int `json:"key"           column:"key"`

	Name         string `json:"name"          column:"name"`
	Code         string `json:"code"          column:"code"`
	Category     string `json:"category"          column:"category"`
	CategoryCode string `json:"categoryCode" column:"category_code"`
	Measurement  string `json:"measurement" column:"measurement"`

	CreatedAt time.Time `json:"createdAt" column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at"`
}

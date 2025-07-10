package new_model

import "time"

type BrandCategory struct {
	ID int `json:"id"         column:"id"`

	BrandId    int `json:"brandId"         column:"brand_id"`
	CategoryId int `json:"categoryId"         column:"category_id"`

	CreatedAt time.Time `json:"createdAt"            column:"created_at"`
	UpdatedAt time.Time `json:"updatedAt"            column:"updated_at"`
}

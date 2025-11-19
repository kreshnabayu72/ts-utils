package model

import (
	"time"
)

type BrandCategory struct {
	ID         int        `json:"_id"        column:"brand_categories.id"`
	BrandId    int        `json:"brandId"    column:"brand_categories.brand_id"`
	CategoryId int        `json:"categoryId" column:"brand_categories.category_id"`
	CreatedAt  time.Time  `json:"createdAt"  column:"brand_categories.created_at"`
	UpdatedAt  time.Time  `json:"updatedAt"  column:"brand_categories.updated_at"`
	IsDeleted  bool       `json:"isDeleted"  column:"brand_categories.is_deleted"`
	DeletedAt  *time.Time `json:"deletedAt"  column:"brand_categories.deleted_at"`
}

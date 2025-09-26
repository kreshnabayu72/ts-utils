package new_model

import (
	"time"
)

type Model struct {
	ID              int        `json:"_id"             column:"models.id"`
	Name            string     `json:"name"            column:"models.name"`
	BrandCategoryId int        `json:"brandCategoryId" column:"models.brand_category_id"`
	IsDeleted       bool       `json:"isDeleted"       column:"models.is_deleted"`
	DeletedAt       *time.Time `json:"deletedAt"       column:"models.deleted_at"`
	CreatedAt       time.Time  `json:"createdAt"       column:"models.created_at"`
	UpdatedAt       time.Time  `json:"updatedAt"       column:"models.updated_at"`
}

package dto

type SetSoftDelete struct {
	IsDeleted bool   `column:"is_deleted"`
	DeletedAt string `column:"deleted_at"`
}

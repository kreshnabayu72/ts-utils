package dto

type PaginationResult[T any] struct {
	TotalRecords int32 `json:"totalRecords" bson:"totalRecords"`
	Data         []T   `json:"data"         bson:"data"`
}

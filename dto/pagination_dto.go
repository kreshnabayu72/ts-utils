package dto

type InterfaceOptionDTO struct {
	Label any `json:"label" bson:"label"`
	Value any `json:"value" bson:"value"`
}

type PaginationResult[T any] struct {
	TotalRecords int `json:"totalRecords" bson:"totalRecords"`
	Data         []T `json:"data"         bson:"data"`
}

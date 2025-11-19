package dto

type DeleteManyResponse struct {
	Success int      `json:"success" bson:"success"`
	Errors  []string `json:"errors"  bson:"errors"`
}

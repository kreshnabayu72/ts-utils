package dto

type StringOptionDTO struct {
	Label string `json:"label" bson:"label"`
	Value string `json:"value" bson:"value"`
}

type IntOptionDTO struct {
	Label string `json:"label" bson:"label"`
	Value int    `json:"value" bson:"value"`
}

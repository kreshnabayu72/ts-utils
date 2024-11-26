package dto

type StringOptionDTO struct {
	Label string `json:"label" bson:"label"`
	Value string `json:"value" bson:"value"`
}

type IntOptionDTO struct {
	Label string `json:"label" bson:"label"`
	Value int    `json:"value" bson:"value"`
}

type InterfaceOptionDTO struct {
	Label interface{} `json:"label" bson:"label"`
	Value interface{} `json:"value" bson:"value"`
}

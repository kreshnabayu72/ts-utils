package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AttributeEmbedded struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"          bson:"name"`
	Key  int                `json:"key"           bson:"key"`
}

type TreeAttributeEmbedded struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name"          bson:"name"`
	FullPath string             `json:"fullPath"      bson:"fullPath"`
	Key      int                `json:"key"           bson:"key"`
}

type UserAttributeEmbedded struct {
	ID             primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	FullName       string             `json:"fullName"       bson:"fullName"`
	ProfilePicture *string            `json:"profilePicture" bson:"profilePicture"`
	Key            int                `json:"key"            bson:"key"`
}

type SKUAttributeEmbedded struct {
	ID    primitive.ObjectID `json:"_id,omitempty"  bson:"_id,omitempty"`
	SKU   string             `json:"sku"            bson:"sku"`
	Image *string            `json:"image"          bson:"image"`
	Key   int                `json:"key"            bson:"key"`
}

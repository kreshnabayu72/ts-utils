package pipeline

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateStringArrayFilter(filter []string) primitive.M {
	if len(filter) > 0 {
		return bson.M{"$in": filter}
	} else {
		return bson.M{"$ne": 0}
	}
}

func GenerateIntArrayFilter(filter []int) primitive.M {
	if len(filter) > 0 {
		return bson.M{"$in": filter}
	} else {
		return bson.M{"$ne": ""}
	}
}
func GenerateSearchCondition(search string) primitive.M {
	if search != "" {
		return bson.M{"$regex": search, "$options": "i"}
	} else {
		return bson.M{"$ne": false}
	}
}

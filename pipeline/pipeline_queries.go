package pipeline

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateExactFilter[T comparable](condition bool, filter T) primitive.M {
	if condition {
		return bson.M{"$eq": filter}
	} else {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}
}

func GenerateArrayFilter[T comparable](filter []T) primitive.M {
	if len(filter) > 0 {
		return bson.M{"$in": filter}
	} else {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}
}
func GenerateSearchCondition(search string) primitive.M {
	if search != "" {
		return bson.M{"$regex": search, "$options": "i"}
	} else {
		return bson.M{"$ne": false}
	}
}

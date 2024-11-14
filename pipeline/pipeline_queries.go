package pipeline

import (
	"time"

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

func GenerateObjectIdFilter(filter string) primitive.M {
	objId, _ := primitive.ObjectIDFromHex(filter)
	if filter != "" {
		return bson.M{"$eq": objId}
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

func SetDateQuery(requestDate []int) primitive.M {
	if requestDate[0] >= 0 && requestDate[1] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		end := time.UnixMilli(int64(requestDate[1]))

		// Add 24 hours minus 1 millisecond to end time
		end = end.Add(24*60*60*1000 - 1)

		return bson.M{
			"$gte": start,
			"$lte": end,
		}
	}
	if requestDate[0] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		start2 := time.UnixMilli(int64(requestDate[0]))
		start2 = start2.Add(24*60*60*1000 - 1)

		return bson.M{
			"$gte": start,
			"$lt":  start2,
		}
	}

	return bson.M{"$ne": "g#*a,Jvb&U.Bg"}
}

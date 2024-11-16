package pipeline

import (
	"strings"
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
	if filter == "" {
		return bson.M{"$ne": "vdjfPyv7ijO5vQeLIZmQHuzPO"} //random string so its always true
	}

	objId, err := primitive.ObjectIDFromHex(filter)
	if err != nil {
		return bson.M{"$eq": "djfPyv7ijO5vQeLIZmQHuzPO"}
	} else {
		return bson.M{"$eq": objId}
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

func GenerateDateFilter(requestDate []int) primitive.M {
	if len(requestDate) < 1 {
		return bson.M{"$ne": "g#*a,Jvb&U.Bg"}
	}

	if len(requestDate) == 2 && requestDate[0] >= 0 && requestDate[1] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		end := time.UnixMilli(int64(requestDate[1]))

		// Add 24 hours minus 1 millisecond to end time
		end = end.Add(24*time.Hour - time.Second)

		return bson.M{
			"$gte": start,
			"$lte": end,
		}
	}
	if len(requestDate) == 1 && requestDate[0] >= 0 {
		start := time.UnixMilli(int64(requestDate[0]))
		start2 := time.UnixMilli(int64(requestDate[0]))

		start2 = start2.Add(24*time.Hour - time.Second)

		return bson.M{
			"$gte": start,
			"$lt":  start2,
		}
	}

	return bson.M{"$ne": "g#*a,Jvb&U.Bg"}
}

func GenerateOptionFacet(
	fieldName string,
	optionsQuery bool,
	labelKey string,
	valueKey string,
) primitive.A {
	lk := labelKey
	if !strings.HasPrefix(lk, "$") {
		lk = "$" + lk
	}

	vk := valueKey
	if !strings.HasPrefix(vk, "$") {
		vk = "$" + vk
	}

	return bson.A{
		bson.M{
			"$match": bson.M{
				fieldName: bson.M{"$ne": nil},
				"$expr": bson.M{
					"$eq": bson.A{optionsQuery, true},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"label": lk,
					"value": vk,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id.label": 1,
			},
		},
	}
}

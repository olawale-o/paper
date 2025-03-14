package utils

import (
	"go-simple-rest/src/v1/articles/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseParamToPrimitiveObjectId(param string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(param)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return oid, nil
}

func HandleQueryParams(params model.QueryParams) bson.M {
	fieldValues := map[string]int{
		"asc":  1,
		"desc": -1,
	}
	filter := bson.M{}
	if params.Date == "desc" {
		filter["createdAtTimestamp"] = -1
	} else {
		filter["createdAtTimestamp"] = fieldValues[params.Date]
	}
	// if params.Likes == "desc" {
	// 	filter["params.Likes"] = -1
	// } else {
	// 	filter["params.Likes"] = fieldValues[params.Likes]
	// }
	// if params.Views == "desc" {
	// 	filter["params.Views"] = -1
	// } else {
	// 	filter["params.Views"] = fieldValues[params.Views]
	// }
	return filter
}

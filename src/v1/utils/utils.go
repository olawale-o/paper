package utils

import (
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

func HandleQueryParams(date, likes, views string) bson.M {
	fieldValues := map[string]int{
		"asc":  1,
		"desc": -1,
	}
	filter := bson.M{}
	if date == "desc" {
		filter["createdAtTimestamp"] = -1
	} else {
		filter["createdAtTimestamp"] = fieldValues[date]
	}
	// if likes == "desc" {
	// 	filter["likes"] = -1
	// } else {
	// 	filter["likes"] = fieldValues[likes]
	// }
	// if views == "desc" {
	// 	filter["views"] = -1
	// } else {
	// 	filter["views"] = fieldValues[views]
	// }
	return filter
}

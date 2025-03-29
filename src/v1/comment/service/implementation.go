package service

import (
	"context"
	"errors"
	"fmt"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/comment/repo"
	"log"
	"slices"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceManager struct {
	repo repo.Repository
}

func New(repo repo.Repository) (Service, error) {
	return &ServiceManager{repo: repo}, nil
}

func (sm *ServiceManager) NewComment(c model.Comment, articleId primitive.ObjectID, userId primitive.ObjectID) (error, interface{}) {
	var article bson.M
	var opts bson.M

	filter := bson.M{"_id": articleId}
	data, err := sm.repo.FindOne(context.TODO(), "articles", filter, article, opts)
	if err != nil {
		log.Println(err)
		return err, "Article not found"
	}

	if data == nil {
		return err, nil
	}

	doc := model.Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: userId, LIKES: 0, CREATEDAT: primitive.NewDateTimeFromTime(time.Now()), UPDATEDAT: primitive.NewDateTimeFromTime(time.Now()), STATUS: "pending", PARENTCOMMENTID: c.PARENTCOMMENTID, CREATEDATIMESTAMP: time.Now().Local().UnixMilli(), UPDATEDATIMESTAMP: time.Now().Local().UnixMilli()}
	res, err := sm.repo.InsertOne(context.TODO(), "article_comments", doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}

	return err, res
}

func (sm *ServiceManager) GetComment(articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) (interface{}, error) {
	// var comment bson.M
	// var opts bson.M

	// filter := bson.M{"_id": commentId, "articleId": articleId}
	// data, err := sm.repo.FindOne(context.TODO(), "article_comments", filter, comment, opts)
	// if err != nil {
	// 	log.Println(err)
	// 	if err == mongo.ErrNoDocuments {
	// 		return nil, errors.New("Comment not found")
	// 	}
	// 	return nil, err
	// }

	// if data == nil {
	// 	return nil, err
	// }

	data, err := _FetchCommentWithReplies(sm.repo, articleId, commentId, next)
	return data, err
}

func (sm *ServiceManager) GetComments(articleId primitive.ObjectID, l int, prev string, next string) (Response, error) {
	data, err := _HandlePaginate(sm.repo, articleId, l, prev, next)
	if err != nil {
		return Response{}, err
	}

	return data, nil
}

func (sm *ServiceManager) ReplyComment(c model.Comment, articleId primitive.ObjectID, commentId primitive.ObjectID, userId primitive.ObjectID) (interface{}, error) {
	var comment bson.M
	var opts bson.M

	filter := bson.M{"_id": commentId, "articleId": articleId}
	data, err := sm.repo.FindOne(context.TODO(), "article_comments", filter, comment, opts)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Comment not found")
		}
		return nil, err
	}

	doc := model.Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: userId, LIKES: 0, CREATEDAT: primitive.NewDateTimeFromTime(time.Now()), UPDATEDAT: primitive.NewDateTimeFromTime(time.Now()), STATUS: "pending", PARENTCOMMENTID: commentId, CREATEDATIMESTAMP: time.Now().Local().UnixMilli(), UPDATEDATIMESTAMP: time.Now().Local().UnixMilli()}
	data, err = sm.repo.InsertOne(context.TODO(), "article_comments", doc)

	if err != nil {
		return nil, err
	}

	val, ok := data.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("Invalid data type")
	}

	data, err = sm.repo.UpdateOne(
		context.TODO(),
		"article_comments",
		filter,
		bson.M{"$push": bson.M{"replies": bson.M{"$each": []model.Reply{
			model.Reply{ID: val, USERID: userId, CREATEDATTIMESTAMP: int(doc.CREATEDATIMESTAMP)},
		},
			"$slice": 2,
			"$sort":  bson.M{"createdAtTimestamp": -1},
		},
		}},
		false,
	)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	return data, err
}

func (sm *ServiceManager) ArticleComments(articleId primitive.ObjectID, next primitive.ObjectID) ([]model.ArticleWithComments, interface{}, error) {
	var matchStage bson.M
	var limitStage bson.M
	var unwindStage bson.M = bson.M{"$unwind": bson.M{"path": "$replies", "preserveNullAndEmptyArrays": true}}
	var sortStage bson.M = bson.M{"$sort": bson.M{"createdAtTimestamp": -1}}
	var lookupStage bson.M = bson.M{
		"$lookup": bson.M{
			"from":         "article_comments",
			"localField":   "replies._id",
			"foreignField": "_id",
			"as":           "comment_replies",
		},
	}

	var projectStage bson.M = bson.M{
		"$project": bson.M{
			"body":               1,
			"articleId":          1,
			"createdAtTimestamp": 1,
			"reply":              bson.M{"$arrayElemAt": []interface{}{"$comment_replies", 0}},
		},
	}

	var groupStage bson.M = bson.M{
		"$group": bson.M{
			"_id":                bson.M{"id": "$_id", "parentId": "$reply.parentCommentId"},
			"id":                 bson.M{"$first": "$_id"},
			"body":               bson.M{"$first": "$body"},
			"articleId":          bson.M{"$first": "$articleId"},
			"createdAtTimestamp": bson.M{"$first": "$createdAtTimestamp"},
			"replies":            bson.M{"$push": bson.M{"commentId": "$reply._id", "body": "$reply.body", "articleId": "$reply.articleId", "createdAtTimestamp": "$reply.createdAtTimestamp"}},
		},
	}

	if next != primitive.NilObjectID {
		matchStage = bson.M{
			"$match": bson.M{"$and": []bson.M{
				bson.M{"articleId": articleId},
				bson.M{"parentCommentId": nil},
				bson.M{"_id": bson.M{"$lte": next}},
			},
			},
		}
		limitStage = bson.M{"$limit": 10}
	} else {
		matchStage = bson.M{
			"$match": bson.M{"$and": []bson.M{
				bson.M{"articleId": articleId},
				bson.M{"parentCommentId": nil},
			},
			},
		}
		limitStage = bson.M{"$limit": 3}
	}

	pipeline := []bson.M{
		matchStage,
		unwindStage,
		lookupStage,
		projectStage,
		groupStage,
		sortStage,
		limitStage,
	}

	data, err := sm.repo.Aggregate(context.TODO(), "article_comments", pipeline)
	if err != nil {
		return nil, "", err
	}

	if len(data) > 2 && next == primitive.NilObjectID {
		return data[:2], data[2].ID, nil
	}

	return data, nil, nil
}

func (sm *ServiceManager) MoreReplies(articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) ([]model.Comment, error) {
	var comments []model.Comment
	data, err := _HandleMoreReplies(sm.repo, articleId, commentId, next)

	if err != nil {
		return comments, err
	}

	return data, nil
}

func _HandleMoreReplies(repository repo.Repository, articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) ([]model.Comment, error) {

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"articleId": articleId},
			bson.M{"parentCommentId": commentId},
			bson.M{"_id": bson.M{"$lt": next}},
		},
	}

	sort := bson.M{"createdAtTimestamp": -1}
	limt := int64(0)

	data, err := repository.Find(context.TODO(), "article_comments", filter, sort, limt)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func _HandlePaginate(repository repo.Repository, articleId primitive.ObjectID, l int, prev string, next string) (Response, error) {
	var sort bson.M = bson.M{"_id": -1}
	var filter bson.M = bson.M{"articleId": articleId}
	var limit int64
	var hasPrev bool
	var hasNext bool
	var lastId primitive.ObjectID
	var firstId primitive.ObjectID

	if prev != "" {
		id, _ := primitive.ObjectIDFromHex(prev)
		filter["_id"] = bson.M{"$gt": id}
		sort["_id"] = 1
	} else if next != "" {
		id, _ := primitive.ObjectIDFromHex(next)
		filter["_id"] = bson.M{"$lt": id}
	}

	if l < 1 {
		limit = int64(1)
	}

	if l > 20 {
		limit = int64(10)
	}

	result, err := repository.Find(context.TODO(), "article_comments", filter, sort, limit)

	if err != nil {
		return Response{}, err
	}

	if len(result) > 0 {
		var opts bson.M = bson.M{}
		var nextComment bson.M
		lastId = result[len(result)-1].ID.(primitive.ObjectID)
		firstId = result[0].ID.(primitive.ObjectID)
		filter["_id"] = bson.M{"$lt": lastId}
		nxtComment, _ := repository.FindOne(context.TODO(), "article_comments", filter, nextComment, opts)
		if nxtComment != nil {
			hasNext = true
		}

		var prevComment bson.M
		filter["_id"] = bson.M{"$gt": firstId}
		prvComment, _ := repository.FindOne(context.TODO(), "article_comments", filter, prevComment, opts)
		if prvComment != nil {
			hasPrev = true
		}
	}

	if prev != "" && hasPrev {
		slices.Reverse(result)
		return Response{Comments: result, HasNext: hasNext, HasPrev: hasPrev, PreviousID: firstId.Hex(), NextID: lastId.Hex()}, nil
	}
	if !hasPrev {
		return Response{Comments: result, HasNext: hasNext, HasPrev: hasPrev, NextID: firstId.Hex()}, nil
	}

	if !hasNext {
		return Response{Comments: result, HasNext: hasNext, HasPrev: hasPrev, PreviousID: lastId.Hex()}, nil
	}

	return Response{Comments: result, HasNext: hasNext, HasPrev: hasPrev, PreviousID: firstId.Hex(), NextID: lastId.Hex()}, nil
}

func _FetchCommentWithReplies(repository repo.Repository, articleId primitive.ObjectID, commentId primitive.ObjectID, next primitive.ObjectID) ([]model.Comment, error) {
	fmt.Println(articleId)
	fmt.Println(commentId)
	fmt.Println(next)
	filter := bson.M{
		"$and": []bson.M{
			bson.M{"articleId": articleId},
			bson.M{"parentCommentId": commentId},
			bson.M{"_id": bson.M{"$lt": next}},
		},
	}
	sort := bson.M{"createdAtTimestamp": -1}
	limit := int64(10)

	result, err := repository.Find(context.TODO(), "article_comments", filter, sort, limit)

	if err != nil {
		return nil, err
	}

	return result, nil
}

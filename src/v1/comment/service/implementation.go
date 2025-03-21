package service

import (
	"context"
	"errors"
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

func (sm *ServiceManager) NewComment(c model.Comment, articleId primitive.ObjectID) (error, interface{}) {
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

	doc := model.Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: c.USERID, LIKES: 0, CREATEDAT: primitive.NewDateTimeFromTime(time.Now()), UPDATEDAT: primitive.NewDateTimeFromTime(time.Now()), STATUS: "pending", PARENTCOMMENTID: c.PARENTCOMMENTID, CREATEDATIMESTAMP: time.Now().Local().UnixMilli(), UPDATEDATIMESTAMP: time.Now().Local().UnixMilli()}
	res, err := sm.repo.InsertOne(context.TODO(), "comments", doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}

	return err, res
}

func (sm *ServiceManager) GetComment(articleId primitive.ObjectID, commentId primitive.ObjectID) (interface{}, error) {
	var comment bson.M
	var opts bson.M

	filter := bson.M{"_id": commentId, "articleId": articleId}
	data, err := sm.repo.FindOne(context.TODO(), "comments", filter, comment, opts)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Comment not found")
		}
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	return data, err
}

func (sm *ServiceManager) GetComments(articleId primitive.ObjectID, l int, prev string, next string) (Response, error) {
	data, err := _HandlePaginate(sm.repo, articleId, l, prev, next)
	if err != nil {
		return Response{}, err
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

	result, err := repository.Get(context.TODO(), "comments", filter, sort, limit)

	if err != nil {
		return Response{}, err
	}

	if len(result) > 0 {
		var opts bson.M = bson.M{}
		var nextComment bson.M
		lastId = result[len(result)-1].ID.(primitive.ObjectID)
		firstId = result[0].ID.(primitive.ObjectID)
		filter["_id"] = bson.M{"$lt": lastId}
		nxtComment, _ := repository.FindOne(context.TODO(), "comments", filter, nextComment, opts)
		if nxtComment != nil {
			hasNext = true
		}

		var prevComment bson.M
		filter["_id"] = bson.M{"$gt": firstId}
		prvComment, _ := repository.FindOne(context.TODO(), "comments", filter, prevComment, opts)
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

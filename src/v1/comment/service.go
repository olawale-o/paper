package comment

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/comment/repo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()
var database = client.Database("go")

var articleCollection = client.Database("go").Collection("articles")
var collection = client.Database("go").Collection("comments")

type Response struct {
	Comments []model.Comment `json:"comments"`
	HasNext  bool            `json:"hasNext"`
	HasPrev  bool            `json:"hasPrev"`
}

func NewComment(c model.Comment, articleId primitive.ObjectID) (error, interface{}) {
	var article model.Article

	filter := bson.M{"_id": articleId}
	if err := articleCollection.FindOne(context.TODO(), filter).Decode(&article); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return err, "Article not found"
		}
		return err, "Article not found"
	}

	doc := model.Comment{BODY: c.BODY, ARTICLEID: articleId, USERID: c.USERID, LIKES: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), STATUS: "pending", PARENTCOMMENTID: c.PARENTCOMMENTID}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return err, ""
	}
	id := res.InsertedID

	return err, id
}

func GetComment(articleId primitive.ObjectID, commentId primitive.ObjectID) (error, interface{}) {
	var comment model.Comment

	filter := bson.M{"_id": commentId, "articleId": articleId}
	if err := collection.FindOne(context.TODO(), filter).Decode(&comment); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return err, "Comment not found"
		}
		return err, "Comment not found"
	}

	return err, comment
}

func GetComments(articleId primitive.ObjectID, l int) (Response, error) {
	var filter bson.M
	// var comments []model.Comment
	var limit int64
	var hasPrev bool
	var hasNext bool
	var lastId primitive.ObjectID
	var firstId primitive.ObjectID
	r, err := repo.New(database)

	if err != nil {
		return Response{}, err
	}

	if l < 1 {
		limit = int64(4)
	}

	if l > 20 {
		limit = int64(10)
	}

	filter = bson.M{"articleId": articleId}
	sort := bson.M{"_id": -1}
	result, err := r.Get(context.TODO(), "comments", filter, sort, limit)

	if err != nil {
		return Response{}, err
	}

	if len(result) > 0 {
		var nextComment bson.M
		lastId = result[len(result)-1].ID.(primitive.ObjectID)
		firstId = result[0].ID.(primitive.ObjectID)
		filter = bson.M{"articleId": articleId, "_id": bson.M{"$lt": lastId}}
		nxtComment, _ := r.FindOne(context.TODO(), "comments", filter, nextComment)
		if nxtComment != nil {
			hasNext = true
		}

		var prevComment bson.M
		filter = bson.M{"articleId": articleId, "_id": bson.M{"$gt": firstId}}
		prvComment, _ := r.FindOne(context.TODO(), "comments", filter, prevComment)
		if prvComment != nil {
			hasPrev = true
		}

	}

	return Response{Comments: result, HasNext: hasNext, HasPrev: hasPrev}, nil
}

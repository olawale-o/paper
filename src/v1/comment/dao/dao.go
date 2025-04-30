package dao

import (
	"context"
	"go-simple-rest/src/v1/comment/model"
	"go-simple-rest/src/v1/comment/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "article_comments"
const articleCollection = "articles"

type CommentDao interface {
	Create(comment model.Comment) (interface{}, error)
	FindAll(filter, sort bson.M, limit int64) ([]model.Comment, error)
	FindById(filter bson.M) (model.Comment, error)
	FindArticleById(id primitive.ObjectID) (interface{}, error)
	FindCommentByIdWithReplies(articleId, commentId, nextId primitive.ObjectID) ([]model.Comment, error)
	UpdateCommentWithReply(id, articleId primitive.ObjectID, update bson.M) (interface{}, error)
	Aggregate(pipeline []bson.M) ([]model.ArticleWithComments, error)
}

type MongoDBCommentDaoManager struct {
	repo repo.Repository
}

func NewCommentDaoManager(repo repo.Repository) CommentDao {
	return &MongoDBCommentDaoManager{repo: repo}
}

func (d *MongoDBCommentDaoManager) Create(doc model.Comment) (interface{}, error) {
	res, err := d.repo.InsertOne(context.TODO(), collectionName, doc)
	return res, err
}

func (d *MongoDBCommentDaoManager) FindAll(filter, sort bson.M, limit int64) ([]model.Comment, error) {
	result, err := d.repo.Find(context.TODO(), collectionName, filter, sort, limit)
	return result, err
}

func (d *MongoDBCommentDaoManager) FindById(filter bson.M) (model.Comment, error) {
	var comment bson.M
	var opts bson.M

	data, err := d.repo.FindOne(context.TODO(), collectionName, filter, comment, opts)
	if err != nil {
		return model.Comment{}, err
	}
	return data.(model.Comment), nil
}

func (d *MongoDBCommentDaoManager) FindArticleById(id primitive.ObjectID) (interface{}, error) {
	var article bson.M
	var opts bson.M

	filter := bson.M{"_id": id}
	data, err := d.repo.FindOne(context.TODO(), articleCollection, filter, article, opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *MongoDBCommentDaoManager) FindCommentByIdWithReplies(articleId, commentId, nextId primitive.ObjectID) ([]model.Comment, error) {

	sort := bson.M{"createdAtTimestamp": -1}
	limit := int64(10)

	filter := bson.M{
		"$and": []bson.M{
			bson.M{"articleId": articleId},
			bson.M{"parentCommentId": commentId},
			bson.M{"_id": bson.M{"$lt": nextId}},
		},
	}
	data, err := d.repo.Find(context.TODO(), collectionName, filter, sort, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *MongoDBCommentDaoManager) UpdateCommentWithReply(id, articleId primitive.ObjectID, update bson.M) (interface{}, error) {
	filter := bson.M{"_id": id, "articleId": articleId}
	res, err := d.repo.UpdateOne(context.TODO(), articleCollection, filter, update, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *MongoDBCommentDaoManager) Aggregate(pipeline []bson.M) ([]model.ArticleWithComments, error) {
	data, err := d.repo.Aggregate(context.TODO(), collectionName, pipeline)
	return data, err
}

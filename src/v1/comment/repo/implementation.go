package repo

import (
	"context"
	"go-simple-rest/src/v1/comment/dao"
	"go-simple-rest/src/v1/comment/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "article_comments"
const articleCollection = "articles"

type MongoDBCommentDaoManager struct {
	dao dao.CommentDAO
}

func NewRepository(dao dao.CommentDAO) Repository {
	return &MongoDBCommentDaoManager{dao: dao}
}

func (d *MongoDBCommentDaoManager) Create(doc model.Comment) (any, error) {
	res, err := d.dao.InsertOne(context.TODO(), collectionName, doc)
	return res, err
}

func (d *MongoDBCommentDaoManager) FindAll(filter, sort bson.M, limit int64) ([]model.Comment, error) {
	result, err := d.dao.Find(context.TODO(), collectionName, filter, sort, limit)
	return result, err
}

func (d *MongoDBCommentDaoManager) FindById(filter bson.M) (model.Comment, error) {
	var comment bson.M
	var opts bson.M

	data, err := d.dao.FindOne(context.TODO(), collectionName, filter, comment, opts)
	if err != nil {
		return model.Comment{}, err
	}
	return data.(model.Comment), nil
}

func (d *MongoDBCommentDaoManager) FindArticleById(id primitive.ObjectID) (any, error) {
	var article bson.M
	var opts bson.M

	filter := bson.M{"_id": id}
	data, err := d.dao.FindOne(context.TODO(), articleCollection, filter, article, opts)
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
			//bson.M{"_id": bson.M{"$lt": nextId}},
		},
	}
	data, err := d.dao.Find(context.TODO(), collectionName, filter, sort, limit)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *MongoDBCommentDaoManager) UpdateCommentWithReply(id, articleId primitive.ObjectID, update bson.M) (any, error) {
	filter := bson.M{"_id": id, "articleId": articleId}
	res, err := d.dao.UpdateOne(context.TODO(), articleCollection, filter, update, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d *MongoDBCommentDaoManager) Aggregate(pipeline []bson.M) ([]model.ArticleWithComments, error) {
	data, err := d.dao.Aggregate(context.TODO(), collectionName, pipeline)
	return data, err
}

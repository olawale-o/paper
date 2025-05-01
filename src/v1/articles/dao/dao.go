package dao

import (
	"context"
	"go-simple-rest/src/v1/articles/model"
	"go-simple-rest/src/v1/articles/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "articles"

var exclude bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}

type ArticleDao interface {
	GetArticles(filter, sort bson.M) ([]model.Article, error)
	GetArticleById(articleId primitive.ObjectID) (model.Article, error)
	UpdateArticle(filter, update bson.M) (any, error)
}

type MongoDBArticleDaoManager struct {
	repo repository.Repository
}

func NewArticleDaoManager(repo repository.Repository) ArticleDao {
	return &MongoDBArticleDaoManager{repo: repo}
}

func (d *MongoDBArticleDaoManager) GetArticles(filter, sort bson.M) ([]model.Article, error) {
	articles, err := d.repo.Find(context.TODO(), collectionName, filter, sort, exclude)
	return articles, err
}

func (d *MongoDBArticleDaoManager) GetArticleById(articleId primitive.ObjectID) (model.Article, error) {
	filter := bson.M{"_id": articleId}
	data, err := d.repo.FindOne(context.TODO(), collectionName, filter, exclude)
	return data, err
}

func (d *MongoDBArticleDaoManager) UpdateArticle(filter, update bson.M) (any, error) {
	res, err := d.repo.UpdateOne(context.TODO(), collectionName, filter, update, true)
	return res, err
}

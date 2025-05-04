package repository

import (
	"context"
	"go-simple-rest/src/v1/articles/dao"
	"go-simple-rest/src/v1/articles/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "articles"

var exclude bson.M = bson.M{"deletedAt": 0, "tags": 0, "categories": 0}

type repository struct {
	dao dao.ArticleDAO
}

func NewRepositoryManager(articleDAO dao.ArticleDAO) Repostiory {
	return &repository{dao: articleDAO}
}

func (d *repository) GetArticles(filter, sort bson.M) ([]model.Article, error) {
	articles, err := d.dao.Find(context.TODO(), collectionName, filter, sort, exclude)
	return articles, err
}

func (d *repository) GetArticleById(articleId primitive.ObjectID) (model.Article, error) {
	filter := bson.M{"_id": articleId}
	data, err := d.dao.FindOne(context.TODO(), collectionName, filter, exclude)
	return data, err
}

func (d *repository) UpdateArticle(filter, update bson.M) (any, error) {
	res, err := d.dao.UpdateOne(context.TODO(), collectionName, filter, update, true)
	return res, err
}

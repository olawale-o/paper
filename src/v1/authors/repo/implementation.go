package repo

import (
	"context"
	"go-simple-rest/src/v1/authors/dao"
	"go-simple-rest/src/v1/authors/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type repository struct {
	dao dao.AuthorDAO
}

func NewRepository(dao dao.AuthorDAO) Repository {
	return &repository{dao: dao}
}

func (d *repository) UpdateArticleAuthor(filter, update bson.M) (model.AuthorArticleUpdateResponse, error) {
	articles, err := d.dao.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	return articles, err
}

func (d *repository) GetAuthorArticles(authorId primitive.ObjectID) ([]model.AuthorArticle, error) {
	articles, err := d.dao.Get(context.TODO(), "articles", bson.M{"authorId": authorId})
	return articles, err
}

func (d *repository) Create(doc model.AuthorArticle) (any, error) {
	insertedId, err := d.dao.InsertOne(context.TODO(), "articles", doc)
	return insertedId, err
}

func (d *repository) Update(filter, update bson.M) (model.AuthorArticleUpdateResponse, error) {
	res, err := d.dao.FindOneAndUpdate(context.TODO(), "articles", filter, update, false)
	return res, err
}

func (d *repository) Delete(articleId primitive.ObjectID) error {
	err := d.dao.DeleteOne(context.TODO(), "articles", bson.M{"_id": articleId})
	return err
}

func (d *repository) GetAuthorById(filter bson.M) (any, error) {
	var author model.Author
	data, err := d.dao.FindOne(context.TODO(), "users", filter, author)
	return data, err

}

func (d *repository) UpdateAuthor(filter, update bson.M) (any, error) {
	res, err := d.dao.UpdateOne(context.TODO(), "users", filter, update, true)
	return res, err
}
func (d *repository) DeleteAuthor(filter bson.M) error {
	err := d.dao.DeleteOne(context.TODO(), "users", filter)
	return err
}

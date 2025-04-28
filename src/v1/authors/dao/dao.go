package dao

import (
	"context"
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorDao interface {
	UpdateArticleAuthor(filter, update bson.M) (model.AuthorArticleUpdateResponse, error)
	GetArticlesByAuthor(authorId primitive.ObjectID) ([]model.AuthorArticle, error)
	Create(doc model.AuthorArticle) (interface{}, error)
	Update(filter, update bson.M) (model.AuthorArticleUpdateResponse, error)
	Delete(articleId primitive.ObjectID) error
	GetAuthorById(filter bson.M) (interface{}, error)
	UpdateAuthor(filter, update bson.M) (interface{}, error)
	DeleteAuthor(filter bson.M) error
}

type MongoDBAuthorDaoManager struct {
	repo repo.Repository
}

func NewArticleDaoManager(repo repo.Repository) AuthorDao {
	return &MongoDBAuthorDaoManager{repo: repo}
}

func (d *MongoDBAuthorDaoManager) UpdateArticleAuthor(filter, update bson.M) (model.AuthorArticleUpdateResponse, error) {
	articles, err := d.repo.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	return articles, err
}

func (d *MongoDBAuthorDaoManager) GetArticlesByAuthor(authorId primitive.ObjectID) ([]model.AuthorArticle, error) {
	articles, err := d.repo.Get(context.TODO(), "articles", bson.M{"authorId": authorId})
	return articles, err
}

func (d *MongoDBAuthorDaoManager) Create(doc model.AuthorArticle) (interface{}, error) {
	insertedId, err := d.repo.InsertOne(context.TODO(), "articles", doc)
	return insertedId, err
}

func (d *MongoDBAuthorDaoManager) Update(filter, update bson.M) (model.AuthorArticleUpdateResponse, error) {
	res, err := d.repo.FindOneAndUpdate(context.TODO(), "articles", filter, update, false)
	return res, err
}

func (d *MongoDBAuthorDaoManager) Delete(articleId primitive.ObjectID) error {
	err := d.repo.DeleteOne(context.TODO(), "articles", bson.M{"_id": articleId})
	return err
}

func (d *MongoDBAuthorDaoManager) GetAuthorById(filter bson.M) (interface{}, error) {
	var author model.Author
	data, err := d.repo.FindOne(context.TODO(), "users", filter, author)
	return data, err

}

func (d *MongoDBAuthorDaoManager) UpdateAuthor(filter, update bson.M) (interface{}, error) {
	res, err := d.repo.UpdateOne(context.TODO(), "users", filter, update, true)
	return res, err
}
func (d *MongoDBAuthorDaoManager) DeleteAuthor(filter bson.M) error {
	err := d.repo.DeleteOne(context.TODO(), "users", filter)
	return err
}

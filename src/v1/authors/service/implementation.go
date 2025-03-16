package service

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var database = client.Database("go")

type ServiceManager struct {
	repo repo.Repository
}

func New(repo repo.Repository) (Service, error) {
	return &ServiceManager{repo: repo}, nil
}

func (s *ServiceManager) AllArticles(authorId primitive.ObjectID) (interface{}, error) {

	articles, err := s.repo.Get(context.TODO(), "articles", bson.M{"authorId": authorId})

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *ServiceManager) CreateArticle(article model.AuthorArticle, authorId primitive.ObjectID) (interface{}, error) {

	doc := model.AuthorArticle{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), TAGS: article.TAGS, CATEGORIES: article.CATEGORIES}
	insertedId, err := s.repo.InsertOne(context.TODO(), "articles", doc)

	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.AuthorArticle{{TITLE: doc.TITLE, ID: insertedId, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}

	res, err := s.repo.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return res, err
}

func (s *ServiceManager) UpdateArticle(article model.AuthorArticle, authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}

	res, err := s.repo.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter = bson.M{"_id": articleId}
	update = bson.M{"$set": bson.M{"title": article.TITLE, "content": article.CONTENT}}
	res, err = s.repo.FindOneAndUpdate(context.TODO(), "articles", filter, update, false)
	if err != nil {
		log.Println(err)
		return "", err
	}
	data, _ := bson.MarshalExtJSON(res, false, false)
	return data, nil
}

func (s *ServiceManager) DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) error {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$pull": bson.M{"articles": bson.M{"_id": articleId}}}

	_, err = s.repo.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.repo.DeleteOne(context.TODO(), "articles", bson.M{"_id": articleId})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceManager) ShowAuthor(authorId primitive.ObjectID) (interface{}, error) {
	var author bson.M

	filter := bson.M{"_id": authorId}

	data, err := s.repo.FindOne(context.TODO(), "users", filter, author)

	if err != nil {
		return author, err
	}

	return data, nil
}

func (s *ServiceManager) UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (interface{}, error) {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$set": bson.M{"firstName": updatedAuthor.FIRSTNAME, "username": updatedAuthor.USERNAME, "lastName": updatedAuthor.LASTNAME}}

	result, err := s.repo.UpdateOne(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return updatedAuthor, err
		}
	}

	return result, nil
}

func (s *ServiceManager) DeleteAuthor(authorId primitive.ObjectID) (int64, error) {

	filter := bson.M{"_id": authorId}
	err = s.repo.DeleteOne(context.TODO(), "users", filter)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return 1, nil
}

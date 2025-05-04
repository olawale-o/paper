package service

import (
	"go-simple-rest/src/v1/authors/model"
	"go-simple-rest/src/v1/authors/repo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceManager struct {
	authorRepo repo.Repository
}

func New(authorRepo repo.Repository) (Service, error) {
	return &ServiceManager{authorRepo: authorRepo}, nil
}

func (s *ServiceManager) AllArticles(authorId primitive.ObjectID) (any, error) {

	articles, err := s.authorRepo.GetAuthorArticles(authorId)

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *ServiceManager) CreateArticle(article model.AuthorArticle, authorId primitive.ObjectID) (any, error) {

	doc := model.AuthorArticle{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), TAGS: article.TAGS, CATEGORIES: article.CATEGORIES}
	insertedId, err := s.authorRepo.Create(doc)

	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.AuthorArticle{{TITLE: doc.TITLE, ID: insertedId, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}

	_, err = s.authorRepo.UpdateArticleAuthor(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return insertedId, err
}

func (s *ServiceManager) UpdateArticle(article model.AuthorArticle, authorId primitive.ObjectID, articleId primitive.ObjectID) (any, error) {

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}

	_, err := s.authorRepo.Update(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter = bson.M{"_id": articleId}
	update = bson.M{"$set": bson.M{"title": article.TITLE, "content": article.CONTENT}}
	res, err := s.authorRepo.Update(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.ID, nil
}

func (s *ServiceManager) DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) error {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$pull": bson.M{"articles": bson.M{"_id": articleId}}}

	_, err := s.authorRepo.Update(filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.authorRepo.Delete(articleId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceManager) ShowAuthor(authorId primitive.ObjectID) (model.Author, error) {
	var author model.Author

	filter := bson.M{"_id": authorId}

	data, err := s.authorRepo.GetAuthorById(filter)

	author = data.(model.Author)

	if err != nil {
		return author, err
	}

	return author, nil
}

func (s *ServiceManager) UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (any, error) {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$set": bson.M{"firstName": updatedAuthor.FIRSTNAME, "username": updatedAuthor.USERNAME, "lastName": updatedAuthor.LASTNAME}}

	result, err := s.authorRepo.UpdateAuthor(filter, update)
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
	err := s.authorRepo.DeleteAuthor(filter)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return 1, nil
}

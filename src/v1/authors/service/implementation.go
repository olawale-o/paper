package service

import (
	"go-simple-rest/src/v1/authors/dao"
	"go-simple-rest/src/v1/authors/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceManager struct {
	authorDao dao.AuthorDao
}

func New(authorDao dao.AuthorDao) (Service, error) {
	return &ServiceManager{authorDao: authorDao}, nil
}

func (s *ServiceManager) AllArticles(authorId primitive.ObjectID) (any, error) {

	articles, err := s.authorDao.GetArticlesByAuthor(authorId)

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *ServiceManager) CreateArticle(article model.AuthorArticle, authorId primitive.ObjectID) (any, error) {

	doc := model.AuthorArticle{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), TAGS: article.TAGS, CATEGORIES: article.CATEGORIES}
	insertedId, err := s.authorDao.Create(doc)

	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.AuthorArticle{{TITLE: doc.TITLE, ID: insertedId, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}

	_, err = s.authorDao.UpdateArticleAuthor(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return insertedId, err
}

func (s *ServiceManager) UpdateArticle(article model.AuthorArticle, authorId primitive.ObjectID, articleId primitive.ObjectID) (any, error) {

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}

	_, err := s.authorDao.Update(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter = bson.M{"_id": articleId}
	update = bson.M{"$set": bson.M{"title": article.TITLE, "content": article.CONTENT}}
	res, err := s.authorDao.Update(filter, update)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return res.ID, nil
}

func (s *ServiceManager) DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) error {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$pull": bson.M{"articles": bson.M{"_id": articleId}}}

	_, err := s.authorDao.Update(filter, update)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.authorDao.Delete(articleId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *ServiceManager) ShowAuthor(authorId primitive.ObjectID) (model.Author, error) {
	var author model.Author

	filter := bson.M{"_id": authorId}

	data, err := s.authorDao.GetAuthorById(filter)

	author = data.(model.Author)

	if err != nil {
		return author, err
	}

	return author, nil
}

func (s *ServiceManager) UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (any, error) {

	filter := bson.M{"_id": authorId}
	update := bson.M{"$set": bson.M{"firstName": updatedAuthor.FIRSTNAME, "username": updatedAuthor.USERNAME, "lastName": updatedAuthor.LASTNAME}}

	result, err := s.authorDao.UpdateAuthor(filter, update)
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
	err := s.authorDao.DeleteAuthor(filter)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return 1, nil
}

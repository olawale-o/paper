package service

import (
	"authors/db"
	"authors/model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, err = db.Connect()

var articleCollection = client.Database("go").Collection("articles")
var userCollection = client.Database("go").Collection("users")

func AllArticles(authorId primitive.ObjectID) (interface{}, error) {

	cursor, err := articleCollection.Find(context.TODO(), bson.M{"authorId": authorId})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var articles []model.Article
	if err = cursor.All(context.TODO(), &articles); err != nil {
		panic(err)
	}
	return articles, nil
}

func CreateArticle(article model.Article, authorId primitive.ObjectID) (interface{}, error) {

	doc := model.Article{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), TAGS: article.TAGS, CATEGORIES: article.CATEGORIES}
	res, err := articleCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.Article{{TITLE: doc.TITLE, ID: res.InsertedID, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var updatedDoc interface{}
	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)

	return res.InsertedID, err
}

func UpdateArticle(article model.Article, authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}
	opts := options.FindOneAndUpdate().SetUpsert(true)

	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts)

	var updatedDoc model.Article
	err := articleCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": articleId}, bson.M{"$set": bson.M{"title": article.TITLE, "content": article.CONTENT}}).Decode(&updatedDoc)

	if err != nil {
		return "", err
	}
	res, _ := bson.MarshalExtJSON(updatedDoc, false, false)
	return res, nil
}

func DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {
	filter := bson.M{"_id": authorId}
	update := bson.M{"$pull": bson.M{"articles": bson.M{"_id": articleId}}}
	opts := options.FindOneAndUpdate().SetUpsert(true)

	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts)

	res, err := articleCollection.DeleteOne(context.TODO(), bson.M{"_id": articleId})
	if err != nil {
		log.Println(err)
		return "", err
	}

	return res.DeletedCount, nil
}

func ShowAuthors() ([]model.Author, error) {
	var authors []model.Author
	filter := bson.M{}
	cursor, err := userCollection.Find(context.TODO(), filter)

	if err = cursor.All(context.TODO(), &authors); err != nil {
		panic(err)
	}

	return authors, nil
}

func ShowAuthor(authorId primitive.ObjectID) (model.Author, error) {
	var author model.Author
	filter := bson.M{"_id": authorId}
	if err := userCollection.FindOne(context.TODO(), filter).Decode(&author); err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return author, err
		}
		return author, err
	}
	return author, nil
}

func UpdateAuthor(authorId primitive.ObjectID, updatedAuthor model.Author) (interface{}, error) {
	filter := bson.M{"_id": authorId}
	update := bson.M{"$set": bson.M{"firstName": updatedAuthor.FIRSTNAME, "username": updatedAuthor.USERNAME, "lastName": updatedAuthor.LASTNAME}}
	opts := options.Update().SetUpsert(true)

	result, err := userCollection.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Println(err)
		if err == mongo.ErrNoDocuments {
			return updatedAuthor, err
		}
	}

	return result.UpsertedID, nil
}

func DeleteAuthor(authorId primitive.ObjectID) (int64, error) {

	filter := bson.M{"_id": authorId}
	result, err := userCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	return result.DeletedCount, nil
}

func UpdateAuthorWithArticle(data model.ArticleData) (interface{}, error) {
	fmt.Println(data)
	authorId, err := primitive.ObjectIDFromHex(data.AUTHORID)
	artilcleId, err := primitive.ObjectIDFromHex(data.ID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []model.Article{{
			TITLE:      data.TITLE,
			ID:         artilcleId,
			CONTENT:    data.CONTENT,
			CREATEDAT:  data.CREATEDAT,
			UPDATEDAT:  data.UPDATEDAT,
			CATEGORIES: data.CATEGORIES,
			TAGS:       data.TAGS,
		}},
			"$sort":  bson.M{"createdAt": -1},
			"$slice": 2}},
		"$inc": bson.M{"articleCount": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var updatedDoc interface{}
	err = userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Println(updatedDoc)

	return updatedDoc, nil
}

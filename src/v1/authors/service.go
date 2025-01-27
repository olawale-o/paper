package authors

import (
	"context"
	"fmt"
	"go-simple-rest/src/v1/articles"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AllArticles(authorId primitive.ObjectID) (interface{}, error) {

	cursor, err := articleCollection.Find(context.TODO(), bson.M{"authorId": authorId})
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	var articles []articles.Article
	if err = cursor.All(context.TODO(), &articles); err != nil {
		panic(err)
	}
	return articles, nil
}

func CreateArtcile(article articles.Article, authorId primitive.ObjectID) (interface{}, error) {

	doc := articles.Article{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now()}
	res, err := articleCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []articles.Article{{TITLE: doc.TITLE, ID: res.InsertedID, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	var updatedDoc interface{}
	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)

	return res.InsertedID, err
}

func UpdateArticle(article articles.Article, authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}
	opts := options.FindOneAndUpdate().SetUpsert(true)

	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts)

	var updatedDoc articles.Article
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

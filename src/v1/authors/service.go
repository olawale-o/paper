package authors

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/articles"
	"go-simple-rest/src/v1/authors/repo"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var articleCollection = client.Database("go").Collection("articles")

var client, ctx, err = db.Connect()

var database = client.Database("go")

func AllArticles(authorId primitive.ObjectID) (interface{}, error) {

	r, err := repo.New(database)

	if err != nil {
		return nil, err
	}

	articles, err := r.Get(context.TODO(), "articles", bson.M{"authorId": authorId})

	if err != nil {
		return nil, err
	}
	return articles, nil
}

func CreateArticle(article articles.Article, authorId primitive.ObjectID) (interface{}, error) {

	r, err := repo.New(database)

	if err != nil {
		return nil, err

	}
	doc := articles.Article{TITLE: article.TITLE, AUTHORID: authorId, CONTENT: article.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now(), TAGS: article.TAGS, CATEGORIES: article.CATEGORIES}
	insertedId, err := r.InsertOne(context.TODO(), "articles", doc)

	if err != nil {
		log.Println(err)
		return "", err
	}

	filter := bson.M{"_id": authorId}
	update := bson.M{
		"$push": bson.M{"articles": bson.M{"$each": []articles.Article{{TITLE: doc.TITLE, ID: insertedId, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
		"$inc":  bson.M{"articleCount": 1}}

	res, err := r.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return res, err
}

func UpdateArticle(article articles.Article, authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {

	r, err := repo.New(database)

	if err != nil {
		return nil, err

	}

	filter := bson.M{"_id": authorId, "articles": bson.M{"$elemMatch": bson.M{"_id": articleId}}}
	update := bson.M{"$set": bson.M{"articles.$.title": article.TITLE, "articles.$.content": article.CONTENT}}

	res, err := r.FindOneAndUpdate(context.TODO(), "users", filter, update, true)
	if err != nil {
		log.Println(err)
		return "", err
	}

	filter = bson.M{"_id": articleId}
	update = bson.M{"$set": bson.M{"title": article.TITLE, "content": article.CONTENT}}
	res, err = r.FindOneAndUpdate(context.TODO(), "articles", filter, update, false)
	if err != nil {
		log.Println(err)
		return "", err
	}
	data, _ := bson.MarshalExtJSON(res, false, false)
	return data, nil
}

// func DeleteArticle(authorId primitive.ObjectID, articleId primitive.ObjectID) (interface{}, error) {
// 	filter := bson.M{"_id": authorId}
// 	update := bson.M{"$pull": bson.M{"articles": bson.M{"_id": articleId}}}
// 	opts := options.FindOneAndUpdate().SetUpsert(true)

// 	userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts)

// 	res, err := articleCollection.DeleteOne(context.TODO(), bson.M{"_id": articleId})
// 	if err != nil {
// 		log.Println(err)
// 		return "", err
// 	}

// 	return res.DeletedCount, nil
// }

// func ShowAuthors() ([]Author, error) {
// 	var authors []Author
// 	filter := bson.M{}
// 	cursor, err := userCollection.Find(context.TODO(), filter)

// 	if err = cursor.All(context.TODO(), &authors); err != nil {
// 		panic(err)
// 	}

// 	return authors, nil
// }

// func ShowAuthor(authorId primitive.ObjectID) (Author, error) {
// 	var author Author
// 	filter := bson.M{"_id": authorId}
// 	if err := userCollection.FindOne(context.TODO(), filter).Decode(&author); err != nil {
// 		log.Println(err)
// 		if err == mongo.ErrNoDocuments {
// 			return author, err
// 		}
// 		return author, err
// 	}
// 	return author, nil
// }

// func UpdateAuthor(authorId primitive.ObjectID, updatedAuthor Author) (interface{}, error) {
// 	filter := bson.M{"_id": authorId}
// 	update := bson.M{"$set": bson.M{"firstName": updatedAuthor.FIRSTNAME, "username": updatedAuthor.USERNAME, "lastName": updatedAuthor.LASTNAME}}
// 	opts := options.Update().SetUpsert(true)

// 	result, err := userCollection.UpdateOne(context.TODO(), filter, update, opts)
// 	if err != nil {
// 		log.Println(err)
// 		if err == mongo.ErrNoDocuments {
// 			return updatedAuthor, err
// 		}
// 	}

// 	return result.UpsertedID, nil
// }

// func DeleteAuthor(authorId primitive.ObjectID) (int64, error) {

// 	filter := bson.M{"_id": authorId}
// 	result, err := userCollection.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		log.Println(err)
// 		return 0, err
// 	}

// 	return result.DeletedCount, nil
// }

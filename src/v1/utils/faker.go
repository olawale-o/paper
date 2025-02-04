package main

import (
	"context"
	"fmt"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/auth"
	"time"

	"github.com/go-faker/faker/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserWithTags struct {
	FIRSTNAME string `faker:"first_name"`
	LASTNAME  string `faker:"last_name"`
	USERNAME  string `faker:"username"`
}
type ArticleWithTags struct {
	TILE    string `faker:"word"`
	CONTENT string `faker:"paragraph"`
}

type Article struct {
	ID        interface{}        `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE     string             `bson:"title" json:"title"`
	CONTENT   string             `bson:"content" json:"content"`
	AUTHORID  primitive.ObjectID `bson:"authorId,omitempty" json:"authorId,omitempty"`
	LIKES     int                `bson:"likes,omitempty" json:"likes,omitempty"`
	VIEWS     int                `bson:"views,omitempty" json:"views,omitempty"`
	CREATEDAT time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UPDATEDAT time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type User struct {
	FIRSTNAME         string    `bson:"firstName" json:"firstName"`
	LASTNAME          string    `bson:"lastName" json:"lastName"`
	USERNAME          string    `bson:"username" json:"username"`
	ARTICLECOUNT      int       `bson:"articleCount,omitempty" json:"articleCount,omitempty"`
	ARTICLELIKESCOUNT int       `bson:"articleLikesCount,omitempty" json:"articleLikesCount,omitempty"`
	CREATEDAT         string    `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	PASSWORD          string    `bson:"password" json:"password"`
	ARTICLES          []Article `bson:"articles,omitempty" json:"articles,omitempty"`
}

type UserWithUpdate struct {
	User
	PASSWORD string    `bson:"password" json:"password"`
	ARTICLES []Article `bson:"articles,omitempty" json:"articles,omitempty"`
}

var client, ctx, err = db.Connect()

var articleCollection = client.Database("go").Collection("articles")
var userCollection = client.Database("go").Collection("users")

func main() {

	u := UserWithTags{}
	a := ArticleWithTags{}
	var users []User

	for i := 0; i < 10; i++ {
		err := faker.FakeData(&u)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, User{FIRSTNAME: u.FIRSTNAME, LASTNAME: u.LASTNAME, USERNAME: u.USERNAME, PASSWORD: "password"})
	}

	for _, v := range users {
		hash, _ := auth.HashPassword(v.PASSWORD)

		res, _ := userCollection.InsertOne(context.TODO(), User{FIRSTNAME: v.FIRSTNAME, LASTNAME: v.LASTNAME, USERNAME: v.USERNAME, ARTICLECOUNT: 0, ARTICLELIKESCOUNT: 0, CREATEDAT: time.Now().Format(time.DateTime), PASSWORD: hash})

		for i := 0; i < 15; i++ {
			err := faker.FakeData(&a)
			if err != nil {
				fmt.Println(err)
			}
			authorId, _ := res.InsertedID.(primitive.ObjectID)
			doc := Article{TITLE: a.TILE, AUTHORID: authorId, CONTENT: a.CONTENT, LIKES: 0, VIEWS: 0, CREATEDAT: time.Now(), UPDATEDAT: time.Now()}
			res, err := articleCollection.InsertOne(context.TODO(), doc)

			filter := bson.M{"_id": authorId}
			update := bson.M{
				"$push": bson.M{"articles": bson.M{"$each": []Article{{TITLE: doc.TITLE, ID: res.InsertedID, CONTENT: doc.CONTENT, CREATEDAT: doc.CREATEDAT, UPDATEDAT: doc.UPDATEDAT, LIKES: doc.LIKES, VIEWS: doc.VIEWS}}, "$sort": bson.M{"createdAt": -1}, "$slice": 2}},
				"$inc":  bson.M{"articleCount": 1}}
			opts := options.FindOneAndUpdate().SetUpsert(true)
			var updatedDoc interface{}
			userCollection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDoc)
		}
	}

	fmt.Println("Done")

}

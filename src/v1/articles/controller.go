package articles

import (
	"go-simple-rest/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client, ctx, _ = db.Connect()

var collection = client.Database("go").Collection("articles")

type Article struct {
	ID     string `bson:"_id,omitempty" json:"id,omitempty"`
	TITLE  string `bson:"title" json:"title"`
	AUTHOR string `bson:"author" json:"author"`
}

func GetArticles(c *gin.Context) {
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	var articles []Article
	if err = cursor.All(ctx, &articles); err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, articles)
}

func NewArticle(c *gin.Context) {
	var newArticle Article
	if err := c.BindJSON(&newArticle); err != nil {
		log.Println(err)
		return
	}
	doc := Article{TITLE: newArticle.TITLE, AUTHOR: newArticle.AUTHOR}
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Println(err)
		return
	}
	id := res.InsertedID
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Article created", "id": id})
}

func ShowArticle(c *gin.Context) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var article Article
	filter := bson.M{"_id": oid}
	if err := collection.FindOne(ctx, filter).Decode(&article); err != nil {
		log.Println(err)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"article": article})
}

func UpdateArticle(c *gin.Context) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var updatedArticle Article
	if err := c.BindJSON(&updatedArticle); err != nil {
		log.Println(err)
		return
	}
	filter := bson.M{"_id": oid}
	update := bson.M{"$set": bson.M{"title": updatedArticle.TITLE, "author": updatedArticle.AUTHOR}}
	opts := options.Update().SetUpsert(true)

	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article updated", "article": result})
}

func DeleteArticle(c *gin.Context) {
	oid, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.M{"_id": oid}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Article not found"})
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Article deleted", "article": result})
}

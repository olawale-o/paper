package auth

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("users")

func Login(c *gin.Context) {
	var user User
	var dbUser User

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		panic(err)
	}

	err := collection.FindOne(context.TODO(), bson.D{{"username", user.USERNAME}}).Decode(&dbUser)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Kindly create an account"})
		}
		return
	}

	match := CheckPasswordHash(user.PASSWORD, dbUser.PASSWORD)
	if !match {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Please provide valid credentials"})
		return
	}

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.USERNAME, "role": "user"})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server error"})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": user.USERNAME})
}

func Register(c *gin.Context) {
	var user User
	var dbUser User

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		panic(err)
	}

	err := collection.FindOne(context.TODO(), bson.D{{"username", user.USERNAME}}).Decode(&dbUser)

	if dbUser.USERNAME != "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to create user"})
		return
	}
	hash, _ := HashPassword(user.PASSWORD)
	doc := User{FIRSTNAME: user.FIRSTNAME, LASTNAME: user.LASTNAME, USERNAME: user.USERNAME, PASSWORD: hash}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User created", "userId": res.InsertedID})
}

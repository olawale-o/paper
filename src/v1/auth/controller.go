package auth

import (
	"context"
	"go-simple-rest/db"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/jwt"
	"go-simple-rest/src/v1/translator"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client, ctx, err = db.Connect()

var collection = client.Database("go").Collection("users")

func Login(c *gin.Context) {
	var user LoginAuth
	var dbUser authors.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(user)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	err := collection.FindOne(context.TODO(), bson.M{"username": user.USERNAME}).Decode(&dbUser)

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

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.ID, "role": "user", "id": dbUser.ID})

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server error"})
		return
	}

	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": gin.H{
		"username": dbUser.USERNAME, "role": jwt.GetRole("user"),
		"id": dbUser.ID,
	}})
}

func Register(c *gin.Context) {
	var user RegisterAuth
	var dbUser authors.User

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		panic(err)
	}

	validate := validator.New()
	err = validate.Struct(user)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
		return
	}

	err := collection.FindOne(context.TODO(), bson.M{"username": user.USERNAME}).Decode(&dbUser)

	if dbUser.USERNAME != "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Unable to create user"})
		return
	}
	hash, _ := HashPassword(user.PASSWORD)
	doc := authors.User{FIRSTNAME: user.FIRSTNAME, LASTNAME: user.LASTNAME, USERNAME: user.USERNAME, PASSWORD: hash}
	res, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User created", "userId": res.InsertedID})
}

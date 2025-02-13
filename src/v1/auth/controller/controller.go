package controller

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/auth/implementation"
	"go-simple-rest/src/v1/auth/repo"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var client, ctx, err = db.Connect()

var database = client.Database("go")

func Login(c *gin.Context) {
	var user auth.LoginAuth
	var dbUser authors.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	msg, error := s.Login(c, user)

	if _, ok := error["err"]; ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}

	c.SetCookie("token", msg, 3600, "/", "127.0.0.1", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": gin.H{
		"username": dbUser.USERNAME, "role": jwt.GetRole("user"),
		"id": dbUser.ID,
	}})
}

func Register(c *gin.Context) {
	var user auth.RegisterAuth

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		panic(err)
	}

	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	msg, error := s.Register(c, user)

	if _, ok := error["err"]; ok {
		c.IndentedJSON(http.StatusInternalServerError, error)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": msg})
}

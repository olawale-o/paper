package controller

import (
	"fmt"
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

// Login godoc
// @Tags User Authentication
// @Summary Login user
// @Description Login user with username and password
// @Accept json
// @Param data body auth.LoginAuth true "User"
// @Produce json
// @Success 200 {object} auth.LoginResponse "Response"
// @Header 200 {string} Cookie "session_id"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var user auth.LoginAuth
	var dbUser authors.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	msg, error := s.Login(c, user)

	if _, ok := error["err"]; ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
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

package controller

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/auth/implementation"
	"go-simple-rest/src/v1/auth/model"
	repo "go-simple-rest/src/v1/auth/repo/implementation"
	"go-simple-rest/src/v1/jwt"
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
// @Param data body model.LoginAuth true "User"
// @Produce json
// @Success 200 {object} model.LoginResponse "Response"
// @Header 200 {string} Cookie "session_id"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var user model.LoginAuth

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	response, error := s.Login(c, user)

	if _, ok := error["err"]; ok {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.SetCookie("token", response.TOKEN, 3600, "/", "127.0.0.1", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Logged in successfully", "user": gin.H{
		"username": response.USER.USERNAME, "role": jwt.GetRole("user"),
		"id": response.USER.ID,
	}})
}

// Sign up godoc
// @Tags User Authentication
// @Summary Register user
// @Description Register user with username,password, firstname and lastname
// @Accept json
// @Param data body model.RegisterAuth true "User"
// @Produce json
// @Success 201 {object} string "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /auth/sign-up [post]
func Register(c *gin.Context) {
	var user model.RegisterAuth

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	msg, error := s.Register(c, user)

	if _, ok := error["err"]; ok {
		c.IndentedJSON(http.StatusInternalServerError, error)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": msg})
}

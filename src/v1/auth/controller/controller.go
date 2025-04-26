package controller

import (
	"go-simple-rest/db"
	"go-simple-rest/src/v1/auth/implementation"
	"go-simple-rest/src/v1/auth/model"
	repo "go-simple-rest/src/v1/auth/repo/implementation"
	"go-simple-rest/src/v1/jwt"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var client, _, _ = db.Connect()

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

	payload := c.MustGet("body").(model.LoginAuth)
	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	response, ok := s.Login(c, payload)

	if ok {
		c.IndentedJSON(http.StatusInternalServerError, response.MESSAGE)
		return
	}

	data := response.DATA.(map[string]interface{})
	userData := data["USER"].(model.UserResponseObject)

	c.SetCookie("token", data["TOKEN"].(string), 3600, "/", "127.0.0.1", false, true)
	c.SetSameSite(http.SameSiteStrictMode)
	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusOK, Success: true, Message: response.MESSAGE, Data: gin.H{
		"username": userData.USERNAME, "role": jwt.GetRole("user"),
		"id": userData.ID,
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
	payload := c.MustGet("body").(model.RegisterAuth)
	rep, _ := repo.New(database)
	s := implementation.NewService(rep)

	response, ok := s.Register(c, payload)

	if ok {
		c.IndentedJSON(http.StatusInternalServerError, response.MESSAGE)
		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusCreated, Success: true, Message: response.MESSAGE, Data: nil})
}

package controller

import (
	"go-simple-rest/src/v1/auth/model"
	"go-simple-rest/src/v1/auth/repo"
	"go-simple-rest/src/v1/auth/service"
	"go-simple-rest/src/v1/jwt"
	"go-simple-rest/src/v1/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	database *mongo.Database
}

func AuthControllerImpl(database *mongo.Database) *AuthController {
	return &AuthController{database: database}
}

func (ac *AuthController) Login(c *gin.Context) {

	payload := c.MustGet("body").(model.LoginAuth)
	rep, _ := repo.New(ac.database)
	s := service.NewService(rep)

	response, ok := s.Login(c, payload)

	if ok {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: response.MESSAGE, Data: nil})

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

func (ac *AuthController) Register(c *gin.Context) {
	payload := c.MustGet("body").(model.RegisterAuth)
	rep, _ := repo.New(ac.database)
	s := service.NewService(rep)

	response, ok := s.Register(c, payload)

	if ok {
		utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusInternalServerError, Success: false, Message: response.MESSAGE, Data: nil})
		return
	}

	utils.TransformResponse(c, utils.Reponse{StatusCode: http.StatusCreated, Success: true, Message: response.MESSAGE, Data: nil})
}

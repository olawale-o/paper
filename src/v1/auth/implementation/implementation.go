package implementation

import (
	"context"
	"go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/jwt"

	"go-simple-rest/src/v1/auth/model"
	"go-simple-rest/src/v1/auth/repo"
	authSvc "go-simple-rest/src/v1/auth/service"

	"github.com/gin-gonic/gin"
)

type service struct {
	repo repo.Repository
}

func NewService(rep repo.Repository) authSvc.Service {
	return &service{repo: rep}
}

func (s *service) Login(ctx *gin.Context, payload model.LoginAuth) (model.AuthResponse, bool) {

	dbUser, err := s.repo.FindOne(context.TODO(), "users", payload.USERNAME)

	if err != nil {
		return model.AuthResponse{MESSAGE: "Cannot find user"}, true
	}

	match := auth.CheckPasswordHash(payload.PASSWORD, dbUser.PASSWORD)
	if !match {
		return model.AuthResponse{MESSAGE: "Please provide valid password credentials"}, true
	}

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.ID, "role": "user", "id": dbUser.ID})

	if err != nil {
		return model.AuthResponse{MESSAGE: "Internal Server error"}, true
	}

	response := model.AuthResponse{
		MESSAGE: "Logged in Succesfully",
		DATA: map[string]interface{}{
			"TOKEN": tokenString,
			"USER": model.UserResponseObject{
				USERNAME: dbUser.USERNAME,
				ID:       dbUser.ID,
				ROLE:     dbUser.ROLE,
			},
		},
	}

	return response, false
}

func (s *service) Register(ctx *gin.Context, payload model.RegisterAuth) (model.AuthResponse, bool) {

	dbUser, _ := s.repo.FindOne(context.TODO(), "users", payload.USERNAME)

	if dbUser.USERNAME != "" {
		return model.AuthResponse{MESSAGE: "Kindly login with your credentials"}, true
	}

	hash, _ := auth.HashPassword(payload.PASSWORD)

	_, err := s.repo.InsertOne(context.TODO(), "users", model.User{
		USERNAME:  payload.USERNAME,
		FIRSTNAME: payload.FIRSTNAME,
		LASTNAME:  payload.LASTNAME,
		PASSWORD:  hash,
	})
	if err != nil {
		return model.AuthResponse{MESSAGE: "Something went wrong"}, true
	}
	return model.AuthResponse{MESSAGE: "User created successfully"}, false
}

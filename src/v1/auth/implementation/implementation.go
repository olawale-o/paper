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

func (s *service) Login(ctx *gin.Context, payload model.LoginAuth) (model.LoginResponse, gin.H) {
	var response model.LoginResponse

	dbUser, err := s.repo.FindOne(context.TODO(), "users", payload.USERNAME)

	if err != nil {
		return response, gin.H{"err": "Cannot find user"}
	}

	match := auth.CheckPasswordHash(payload.PASSWORD, dbUser.PASSWORD)
	if !match {
		return response, gin.H{"err": "Please provide valid password credentials"}
	}

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.ID, "role": "user", "id": dbUser.ID})

	if err != nil {
		return response, gin.H{"err": "Internal Server error"}
	}

	response = model.LoginResponse{
		MESSAGE: "Logged in Succesfully",
		TOKEN:   tokenString,
		USER: model.UserResponseObject{
			USERNAME: dbUser.USERNAME,
			ID:       dbUser.ID,
			ROLE:     dbUser.ROLE,
		},
	}

	return response, gin.H{}
}

func (s *service) Register(ctx *gin.Context, payload model.RegisterAuth) (string, gin.H) {

	dbUser, _ := s.repo.FindOne(context.TODO(), "users", payload.USERNAME)

	if dbUser.USERNAME != "" {
		return "", gin.H{"err": "Kindly login with your credentials"}
	}

	hash, _ := auth.HashPassword(payload.PASSWORD)

	_, err := s.repo.InsertOne(context.TODO(), "users", model.User{
		USERNAME:  payload.USERNAME,
		FIRSTNAME: payload.FIRSTNAME,
		LASTNAME:  payload.LASTNAME,
		PASSWORD:  hash,
	})
	if err != nil {
		return "", gin.H{"err": "Something went wrong"}
	}
	return "User created", gin.H{}
}

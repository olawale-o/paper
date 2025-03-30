package implementation

import (
	"auth/bcrypt"
	"auth/jwt"
	"auth/model"
	"auth/translator"
	"context"

	authSvc "auth/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repo model.Repository
}

func NewService(rep model.Repository) authSvc.Service {
	return &service{repo: rep}
}

func (s *service) Login(ctx *gin.Context, payload model.LoginAuth) (interface{}, gin.H) {
	validate := validator.New()
	err := validate.Struct(payload)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {
		return "", gin.H{"err": errs}
	}

	dbUser, err := s.repo.GetUser(context.TODO(), "users", payload.USERNAME)

	if err != nil {
		return model.LoginResponse{}, gin.H{"err": "Cannot find user"}
	}

	match := bcrypt.CheckPasswordHash(payload.PASSWORD, dbUser.PASSWORD)
	if !match {
		return "", gin.H{"err": "Please provide valid password credentials"}
	}

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.ID, "role": "user", "id": dbUser.ID})

	if err != nil {
		return "", gin.H{"err": "Internal Server error"}
	}

	return model.LoginResponse{TOKEN: tokenString, USER: model.UserResponseObject{
		USERNAME: dbUser.USERNAME,
		ID:       dbUser.ID,
		ROLE:     dbUser.ROLE,
	}}, gin.H{}
}

func (s *service) Register(ctx *gin.Context, payload model.RegisterAuth) (string, gin.H) {

	validate := validator.New()
	err := validate.Struct(payload)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {

		return "", gin.H{"err": errs}
	}

	dbUser, err := s.repo.GetUser(context.TODO(), "users", payload.USERNAME)

	if dbUser.USERNAME != "" {

		return "", gin.H{"err": "Unable to create user"}
	}
	hash, _ := bcrypt.HashPassword(payload.PASSWORD)

	_, err = s.repo.InsertUser(context.TODO(), "users", model.User{
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

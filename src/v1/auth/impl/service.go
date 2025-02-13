package impl

import (
	"context"
	"go-simple-rest/src/v1/auth"
	authsvc "go-simple-rest/src/v1/auth"
	"go-simple-rest/src/v1/auth/bcrypt"
	"go-simple-rest/src/v1/authors"
	"go-simple-rest/src/v1/jwt"
	"go-simple-rest/src/v1/translator"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/go-playground/validator/v10"
)

type service struct {
	repo   authsvc.Repository
	logger log.Logger
}

func NewService(rep authsvc.Repository, logger log.Logger) auth.Service {
	return &service{repo: rep, logger: logger}
}

func (s *service) Login(ctx *gin.Context, payload auth.LoginAuth) (string, gin.H) {
	// logger := log.With(s.logger, "method", "Create")

	validate := validator.New()
	err := validate.Struct(payload)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {
		return "", gin.H{"err": errs}
	}

	dbUser, err := s.repo.GetUser(context.TODO(), "users", payload.USERNAME)

	if err != nil {
		return "", gin.H{"err": "Cannot find user"}
	}

	match := bcrypt.CheckPasswordHash(payload.PASSWORD, dbUser.PASSWORD)
	if !match {

		return "", gin.H{"err": "Please provide valid password credentials"}
	}

	tokenString, err := jwt.CreateToken(map[string]string{"username": dbUser.ID, "role": "user", "id": dbUser.ID})

	if err != nil {
		return "", gin.H{"err": "Internal Server error"}
	}

	return tokenString, gin.H{}
}
func (s *service) Register(ctx *gin.Context, payload auth.RegisterAuth) (string, gin.H) {
	// logger := log.With(s.logger, "method", "Create")

	validate := validator.New()
	err := validate.Struct(payload)
	errs := translator.Translate(validate, err)

	if len(errs) > 0 {
		return "", gin.H{"err": errs}
	}

	dbUser, err := s.repo.GetUser(context.TODO(), "users", payload.USERNAME)

	if dbUser.USERNAME != "" {
		return "", gin.H{"err": "Unable to create user. User already exists"}

	}
	hash, _ := bcrypt.HashPassword(payload.PASSWORD)
	payload.PASSWORD = hash
	_, err = s.repo.InsertUser(context.TODO(), "users", authors.User{
		FIRSTNAME: payload.FIRSTNAME,
		LASTNAME:  payload.LASTNAME,
		PASSWORD:  hash,
		USERNAME:  payload.USERNAME,
	})

	if err != nil {
		return "", gin.H{"err": "Cannot create user"}
	}
	return "User created", gin.H{}
}

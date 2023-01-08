// Package usecase is described all business logic in app
package usecase

import (
	"context"
	"os"
	"time"

	"github.com/MochamadAkbar/ordent-test/api"
	"github.com/MochamadAkbar/ordent-test/common/constants"
	commonErr "github.com/MochamadAkbar/ordent-test/common/errors"
	commonJwt "github.com/MochamadAkbar/ordent-test/common/jwt"
	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(ctx context.Context, user *entity.User) (api.UserResponse, error)
	Login(ctx context.Context, user *entity.User) (api.UserResponse, error)
}

type UserUsecaseImpl struct {
	Repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		Repository: repository,
	}
}

func (usecase *UserUsecaseImpl) Register(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse

	expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), constants.Salt)
	if err != nil {
		return resp, commonErr.ErrInternalServer
	}

	user.Password = string(hash)

	result, ok := usecase.Repository.Register(ctx, user)

	if !ok {
		return resp, commonErr.ErrInternalServer
	}

	token, err := commonJwt.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":   result.ID,
			"expiresIn": expiresIn,
		},
		[]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		panic(err)
	}

	resp.UserID = result.ID
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}

func (usecase *UserUsecaseImpl) Login(ctx context.Context, user *entity.User) (api.UserResponse, error) {
	var resp api.UserResponse

	expiresIn := time.Now().Add(time.Duration(1) * time.Minute).Unix()

	result, ok := usecase.Repository.Login(ctx, user)
	if !ok {
		return api.UserResponse{}, commonErr.ErrNotFound
	}

	errCreds := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if errCreds != nil {
		return resp, commonErr.ErrUnauthorized
	}

	token, err := commonJwt.JwtClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id":    result.ID,
			"expires_in": expiresIn,
		},
		[]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		panic(err)
	}
	resp.UserID = result.ID
	resp.Token = token
	resp.ExpiresIn = expiresIn

	return resp, nil
}

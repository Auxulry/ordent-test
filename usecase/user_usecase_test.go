package usecase

import (
	"context"
	"testing"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userUsecase        = UserUsecaseImpl{Repository: userRepositoryMock}
	userRepositoryMock = &repository.UserRepositoryMock{Mock: mock.Mock{}}
)

func TestUserUsecaseImpl_Register(t *testing.T) {
	t.Run("Test User Use Case Register Success", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "johndoe@gmail.com",
			Password: "Password",
		}

		ctx := context.Background()

		userRepositoryMock.Mock.On("Register", &user).Return(user, true)

		result, err := userUsecase.Register(ctx, &user)

		assert.Nil(t, err)
		assert.Equal(t, 0, result.UserID)
		assert.NotEqual(t, "", result.Token)
		assert.NotEqual(t, int64(0), result.ExpiresIn)
	})

	t.Run("Test User Use Case Register Fail", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "johndoe@gmail.com",
			Password: "Password",
		}

		ctx := context.Background()

		userRepositoryMock.Mock.On("Register", &user).Return(nil, false)

		result, err := userUsecase.Register(ctx, &user)

		assert.NotNil(t, err)
		assert.Equal(t, "internal server error", err.Error())
		assert.Equal(t, 0, result.UserID)
		assert.Equal(t, "", result.Token)
		assert.Equal(t, int64(0), result.ExpiresIn)
	})
}

func TestUserUsecase_Login(t *testing.T) {
	t.Run("Test User Use Case Login Success", func(t *testing.T) {
		user := entity.User{
			Email:    "johndoe@gmail.com",
			Password: "Password",
		}

		ctx := context.Background()

		userRepositoryMock.Mock.On("Login", &user).Return(user, true)

		result, err := userUsecase.Login(ctx, &user)

		assert.Nil(t, err)
		assert.Equal(t, 0, result.UserID)
		assert.NotEqual(t, "", result.Token)
		assert.NotEqual(t, 0, result.ExpiresIn)
	})

	t.Run("Test User Use Case Login Failed", func(t *testing.T) {
		user := entity.User{
			Name:     "John Doe",
			Email:    "johndoe@gmail.com",
			Password: "Password",
		}

		ctx := context.Background()

		userRepositoryMock.Mock.On("Login", &user).Return(nil, false)

		result, err := userUsecase.Login(ctx, &user)

		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Equal(t, 0, result.UserID)
		assert.Equal(t, "", result.Token)
		assert.Equal(t, int64(0), result.ExpiresIn)
	})
}

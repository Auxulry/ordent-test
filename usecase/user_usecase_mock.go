package usecase

import (
	"context"
	"errors"

	"github.com/MochamadAkbar/ordent-test/api"
	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/stretchr/testify/mock"
)

type UserUsecaseMock struct {
	Mock       mock.Mock
	Repository repository.UserRepositoryMock
}

func (usecase *UserUsecaseMock) Register(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)
	if args.Get(0) == nil {
		return api.UserResponse{}, errors.New("internal server error")
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}

func (usecase *UserUsecaseMock) Login(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)

	if args.Get(0) == nil {
		return api.UserResponse{}, errors.New("user not found")
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}

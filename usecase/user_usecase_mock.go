package usecase

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/api"
	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	Mock       mock.Mock
	Repository repository.UserRepositoryMock
}

func (usecase *UserUseCaseMock) Register(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)
	if args.Get(0) == nil {
		return api.UserResponse{}, args.Error(1)
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}

func (usecase *UserUseCaseMock) Login(_ context.Context, user *entity.User) (api.UserResponse, error) {
	args := usecase.Mock.Called(user)

	if args.Get(0) == nil {
		return api.UserResponse{}, args.Error(1)
	} else {
		result := args.Get(0).(api.UserResponse)
		return result, nil
	}
}

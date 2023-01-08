package repository

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (repository *UserRepositoryMock) Register(_ context.Context, user *entity.User) (entity.User, bool) {
	args := repository.Mock.Called(user)

	if args.Get(0) == nil {
		return entity.User{}, false
	} else {
		result := args.Get(0).(entity.User)
		return result, true
	}
}

func (repository *UserRepositoryMock) Login(_ context.Context, user *entity.User) (entity.User, bool) {
	args := repository.Mock.Called(user)

	if args.Get(0) == nil {
		return entity.User{}, false
	} else {
		result := args.Get(0).(entity.User)
		return result, true
	}
}

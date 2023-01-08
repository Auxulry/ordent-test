//go:build wireinject
// +build wireinject

// Package injector describe all dependency injectors
package injector

import (
	"github.com/MochamadAkbar/ordent-test/handler"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/MochamadAkbar/ordent-test/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeUserService(conn *pgxpool.Pool, router chi.Router) error {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
	)
	return nil
}

// Package repository is describe all action to database
package repository

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) bool
	Login(ctx context.Context, user *entity.User) (entity.User, bool)
}

type UserRepositoryImpl struct {
	Conn *pgxpool.Pool
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, user *entity.User) bool {
	statement := `INSERT INTO "users" ("name", "email", "password") VALUES($1, $2, $3);`
	_, err := repository.Conn.Exec(ctx, statement, user.Name, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	return true
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user *entity.User) (entity.User, bool) {
	statement := `SELECT "id", "email", "password" FROM "users" WHERE "email" = $1;`
	var result entity.User
	err := repository.Conn.QueryRow(ctx, statement, user.Email).
		Scan(&result.ID, &result.Email, &result.Password)

	if err != nil {
		panic(err.Error())
	}

	return result, true
}

package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func NewDBTest(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)

	return pool, err
}

func TestConnectionSuccess(t *testing.T) {
	t.Run("Test Scenario DB Connection Success", func(t *testing.T) {
		ctx := context.Background()
		urlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"))

		pool, err := NewDBTest(ctx, urlString)
		defer pool.Close()

		err = pool.Ping(ctx)

		assert.Nil(t, err)
		assert.NotNil(t, pool)
	})
}

func TestConnectionFailed(t *testing.T) {
	t.Run("Test Scenario DB Connection Failed", func(t *testing.T) {
		ctx := context.Background()
		urlString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			"must_failed",
			"must_failed",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			"123")

		pool, err := NewDBTest(ctx, urlString)
		defer pool.Close()

		err = pool.Ping(ctx)

		assert.NotNil(t, err)
		assert.NotNil(t, pool)
	})
}

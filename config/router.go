// Package config is describe all configuration
package config

import (
	"github.com/MochamadAkbar/ordent-test/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recovery)
	router.Use(middleware.CORS)
	return router
}

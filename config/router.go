// Package config is describe all configuration
package config

import "github.com/go-chi/chi/v5"

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	return router
}

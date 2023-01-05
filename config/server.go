package config

import (
	"net/http"
	"os"

	"github.com/MochamadAkbar/ordent-test/common/constants"
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT"),
		Handler:        handler,
		ReadTimeout:    constants.ReadTimeout,
		WriteTimeout:   constants.WriteTimeout,
		MaxHeaderBytes: constants.MaxHeaderBytes,
	}
}

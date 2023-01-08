// Package middleware is describe all chi middleware
package middleware

import (
	"net/http"

	"github.com/MochamadAkbar/ordent-test/api"
	"github.com/MochamadAkbar/ordent-test/common/serializer"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if errRecover := recover(); errRecover != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				errResponse := api.ErrResponse{
					Code:    http.StatusInternalServerError,
					Status:  http.StatusText(http.StatusInternalServerError),
					Message: "There was an internal server error",
				}

				_ = serializer.SerializeWriter(w, http.StatusInternalServerError, errResponse)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

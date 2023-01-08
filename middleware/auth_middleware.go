package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/MochamadAkbar/ordent-test/api"
	commonErr "github.com/MochamadAkbar/ordent-test/common/errors"
	commonJwt "github.com/MochamadAkbar/ordent-test/common/jwt"
	"github.com/MochamadAkbar/ordent-test/common/serializer"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != "" {
			token := strings.Split(header, "Bearer ")
			_, err := commonJwt.JwtValidate(commonJwt.HMAC, token[1], []byte(os.Getenv("SECRET_KEY")))
			if err != nil {
				res := api.ErrResponse{
					Code:    commonErr.GetHTTPErrorCode(err),
					Status:  http.StatusText(commonErr.GetHTTPErrorCode(err)),
					Message: err.Error(),
				}
				_ = serializer.SerializeWriter(w, res.Code, res)
				return
			}
			next.ServeHTTP(w, r)
		} else {
			res := api.ErrResponse{
				Code:    commonErr.GetHTTPErrorCode(commonErr.ErrUnauthorized),
				Status:  http.StatusText(commonErr.GetHTTPErrorCode(commonErr.ErrUnauthorized)),
				Message: commonErr.ErrUnauthorized.Error(),
			}
			_ = serializer.SerializeWriter(w, res.Code, res)
			return
		}
	})
}

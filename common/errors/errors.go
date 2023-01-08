// Package errors is utils to catch error from use case
package errors

import (
	"errors"
	"net/http"

	"github.com/MochamadAkbar/ordent-test/common/constants"
)

var (
	ErrNotFound       error = errors.New("NOT_FOUND")
	ErrInternalServer       = errors.New("INTERNAL_SERVER_ERROR")
)

func GetHTTPErrorCode(err error) int {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	} else if errors.Is(err, ErrInternalServer) {
		return http.StatusInternalServerError
	} else {
		return constants.DefaultErrCode
	}
}

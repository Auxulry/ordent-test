// Package serializer is describe json serializer
package serializer

import (
	"encoding/json"
	"net/http"
)

func SerializeRequest[T any](r *http.Request, result T) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func SerializeWriter[T any](w http.ResponseWriter, code int, response T) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		return err
	}
	return nil
}

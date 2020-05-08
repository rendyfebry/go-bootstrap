package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/rendyfebry/go-streamer/service"
)

// HTTPError ...
type HTTPError struct {
	Code    string
	Message string
}

func MakeHTTPRoutes(svc service.SomeService) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", MakeIndexHandler(svc))
	r.HandleFunc("/health", MakeHealthHandler(svc))

	return r
}

// MakeIndexHandler ...
func MakeIndexHandler(svc service.SomeService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := svc.GetIndex()

		w.Write([]byte(data))
	})
}

// MakeHealthHandler ...
func MakeHealthHandler(svc service.SomeService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := svc.GetHealth()

		w.Write([]byte(data))
	})
}

// MakeNotFoundHandler ...
func MakeNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.
			NewEncoder(w).
			Encode(
				HTTPError{
					Code:    string(http.StatusNotFound),
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}

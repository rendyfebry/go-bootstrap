package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPError ...
type HTTPError struct {
	Code    string
	Message string
}

func MakeHTTPRoutes() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/", MakeIndexHandler())
	r.HandleFunc("/health", MakeHealthHandler())

	return r
}

// MakeIndexHandler ...
func MakeIndexHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Index route"))
	})
}

// MakeHealthHandler ...
func MakeHealthHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Health route"))
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

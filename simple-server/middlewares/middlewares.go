package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/* Middlewares */
func InjectMiddlewares(h *mux.Router) *mux.Router {
	h.Use(setHeaders)
	h.Use(authenticate)

	return h
}

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bearer string = r.Header.Get("Authorization")
		if bearer == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("You must authenticate")
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func setHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

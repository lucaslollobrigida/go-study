package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r = InjectRoutes(r)
	r.Use(ParseResponse)
	r.Use(Authenticate)

	http.ListenAndServe(":8081", r)
}

/* Routes */
func InjectRoutes(h *mux.Router) *mux.Router {
	h.HandleFunc("/hi", HandleHi)
	h.HandleFunc("/check", HandleCheck)

	return h
}

func HandleHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello")
}

func HandleCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("All OK")
}

/* Middlewares */
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bearer string = r.Header.Get("Authorization")
		if bearer == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("You must authenticate")
		} else {
			fmt.Print("Received %s", bearer)
			next.ServeHTTP(w, r)
		}
	})
}

func ParseResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

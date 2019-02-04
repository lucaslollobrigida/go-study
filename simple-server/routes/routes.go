package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/* Routes */
func InjectRoutes(h *mux.Router) *mux.Router {
	h.HandleFunc("/hi", handleHi)
	h.HandleFunc("/check", handleCheck)

	return h
}

func handleHi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hello")
}

func handleCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("All OK")
}

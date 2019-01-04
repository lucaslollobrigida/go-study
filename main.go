package main

import (
	"encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/hello/{name}", SayHello).Methods("GET")
		router.HandleFunc("/", Welcome).Methods("GET")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func SayHello(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		message := "Hello, " + params["name"]
		json.NewEncoder(w).Encode(message)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Welcome")
}

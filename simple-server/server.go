// Package main provides a simple web server example
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaslollobrigida/go-study/simple-server/middlewares"
	"github.com/lucaslollobrigida/go-study/simple-server/routes"
)

func main() {
	r := mux.NewRouter()

	r = routes.InjectRoutes(r)
	r = middlewares.InjectMiddlewares(r)

	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}

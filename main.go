package main

import (
	"go-simple-rest/src/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	routes.Routes(r)
	log.Println("Starting server... on port ", 7000)
	http.ListenAndServe(":7000", r) // Gateway listens on port 8000
}

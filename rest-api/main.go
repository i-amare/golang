package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/i-amare/rest-api/db"
	"github.com/i-amare/rest-api/routes"
)

func main() {
	db.InitDB()

	router := chi.NewRouter()
	routes.InitRoutes(router)

	http.ListenAndServe(":3000", router)
}

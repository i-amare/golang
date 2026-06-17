package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/i-amare/rest-api/middleware"
	"github.com/i-amare/rest-api/utils"
)

func InitRoutes(router chi.Router) {
	router.Get("/", ping)

	loadUserRoutes(router)
	loadVendorRoutes(router)

	router.Route("/vendor", loadVendorRoutes)
}

func loadUserRoutes(router chi.Router) {
	router.Post("/signup", createUser)
	router.Post("/login", loginUser)
	router.Get("/users", getAllUsers)
}

func loadVendorRoutes(router chi.Router) {
	router.Get("/", getAllVendors)
	router.Get("/{id}", getVendor)

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.Authenticate)
		r.Post("/", createVendor)
		r.Put("/{id}", updateVendor)
		r.Delete("/{id}", deleteVendor)
		r.Post("/menu", createMenuItem)
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccess(w, http.StatusOK, "Hello World", map[string]int{
		"res": 200,
	})
}

package routes

import (
	"GO-JWT/controllers"
	"GO-JWT/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r * mux.Router){
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middleware.Auth)
	
	router.HandleFunc("/me", controllers.Me).Methods("GET")
}
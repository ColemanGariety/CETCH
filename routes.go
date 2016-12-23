package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/JacksonGariety/wetch/controllers"
)

func NewRouter() *mux.Router {
	// The router
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// RESTful Routes
	router.Methods("Get").Path("/").HandlerFunc(controllers.Index)
	router.Methods("Get").Path("/login").HandlerFunc(controllers.LoginShow)

	return router
}

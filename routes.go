package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/JacksonGariety/wetch/controllers"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("session-secret"))

func NewRouter() *mux.Router {
	// The router
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// RESTful Routes
	router.Methods("Get").Path("/").HandlerFunc(controllers.Index)
	router.Methods("Get").Path("/login").HandlerFunc(controllers.LoginShow)
	router.Methods("Post").Path("/login").HandlerFunc(controllers.LoginPost)
	router.Methods("Get").Path("/signup").HandlerFunc(controllers.SignupShow)
	router.Methods("Post").Path("/signup").HandlerFunc(controllers.SignupPost)

	return router
}

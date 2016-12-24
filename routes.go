package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/JacksonGariety/wetch/controllers"
	"github.com/JacksonGariety/wetch/utils"
)

func NewRouter() *mux.Router {
	// The router
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// WTful Routes
	router.Methods("Get").Path("/").HandlerFunc(controllers.Index)

	// login
	router.Methods("Get").Path("/login").HandlerFunc(controllers.LoginShow)
	router.Methods("Post").Path("/login").HandlerFunc(controllers.LoginPost)
	router.Methods("Get").Path("/logout").HandlerFunc(utils.AuthorizeClaims(controllers.LogoutShow))

	// signup
	router.Methods("Get").Path("/signup").HandlerFunc(controllers.SignupShow)
	router.Methods("Post").Path("/signup").HandlerFunc(controllers.SignupPost)

	router.Methods("Get", "Post").Path("/profile").HandlerFunc(utils.AuthorizeClaims(controllers.ProfileShow))

	return router
}

package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/justinas/alice"
	"github.com/NYTimes/gziphandler"

	"github.com/JacksonGariety/cetch/app/controllers"
	"github.com/JacksonGariety/cetch/app/middleware"
)

func NewRouter() *mux.Router {
	// The router
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Middleware
	chain := alice.New(
		middleware.Timeout,
		middleware.Authorize,
		gziphandler.GzipHandler,
	)

	// Routes
	router.Methods("Get").Path("/").Handler(chain.ThenFunc(controllers.Index))

	// login
	router.Methods("Get").Path("/login").Handler(chain.ThenFunc(controllers.LoginShow))
	router.Methods("Post").Path("/login").Handler(chain.ThenFunc(controllers.LoginPost))
	router.Methods("Get").Path("/logout").Handler(chain.ThenFunc(controllers.LogoutShow))

	// signup
	router.Methods("Get").Path("/signup").Handler(chain.ThenFunc(controllers.SignupShow))
	router.Methods("Post").Path("/signup").Handler(chain.ThenFunc(controllers.SignupPost))

	// profile
	router.Methods("Get", "Post").Path("/profile").Handler(chain.ThenFunc(controllers.ProfileShow))

	return router
}

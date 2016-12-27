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

	// Index
	router.Methods("Get").Path("/").Handler(chain.ThenFunc(controllers.Index))

	// login/logout
	router.Methods("Get").Path("/login").Handler(chain.ThenFunc(controllers.LoginShow))
	router.Methods("Post").Path("/login").Handler(chain.ThenFunc(controllers.LoginPost))
	router.Methods("Get").Path("/logout").Handler(chain.ThenFunc(controllers.LogoutShow))

	// forgotten/reset
	router.Methods("Get").Path("/forgotten").Handler(chain.ThenFunc(controllers.ForgottenShow))
	router.Methods("Post").Path("/forgotten").Handler(chain.ThenFunc(controllers.ForgottenPost))

	// signup
	router.Methods("Get").Path("/signup").Handler(chain.ThenFunc(controllers.SignupShow))
	router.Methods("Post").Path("/signup").Handler(chain.ThenFunc(controllers.SignupPost))

	// profile
	router.Methods("Get", "Post").Path("/user/{name}").Handler(chain.ThenFunc(controllers.UserShow))

	// competitions
	router.Methods("Get").Path("/competition/new").Handler(chain.Append(middleware.Protect).ThenFunc(controllers.CompetitionNew))
	router.Methods("Get").Path("/competition/{id}").Handler(chain.ThenFunc(controllers.CompetitionShow))
	router.Methods("Post").Path("/competition/new").Handler(chain.Append(moddleare.Protect).ThenFunc(controllers.CompetitionCreate))
	router.Methods("Get", "Post").Path("/competitions").Handler(chain.ThenFunc(controllers.CompetitionsShow))

	return router
}

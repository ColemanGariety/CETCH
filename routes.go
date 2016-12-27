package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/justinas/alice"
	"github.com/NYTimes/gziphandler"

	c "github.com/JacksonGariety/cetch/app/controllers"
	m "github.com/JacksonGariety/cetch/app/middleware"
)

func NewRouter() *mux.Router {
	// The router
	router := mux.NewRouter()

	// Static files
	fs := http.FileServer(http.Dir("static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// Middleware
	chain := alice.New(
		m.Timeout,
		gziphandler.GzipHandler,
		m.Authenticate,
	)

	// Index
	router.Methods("Get").Path("/").Handler(chain.ThenFunc(c.Index))

	// login/logout
	router.Methods("Get").Path("/login").Handler(chain.Append(m.Retain).ThenFunc(c.LoginShow))
	router.Methods("Post").Path("/login").Handler(chain.Append(m.Retain).ThenFunc(c.LoginPost))
	router.Methods("Get").Path("/logout").Handler(chain.ThenFunc(c.LogoutShow))

	// forgotten/reset
	router.Methods("Get").Path("/forgotten").Handler(chain.Append(m.Retain).ThenFunc(c.ForgottenShow))
	router.Methods("Post").Path("/forgotten").Handler(chain.Append(m.Retain).ThenFunc(c.ForgottenPost))

	// signup
	router.Methods("Get").Path("/signup").Handler(chain.Append(m.Retain).ThenFunc(c.SignupShow))
	router.Methods("Post").Path("/signup").Handler(chain.Append(m.Retain).ThenFunc(c.SignupPost))

	// profile
	router.Methods("Get", "Post").Path("/user/{name}").Handler(chain.ThenFunc(c.UserShow))

	// competitions
	router.Methods("Get").Path("/competition/new").Handler(chain.Append(m.Forbid).ThenFunc(c.CompetitionNew))
	router.Methods("Get").Path("/competition/{id}").Handler(chain.ThenFunc(c.CompetitionShow))
	router.Methods("Post").Path("/competition/new").Handler(chain.Append(m.Forbid).ThenFunc(c.CompetitionCreate))
	router.Methods("Get", "Post").Path("/competitions").Handler(chain.ThenFunc(c.CompetitionsShow))

	return router
}

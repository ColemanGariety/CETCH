package main

import (
	"net/http"
	"github.com/go-zoo/bone"

	"github.com/justinas/alice"
	"github.com/NYTimes/gziphandler"

	c "github.com/JacksonGariety/cetch/app/controllers"
	m "github.com/JacksonGariety/cetch/app/middleware"
)

func NewRouter() http.Handler {
	mux := bone.New()

	// Middleware
	chain := alice.New(
		m.Timeout,
		gziphandler.GzipHandler,
		m.Authenticate,
	)

	mux.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Get("/", chain.ThenFunc(c.Index))
	mux.Get("/login", chain.Append(m.Retain).ThenFunc(c.LoginShow))
	mux.Post("/login", chain.Append(m.Retain).ThenFunc(c.LoginPost))
	mux.Get("/logout", chain.ThenFunc(c.LogoutShow))
	mux.Get("/forgotten", chain.Append(m.Retain).ThenFunc(c.ForgottenShow))
	mux.Post("/forgotten", chain.Append(m.Retain).ThenFunc(c.ForgottenPost))
	mux.Get("/signup", chain.Append(m.Retain).ThenFunc(c.SignupShow))
	mux.Post("/signup", chain.Append(m.Retain).ThenFunc(c.SignupPost))
	mux.Get("/user/:name", chain.ThenFunc(c.UserShow))
	mux.Post("/user/:name", chain.ThenFunc(c.UserShow))
	mux.Get("/competition/new", chain.Append(m.Forbid).ThenFunc(c.CompetitionNew))
	mux.Get("/competition/:id", chain.ThenFunc(c.CompetitionShow))
	mux.Post("/competition/new", chain.Append(m.Forbid).ThenFunc(c.CompetitionCreate))
	mux.Get("/competitions", chain.ThenFunc(c.CompetitionsShow))
	mux.Post("/competitions", chain.ThenFunc(c.CompetitionsShow))

	return mux
}

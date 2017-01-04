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
		m.StickyEntry,
	)

	mux.Get("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Get("/", chain.ThenFunc(c.Index))
	mux.Get("/rules", chain.ThenFunc(c.Rules))
	mux.Get("/login", chain.Append(m.Retain).ThenFunc(c.LoginShow))
	mux.Post("/login", chain.Append(m.Retain).ThenFunc(c.Login))
	mux.Get("/logout", chain.ThenFunc(c.LogoutShow))
	mux.Get("/forgotten", chain.Append(m.Retain).ThenFunc(c.ForgottenShow))
	mux.Post("/forgotten", chain.Append(m.Retain).ThenFunc(c.ForgottenSendEmail))
	mux.Get("/signup", chain.Append(m.Retain).ThenFunc(c.SignupShow))
	mux.Post("/signup", chain.Append(m.Retain).ThenFunc(c.SignupPost))
	mux.Get("/user/:name", chain.ThenFunc(c.UserShow))
	mux.Post("/user/:name", chain.ThenFunc(c.UserShow))
	mux.Get("/competitors", chain.ThenFunc(c.UsersShow))
	mux.Get("/competition/new", chain.Append(m.Forbid).ThenFunc(c.CompetitionNew))
	mux.Post("/competition/new", chain.Append(m.Forbid).ThenFunc(c.CompetitionCreate))
	mux.Get("/competition/:id", chain.ThenFunc(c.CompetitionShow))
	mux.Get("/competition/:id/edit", chain.Append(m.Forbid).ThenFunc(c.CompetitionEdit))
	mux.Post("/competition/:id/edit", chain.Append(m.Forbid).ThenFunc(c.CompetitionUpdate))
	mux.Post("/archive", chain.Append(m.Forbid).ThenFunc(c.Archive))
	mux.Get("/archive", chain.ThenFunc(c.Archive))
	mux.Get("/schedule", chain.Append(m.Forbid).ThenFunc(c.ScheduleShow))
	mux.Post("/schedule", chain.Append(m.Forbid).ThenFunc(c.SchedulePost))
	mux.Get("/entry/new", chain.ThenFunc(c.EntryNew))
	mux.Post("/entry/new", chain.Append(m.Protect).ThenFunc(c.EntryCreate))
	mux.Get("/entry/:id", chain.Append(m.Protect).ThenFunc(c.EntryShow))
	mux.Post("/entry/:id", chain.Append(m.Protect).ThenFunc(c.EntryShow))

	return mux
}

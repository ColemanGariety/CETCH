package controllers

import (
	"net/http"
	"html/template"
)

func LoginShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/login.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// validation here
}

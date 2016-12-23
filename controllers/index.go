package controllers

import (
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/index.html")
	t.ExecuteTemplate(w, "layout", nil)
}

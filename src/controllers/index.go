package controllers

import (
	"path"
	"net/http"
	"html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	lp := path.Join("templates", "layout.html")
	fp := path.Join("templates", "index.html")

	t, _ := template.ParseFiles(lp, fp)
	t.ExecuteTemplate(w, "layout", nil)
}

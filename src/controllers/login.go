package controllers

import (
	"path"
	"net/http"
	"html/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		lp := path.Join("./templates", "layout.html")
		fp := path.Join("./templates", "login.html")

		t, _ := template.ParseFiles(lp, fp)
		t.ExecuteTemplate(w, "layout", nil)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

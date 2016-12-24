package utils

import (
	"net/http"
	"html/template"
	"path"
)

func Render(w http.ResponseWriter, filename string, data interface{}) {
  tmpl, err := template.ParseFiles("views/layout.html", path.Join("views", filename))
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

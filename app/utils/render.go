package utils

import (
	"net/http"
	"html/template"
	"path"
)

type Props map[string]interface{}

func Render(w http.ResponseWriter, filename string, props interface{}) {
	tmpl := template.Must(template.New("base").ParseFiles("app/views/layout.html", path.Join("app/views", filename)))

  if err := tmpl.ExecuteTemplate(w, "layout", props); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

package utils

import (
	"net/http"
	"html/template"
	"path"
	"os"
)

type Props map[string]interface{}

var basePath = os.Getenv("base_path")

func Render(w http.ResponseWriter, filename string, props interface{}) {
	tmpl := template.Must(template.New("base").ParseFiles(path.Join(basePath, "./app/views/layout.html"), path.Join(basePath, "app/views", filename)))

  if err := tmpl.ExecuteTemplate(w, "layout", props); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

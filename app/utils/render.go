package utils

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"time"
	"strconv"
	"path/filepath"
	"log"
)

var BasePath = os.Getenv("base_path")

func formatDate(date time.Time) string {
	return date.Format("January 2")
}

func formatDateForForm(date time.Time) string {
	return date.Format("2006-01-02")
}

func formatSolution(solution float64) string {
	return strconv.FormatFloat(solution, 'f', 6, 64)
}

var funcMap = template.FuncMap{
	"formatDate": formatDate,
	"formatDateForForm": formatDateForForm,
	"formatSolution": formatSolution,
}

var templates map[string]*template.Template

// Load templates on program initialisation
func InitTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	tmpls, err := filepath.Glob(path.Join(BasePath, "app/views/*.html"))
	if err != nil {
		log.Fatal(err)
	}

	partials, err := filepath.Glob(path.Join(BasePath, "app/views/partials/*.html"))
	if err != nil {
		log.Fatal(err)
	}

	for _, tmpl := range tmpls {
		files := append(partials, path.Join(BasePath, "app/views/layout.html"))
		files = append(files, tmpl)
		templates[filepath.Base(tmpl)] = template.Must(template.New("base").Funcs(funcMap).ParseFiles(files...))
	}
}

// At the end of every GET controller action
func Render(w http.ResponseWriter, r *http.Request, filename string, props interface{}) {
	tmpl := templates[filename]

	endProps := make(map[string]interface{})
	for k, v := range *props.(*Props) {
		endProps[k] = v
	}

	data, ok := r.Context().Value("data").(*Props)

	if ok {
		for k, v := range *data {
			endProps[k] = v
		}
	}

	if err := tmpl.ExecuteTemplate(w, "layout", endProps); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NotAuthorized(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/signup", 307)
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "403 forbidden")
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	Render(w, r, "404.html", &Props{})
}

package utils

import (
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, filename string, props interface{}) {
	tmpl := templates[filename]

	if tmpl != nil {
		data, _ := r.Context().Value("data").(*Props)
		for k, v := range *props.(*Props) {
			(*data)[k] = v
		}

		if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		NotFound(w, r)
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

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "400 bad request")
}

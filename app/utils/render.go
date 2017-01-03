package utils

import (
	"fmt"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, filename string, props interface{}) {
	tmpl := templates[filename]

	if tmpl != nil {
		data := r.Context().Value("data")
		if data != nil {
			for k, v := range *data.(*Props) {
				(*props.(*Props))[k] = v
			}
		}

		if err := tmpl.ExecuteTemplate(w, "layout", props); err != nil {
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
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 not found")
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "400 bad request")
}

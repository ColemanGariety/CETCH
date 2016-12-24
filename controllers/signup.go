package controllers

import (
	"log"
	"net/http"
	"html/template"

	"github.com/JacksonGariety/wetch/models"
)

func SignupShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/signup.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// handle parse error
	}

	name := r.PostFormValue("username")
	password := r.PostFormValue("password")
	passwordConfirmation := r.PostFormValue("password_confirmation")

	if password != passwordConfirmation {
		return
		// flash error
	}

	if models.UserExistsByName(name)  {
		log.Println("user exists")
		http.Redirect(w, r, "/signup", http.StatusFound)
		// flash user exists
		return
	} else {
		err := models.UserCreate(name, password)
		log.Println(err)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

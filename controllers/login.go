package controllers

import (
	"log"
	"net/http"
	"html/template"
	"golang.org/x/crypto/bcrypt"

	"github.com/JacksonGariety/wetch/models"
)

func LoginShow(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/layout.html", "views/login.html")
	t.ExecuteTemplate(w, "layout", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// handle parse error
	}

	name := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, _ := models.UserByName(name)

	if user == nil {
		log.Println("user does not exist")
		// flash user does not exist
		return
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err == nil {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
			// flash password is wrong
		}
	}
}

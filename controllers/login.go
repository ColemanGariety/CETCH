package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"

	"github.com/JacksonGariety/wetch/models"
	"github.com/JacksonGariety/wetch/utils"
)

// Actions

func LoginShow(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "login.html", nil)
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	form := &LoginForm{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if (form.Validate() == false) {
		utils.Render(w, "login.html", form)
	} else {
		utils.Render(w, "index.html", nil)
	}
}

// Validations

type LoginForm struct {
	utils.Form
	Username string
	Password string
}

func (form *LoginForm) Validate() bool {
	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Password, "Password")
	form.ValidatePresence(form.Username, "Username")

	if form.IsValid() {
		user, err := models.UserByName(form.Username)
		if err != nil {
			form.SetError("Username", "Username does not exist")
		} else {
			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password))
			if err != nil {
				form.SetError("Password", "Password is incorrect")
			}
		}
	}

	return form.IsValid()
}

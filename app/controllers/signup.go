package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/middleware"
)

// Actions

func SignupShow(w http.ResponseWriter, r *http.Request) {
	if _, ok := middleware.CurrentUser(r); !ok {
		utils.Render(w, "signup.html", nil)
	} else {
		http.Redirect(w, r, "/profile", 307)
	}
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	form := &SignupForm{
		Email: r.FormValue("email"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	if (form.validate() == false) {
		utils.Render(w, "signup.html", &utils.Props{
			"errors": form.Errors,
			"email": form.Email,
			"username": form.Username,
			"password": form.Password,
			"passwordConfirmation": form.PasswordConfirmation,
		})
	} else {
		(&models.User{ Name: form.Username, Email: form.Email }).CreateFromPassword(form.Password)
		http.Redirect(w, r, "/profile", 307)
	}
}

// Validations

type SignupForm struct {
	models.Form
	Email                string
	Username             string
	Password             string
	PasswordConfirmation string
}

func (form *SignupForm) validate() (bool) {
	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Email, "Email")

	if form.ValidatePresence(form.Password, "Password") {
		form.ValidateLength(form.Password, "Password", 5, 30)
	}

	if form.ValidatePresence(form.Username, "Username") {
		form.ValidateNoSpace(form.Username, "Username")

		exists, _ := (&models.User{ Name: form.Username }).Exists()
		if exists {
			form.SetError("Username", "Username is already in use")
		}
	}

	if form.ValidatePresence(form.PasswordConfirmation, "PasswordConfirmation") {
		form.ValidateConfirmation(form.Password, "Password", form.PasswordConfirmation, "PasswordConfirmation")
	}

	return form.IsValid()
}

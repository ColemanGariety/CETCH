package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

// Actions

func SignupShow(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "signup.html", &utils.Props{})
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors":                make(map[string]string),
		"email":                 r.FormValue("email"),
		"username":              r.FormValue("username"),
		"password":              r.FormValue("password"),
		"password_confirmation": r.FormValue("password_confirmation"),
	}

	if validateSignupForm(form) == false {
		utils.Render(w, r, "signup.html", &form)
	} else {
		(&models.User{Name: form["username"].(string), Email: form["email"].(string)}).CreateFromPassword(form["password"].(string))
		signedToken, expireCookie, claims := models.ClaimsCreate(form["username"].(string)) // creates a JWT token
		cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, claims.Userpath(), 307)
	}
}

// Validations

func validateSignupForm(form utils.Props) bool {
	if form.ValidatePresence("email") {
		if form.ValidateEmail("email") {
			exists, _ := (&models.User{Email: form["email"].(string)}).Exists()
			if exists {
				form.SetError("email", "email is already in use")
			}
		}
	}

	if form.ValidatePresence("password") {
		form.ValidateLength("password", 5, 30)
	}

	if form.ValidatePresence("username") {
		form.ValidateNoSpace("username")

		exists, _ := (&models.User{Name: form["username"].(string)}).Exists()
		if exists {
			form.SetError("username", "username is already in use")
		}
	}

	if form.ValidatePresence("password_confirmation") {
		form.ValidateConfirmation("password", "password_confirmation")
	}

	return form.IsValid()
}

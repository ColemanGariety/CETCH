package controllers

import (
	"net/http"

	"github.com/JacksonGariety/wetch/models"
	"github.com/JacksonGariety/wetch/utils"
)

// Actions

func SignupShow(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "signup.html", nil)
}

func SignupPost(w http.ResponseWriter, r *http.Request) {
	form := &SignupForm{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
		PasswordConfirmation: r.FormValue("password_confirmation"),
	}

	if (form.Validate() == false) {
		utils.Render(w, "signup.html", form)
	} else {
		models.UserCreate(form.Username, form.Password)
		utils.Render(w, "index.html", nil)
	}
}
// Validations

type SignupForm struct {
	utils.Form
	Username             string
	Password             string
	PasswordConfirmation string
}

func (form *SignupForm) Validate() bool {
	form.Errors = make(map[string]string)

	form.ValidatePresence(form.Password, "Password")
	form.ValidatePresence(form.PasswordConfirmation, "PasswordConfirmation")

	if form.ValidatePresence(form.Username, "Username") {
		form.ValidateNoSpace(form.Username, "Username")

		if models.UserExistsByName(form.Username)  {
			form.SetError("Username", "Username is already in use")
		}
	}

	form.ValidateConfirmation(form.Password, "Password", form.PasswordConfirmation, "PasswordConfirmation")

	return form.IsValid()
}

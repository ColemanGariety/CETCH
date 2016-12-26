package controllers

import (
	"net/http"
	// "net/smtp"
	// "log"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

// Actions

func ForgottenShow(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "forgotten.html", nil)
}

func ForgottenPost(w http.ResponseWriter, r *http.Request) {
	form := models.Form{
		"errors": make(map[string]string),
		"email": r.FormValue("email"),
	}

	if (validateForgottenForm(form) == false) {
		utils.Render(w, "forgotten.html", form)
	} else {
		// err := smtp.SendMail(
		// 	"mail.example.com:25",
		// 	smtp.PlainAuth(
		// 		"",
		// 		"user@example.com",
		// 		"password",
		// 		"mail.example.com",
		// 	),
		// 	"sender@example.org",
		// 	[]string{form["email"]},
		// 	[]byte("This is the email body."),
		// )
		// if err != nil {
		// 	log.Fatal(err)
		// }
	}
}

// Validations

func validateForgottenForm(form models.Form) (bool) {
	if form.ValidatePresence("email") {
		exists, _ := (&models.User{ Email: form["email"].(string) }).Exists()
		if !exists {
			form.SetError("email", "unrecognized email")
		}
	}
	return form.IsValid()
}

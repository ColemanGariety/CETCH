package controllers

import (
	"net/http"
	// "net/smtp"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

// Actions

func ForgottenShow(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "forgotten.html", &utils.Props{})
}

func ForgottenSendEmail(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors": make(map[string]string),
		"email":  r.FormValue("email"),
	}

	if validateForgottenForm(form) == false {
		utils.Render(w, r, "forgotten.html", &form)
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

func validateForgottenForm(form utils.Props) bool {
	if form.ValidatePresence("email") {
		exists := models.Exists(&models.User{Email: form["email"].(string)})
		if !exists {
			form.SetError("email", "unrecognized email")
		}
	}
	return form.IsValid()
}

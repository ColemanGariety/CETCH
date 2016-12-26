package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/middleware"
)

// Actions

func LoginShow(w http.ResponseWriter, r *http.Request) {
	if claims, ok := middleware.CurrentUser(r); !ok {
		utils.Render(w, "login.html", nil)
	} else {
		http.Redirect(w, r, claims.Userpath(), 307)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	form := models.Form{
		"errors": make(map[string]string),
		"username": r.FormValue("username"),
		"password": r.FormValue("password"),
	}

	if (validateLoginForm(form) == false) {
		utils.Render(w, "login.html", form)
	} else {
	  signedToken, expireCookie, claims := models.ClaimsCreate(form["username"].(string)) // creates a JWT token
		cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, claims.Userpath(), 307)
	}
}

func LogoutShow(w http.ResponseWriter, r *http.Request) {
    deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
    http.SetCookie(w, &deleteCookie)
		http.Redirect(w, r, "/", 307)
}

// Validations

func validateLoginForm(form models.Form) (bool) {
	hasPassword := form.ValidatePresence("password")

	if form.ValidatePresence("username") {
		user, err := (&models.User{ Name: form["username"].(string) }).Find()
		if err != nil {
			form.SetError("username", "invalid username or password")
		} else if hasPassword {
			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form["password"].(string)))
			if err != nil {
				form.SetError("username", "invalid username or password")
			}
		}
	}

	return form.IsValid()
}

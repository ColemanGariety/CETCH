package controllers

import (
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/JacksonGariety/wetch/app/models"
	"github.com/JacksonGariety/wetch/app/utils"
	"github.com/JacksonGariety/wetch/app/middleware"
)

// Actions

func LoginShow(w http.ResponseWriter, r *http.Request) {
	if _, ok := middleware.CurrentUser(r); !ok {
		utils.Render(w, "login.html", nil)
	} else {
		http.Redirect(w, r, "/profile", 307)
	}
}

func LoginPost(w http.ResponseWriter, r *http.Request) {
	form := &LoginForm{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if (form.validate() == false) {
		utils.Render(w, "login.html", &utils.Props{
			"errors": form.Errors,
			"username": form.Username,
			"password": form.Password,
		})
	} else {
	  signedToken, expireCookie := models.ClaimsCreate(form.Username) // creates a JWT token
		cookie := http.Cookie{Name: "Auth", Value: signedToken, Expires: expireCookie, HttpOnly: true}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/profile", 307)
	}
}

func LogoutShow(w http.ResponseWriter, r *http.Request){
    deleteCookie := http.Cookie{Name: "Auth", Value: "none", Expires: time.Now()}
    http.SetCookie(w, &deleteCookie)
		http.Redirect(w, r, "/", 307)
}

// Validations

type LoginForm struct {
	utils.Form
	Username string
	Password string
}

func (form *LoginForm) validate() (bool) {
	form.Errors = make(map[string]string)

	hasPassword := form.ValidatePresence(form.Password, "Password")

	if form.ValidatePresence(form.Username, "Username") {
		user, err := models.UserByName(form.Username)
		if err != nil {
			form.SetError("Username", "Username does not exist")
		} else if hasPassword {
			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(form.Password))
			if err != nil {
				form.SetError("Password", "Password is incorrect")
			}
		}
	}

	return form.IsValid()
}

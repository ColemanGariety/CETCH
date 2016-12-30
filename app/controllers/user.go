package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func UserShow(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: bone.GetValue(r, "name")}
	if exists, _ := user.Exists(); exists {
		utils.Render(w, r, "user.html", &utils.Props{
			"username": user.Name,
			"admin":    user.Admin,
		})
	} else {
		utils.NotFound(w, r)
	}
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	users, _ := (&models.Users{}).FindAll()
	utils.Render(w, r, "users_show.html", &utils.Props{
		"users": users,
	})
}

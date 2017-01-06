package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func UserShow(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: bone.GetValue(r, "name")}
	if models.Exists(&user) {
		entries := (&models.Entries{}).FindByUserId(user.ID)

		for i, entry := range *entries {
			models.DB.Model(entry).Related(&(*entries)[i].Competition)
		}

		utils.Render(w, r, "user.html", &utils.Props{
			"username": user.Name,
			"admin":    user.Admin,
			"entries": entries,
		})
	} else {
		utils.NotFound(w, r)
	}
}

func UsersShow(w http.ResponseWriter, r *http.Request) {
	users := &models.Users{}
	models.All(users)
	utils.Render(w, r, "users_show.html", &utils.Props{
		"users": users,
	})
}

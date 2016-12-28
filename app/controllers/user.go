package controllers

import (
	"fmt"
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
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

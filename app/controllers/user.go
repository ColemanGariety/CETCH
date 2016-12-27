package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"

	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/models"
)

func UserShow(w http.ResponseWriter, r *http.Request){
	user := &models.User{ Name: mux.Vars(r)["name"] }
	if exists, _ := user.Exists(); exists {
		utils.Render(w, r, "user.html", &utils.Props{
			"username": user.Name,
			"admin": user.Admin,
		})
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

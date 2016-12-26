package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	// "log"

	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/middleware"
)

func UserShow(w http.ResponseWriter, r *http.Request){
	user := &models.User{ Name: mux.Vars(r)["name"] }
	if exists, _ := user.Exists(); exists {
		claims, ok := middleware.CurrentUser(r)
		utils.Render(w, "user.html", &utils.Props{
			"authorized": ok,
			"authorized_username": claims.Username,
			"username": user.Name,
		})
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

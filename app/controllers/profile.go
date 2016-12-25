package controllers

import (
	"net/http"

	"github.com/JacksonGariety/wetch/app/utils"
	"github.com/JacksonGariety/wetch/app/middleware"
)

func ProfileShow(w http.ResponseWriter, r *http.Request){
	if claims, ok := middleware.CurrentUser(r); !ok {
		http.Redirect(w, r, "/login", 307)
	} else {
		utils.Render(w, "profile.html", &utils.Props{
			"authorized": ok,
			"username": claims.Username,
		})
	}
}

package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/middleware"
)

func Index(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.CurrentUser(r)

	utils.Render(w, "index.html", &utils.Props{
		"authorized": ok,
		"authorized_username": claims.Username,
		"userpath": claims.Userpath(),
	})
}

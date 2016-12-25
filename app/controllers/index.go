package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/utils"
	"github.com/JacksonGariety/cetch/app/middleware"
)

func Index(w http.ResponseWriter, r *http.Request) {
	_, ok := middleware.CurrentUser(r)

utils.Render(w, "index.html", &utils.Props{
		"authorized": ok,
	})
}

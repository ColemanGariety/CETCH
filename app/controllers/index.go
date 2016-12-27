package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "index.html", &utils.Props{})
}

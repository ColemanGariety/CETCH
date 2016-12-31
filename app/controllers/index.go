package controllers

import (
	"net/http"
	// "time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	utils.Render(w, r, "index.html", &utils.Props{
		"current_competition": comp,
	})
}

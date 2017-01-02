package controllers

import (
	"net/http"
	// "time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]
	entry := &models.Entry{
		UserID: current_user.(models.User).ID,
		CompetitionID: comp.ID,
	}
	models.Find(entry)
	utils.Render(w, r, "index.html", &utils.Props{
		"competition": comp,
		"entry": entry,
	})
}

func Rules(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "rules.html", &utils.Props{})
}

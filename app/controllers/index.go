package controllers

import (
	"net/http"
	// "time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()
	currentUser := (*r.Context().Value("data").(*utils.Props))["current_user"]

	entry := models.Entry{}
	if currentUser != nil {
		entry.UserID = currentUser.(models.User).ID
		entry.CompetitionID = comp.ID
	}

	models.DB.Where(&entry).First(&entry)

	utils.Render(w, r, "index.html", &utils.Props{
		"competition": comp,
		"entry": entry,
	})
}

func Rules(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "rules.html", &utils.Props{})
}

package controllers

import (
	"net/http"
	// "time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	comp, _ := (&models.Competition{}).Current()

	var entry models.Entry
	currentUser := (*r.Context().Value("data").(*utils.Props))["current_user"]
	if currentUser != nil {
		entry = models.Entry{
			UserID: currentUser.(models.User).ID,
			CompetitionID: comp.ID,
		}
	}

	models.Find(&entry)

	utils.Render(w, r, "index.html", &utils.Props{
		"competition": comp,
		"entry": entry,
	})
}

func Rules(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "rules.html", &utils.Props{})
}

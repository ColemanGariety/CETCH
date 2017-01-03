package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	prev, _ := (&models.Competition{}).Previous()
	comp, _ := (&models.Competition{}).Current()

	entry := models.Entry{}
	data := r.Context().Value("data")
	if data != nil {
		currentUser := (*data.(*utils.Props))["current_user"]
		if currentUser != nil {
			entry.UserID = currentUser.(models.User).ID
			entry.CompetitionID = comp.ID
		}

		models.DB.Where(&entry).First(&entry)
		models.DB.Model(entry).Related(&entry.Competition)
	}

	winner := prev.Winner()
	models.DB.Model(winner).Related(&winner.Competition)

	utils.Render(w, r, "index.html", &utils.Props{
		"competition": comp,
		"entry": entry,
		"winner": winner,
	})
}

func Rules(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "rules.html", &utils.Props{})
}

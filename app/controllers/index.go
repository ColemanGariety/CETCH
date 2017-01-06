package controllers

import (
	"net/http"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	prev, _ := (&models.Competition{}).Previous()
	curr, _ := (&models.Competition{}).Current()

	winner := prev.Winner()

	utils.Render(w, r, "index.html", &utils.Props{
		"competition": curr,
		"winner": winner,
	})
}

func Rules(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "rules.html", &utils.Props{})
}

func About(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "about.html", &utils.Props{})
}

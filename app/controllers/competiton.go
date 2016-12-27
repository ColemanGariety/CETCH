package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func CompetitionsShow(w http.ResponseWriter, r *http.Request) {
	competitions, _ := (&models.Competitions{}).FindAll()
	utils.Render(w, r, "competitions.html", &utils.Props{ "competitions": competitions })
}

func CompetitionShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	comp := &models.Competition{ Id: id }
	if exists, _ := comp.Exists(); exists {
		utils.Render(w, r, "competition_show.html", &utils.Props{ "competition": comp })
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

func CompetitionNew(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "competition_new.html", &utils.Props{})
}

func CompetitionCreate(w http.ResponseWriter, r *http.Request) {
	form := utils.Props{
		"errors": make(map[string]string),
		"name": r.FormValue("name"),
		"description": r.FormValue("description"),
	}

	if (validateCompetitionCreateForm(form) == false) {
		utils.Render(w, r, "competition_new.html", &form)
	} else {
		(&models.Competition{ Name: form["name"].(string), Description: form["description"].(string) }).Create()
		http.Redirect(w, r, "/competitions", 307)
	}
}

// Validations

func validateCompetitionCreateForm(form utils.Props) (bool) {
	return true
}

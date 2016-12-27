package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"

	"github.com/JacksonGariety/cetch/app/middleware"
	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func CompetitionsShow(w http.ResponseWriter, r *http.Request) {
	competitions, _ := (&models.Competitions{}).FindAll()
	user, ok := middleware.CurrentUser(r)
	utils.Render(w, "competitions.html", &utils.Props{
		"authorized": ok,
		"authorized_username": user.Name,
		"competitions": competitions,
		"userpath": user.Userpath(),
		"admin": ok && user.Admin,
	})
}

func CompetitionShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	comp := &models.Competition{ Id: id }
	if exists, _ := comp.Exists(); exists {
		user, ok := middleware.CurrentUser(r)
		utils.Render(w, "competition_show.html", &utils.Props{
			"authorized": ok,
			"authorized_username": user.Name,
			"competition": comp,
			"userpath": user.Userpath(),
		})
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

func CompetitionNew(w http.ResponseWriter, r *http.Request) {
	user, ok := middleware.CurrentUser(r)
	utils.Render(w, "competition_new.html", &utils.Props{
		"authorized": ok,
		"authorized_username": user.Name,
		"userpath": user.Userpath(),
	})
}

func CompetitionCreate(w http.ResponseWriter, r *http.Request) {
	form := models.Form{
		"errors": make(map[string]string),
		"name": r.FormValue("name"),
		"description": r.FormValue("description"),
	}

	if (validateCompetitionCreateForm(form) == false) {
		utils.Render(w, "competition_new.html", form)
	} else {
		(&models.Competition{ Name: form["name"].(string), Description: form["description"].(string) }).Create()
		http.Redirect(w, r, "/competitions", 307)
	}
}

// Validations

func validateCompetitionCreateForm(form models.Form) (bool) {
	return true
}

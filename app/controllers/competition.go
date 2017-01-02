package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"
	"strconv"
	"time"
	"fmt"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Archive(w http.ResponseWriter, r *http.Request) {
	comps := new(models.Competitions)
	models.Where(comps, "date < NOW() AND date != '0001-01-01'")
	utils.Render(w, r, "archive.html", &utils.Props{"competitions": comps})
}

func Current(w http.ResponseWriter, r *http.Request) {
	current, _ := new(models.Competition).Current()
	http.Redirect(w, r, fmt.Sprintf("/competition/%v", current.ID), 307)
}

func CompetitionShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp := &models.Competition{}

	if models.ExistsById(comp, id) {
		current_user := (*r.Context().Value("data").(*utils.Props))["current_user"]

		var entry *models.Entry
		if current_user != nil {
			entry = &models.Entry{
				UserID: current_user.(models.User).ID,
				CompetitionID: comp.ID,
			}

			models.DB.First(&entry)
		} else {
			entry = nil
		}

		utils.Render(w, r, "competition_show.html", &utils.Props{
			"competition": comp,
			"current": comp.IsCurrent(),
			"entry": entry,
		})
	} else {
		utils.NotFound(w, r)
	}
}

func CompetitionNew(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "competition_new.html", &utils.Props{})
}

func CompetitionCreate(w http.ResponseWriter, r *http.Request) {
	solution, _ := strconv.ParseFloat(r.FormValue("solution"), 64)
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))

	form := utils.Props{
		"errors":      make(map[string]string),
		"name":        r.FormValue("name"),
		"description": r.FormValue("description"),
		"date":        date,
		"solution":    solution,
	}

	form.ValidatePresence("name")
	form.ValidatePresence("description")
	form.ValidatePresence("solution")

	if form.IsValid() == false {
		utils.Render(w, r, "competition_new.html", &form)
	} else {
		models.Create(&models.Competition{
			Name: form["name"].(string),
			Description: form["description"].(string),
			Solution: form["solution"].(float64),
		})
		http.Redirect(w, r, "/archive", 307)
	}
}

func CompetitionEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp := models.Competition{}
	models.FindById(&comp, id)
	utils.Render(w, r, "competition_edit.html", &utils.Props{
		"competition": comp,
	})
}

func CompetitionUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	solution, _ := strconv.ParseFloat(r.FormValue("solution"), 64)
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))

	form := utils.Props{
		"errors":      make(map[string]string),
		"name":        r.FormValue("name"),
		"description": r.FormValue("description"),
		"date":        date,
		"solution":    solution,
	}

	form.ValidatePresence("name")
	form.ValidatePresence("description")
	form.ValidatePresence("solution")

	comp := models.Competition{}
	models.FindById(&comp, id)
	if form.IsValid() {
		comp.Name = form["name"].(string)
		comp.Description = form["description"].(string)
		comp.Date = form["date"].(time.Time)
		comp.Solution = form["solution"].(float64)
		models.Save(&comp)
	}

	utils.Render(w, r, "competition_edit.html", &utils.Props{
		"errors": form["errors"],
		"competition": comp,
	})
}

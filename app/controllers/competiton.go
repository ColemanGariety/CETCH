package controllers

import (
	"fmt"
	"github.com/go-zoo/bone"
	"net/http"
	"strconv"
	"time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func CompetitionsShow(w http.ResponseWriter, r *http.Request) {
	competitions, _ := (&models.Competitions{}).FindAll()
	utils.Render(w, r, "competitions.html", &utils.Props{"competitions": *competitions})
}

func CompetitionShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp := &models.Competition{}
	if exists, _ := comp.ExistsById(id); exists {
		utils.Render(w, r, "competition_show.html", &utils.Props{"competition": comp})
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 not found")
	}
}

func CompetitionNew(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, r, "competition_new.html", &utils.Props{})
}

func CompetitionCreate(w http.ResponseWriter, r *http.Request) {
	solution, _ := strconv.ParseFloat(r.FormValue("solution"), 64)

	form := utils.Props{
		"errors":      make(map[string]string),
		"name":        r.FormValue("name"),
		"description": r.FormValue("description"),
		"solution":    solution,
	}

	form.ValidatePresence("name")
	form.ValidatePresence("description")
	form.ValidatePresence("solution")

	if form.IsValid() == false {
		utils.Render(w, r, "competition_new.html", &form)
	} else {
		(&models.Competition{
			Name: form["name"].(string),
			Description: form["description"].(string),
			Solution: form["solution"].(float64),
		}).Create()
		http.Redirect(w, r, "/competitions", 307)
	}
}

type CompetitionDay struct {
	Date        string
	Competition models.Competition
}

func ScheduleShow(w http.ResponseWriter, r *http.Request) {
	date := utils.NextFriday()
	
	// sort competitions by date
	comps, _ := (&models.Competitions{}).Where("date > ? OR date = ?", time.Now(), "0001-01-01")

	competitionsByDate := make(map[time.Time]interface{})
	for _, v := range *comps {
		competitionsByDate[v.Date.UTC()] = v
	}

	days := [10]CompetitionDay{}
	for i := 0; i < 10; i++ {
		var comp models.Competition
		if competitionsByDate[date] == nil {
			comp = models.Competition{}
		} else {
			comp = competitionsByDate[date].(models.Competition)
			delete(competitionsByDate, date)
		}
		days[i] = CompetitionDay{
			Date: date.Format("Mon Jan 02 2006"),
			Competition: comp,
		}
		date = date.AddDate(0, 0, 7)
	}

	utils.Render(w, r, "schedule.html", &utils.Props{ "days": days, "comps": competitionsByDate })
}

func SchedulePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, values := range r.Form {   // range over map
		for _, value := range values {    // range over []string
			if value != "" {
				comp, _ := (&models.Competition{ Name: value }).Find()
				date, _ := time.Parse("Mon Jan 02 2006", key)
				comp.Date = date
				comp.Save()
			}
		}
	}

	ScheduleShow(w, r)
}

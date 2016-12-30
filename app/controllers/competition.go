package controllers

import (
	"github.com/go-zoo/bone"
	"net/http"
	"strconv"
	"time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func Archive(w http.ResponseWriter, r *http.Request) {
	comps, _ := (&models.Competitions{}).Where("date < NOW() OR date = NOW()")
	utils.Render(w, r, "archive.html", &utils.Props{"competitions": comps})
}

func CompetitionShow(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp := &models.Competition{}
	if exists, _ := comp.ExistsById(id); exists {
		utils.Render(w, r, "competition_show.html", &utils.Props{"competition": comp})
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
		(&models.Competition{
			Name: form["name"].(string),
			Description: form["description"].(string),
			Solution: form["solution"].(float64),
		}).Create()
		http.Redirect(w, r, "/archive", 307)
	}
}

func CompetitionEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp, _ := (&models.Competition{}).FindById(id)
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

	comp, _ := (&models.Competition{}).FindById(id)
	if form.IsValid() {
		comp.Name = form["name"].(string)
		comp.Description = form["description"].(string)
		comp.Date = form["date"].(time.Time)
		comp.Solution = form["solution"].(float64)
		comp.Save()
	} else {

	}

	utils.Render(w, r, "competition_edit.html", &utils.Props{ "competition": comp })
}

type CompetitionDay struct {
	Date        string
	Competition models.Competition
}

func ScheduleShow(w http.ResponseWriter, r *http.Request) {
	date := utils.NextFriday()

	// sort competitions by date
	comps, _ := (&models.Competitions{}).Where("date > NOW() OR date = ?", "0001-01-01")

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

	all, _ := (&models.Competitions{}).FindAll()

	utils.Render(w, r, "schedule.html", &utils.Props{
		"days": days,
		"comps": competitionsByDate,
		"all": all,
	})
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

func CompetitionJoin(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	comp := &models.Competition{}
	if exists, _ := comp.ExistsById(id); exists {
		utils.Render(w, r, "competition_join.html", &utils.Props{"competition": comp})
	} else {
		utils.NotFound(w, r)
	}
}

func CompetitionJoinComp(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(bone.GetValue(r, "id"))
	(&models.Entry{
		CompetitionID: id,
		UserID: 1,
	}).Create()
	http.Redirect(w, r, "/archive", 307)
}

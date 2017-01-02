package controllers

import (
	"net/http"
	"time"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

func ScheduleShow(w http.ResponseWriter, r *http.Request) {
	date := utils.NextSaturday()

	// sort competitions by date
	comps := &models.Competitions{}
	models.Where(comps, "date > NOW()")

	competitionsByDate := make(map[time.Time]interface{})
	for _, v := range *comps {
		competitionsByDate[v.Date.UTC()] = v
	}

	days := [10]models.Schedule{}
	for i := 0; i < 10; i++ {
		var comp models.Competition
		if competitionsByDate[date] == nil {
			comp = models.Competition{}
		} else {
			comp = competitionsByDate[date].(models.Competition)
		}
		days[i] = models.Schedule{
			Date: date.Format("Mon Jan 02 2006"),
			Competition: comp,
		}
		date = date.AddDate(0, 0, 7)
	}

	all := new(models.Competitions)
	models.DB.Order("created_at desc").Find(&all)

	unscheduled := new(models.Competitions)
	models.Where(unscheduled, "date = ?", "0001-01-01")

	utils.Render(w, r, "schedule.html", &utils.Props{
		"days": days,
		"comps": unscheduled,
		"all": all,
	})
}

func SchedulePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	for key, values := range r.Form {   // range over map
		for _, value := range values {    // range over []string
			if value != "" {
				comp := &models.Competition{ Name: value }
				models.Find(comp)
				date, _ := time.Parse("Mon Jan 02 2006", key)
				comp.Date = date
				models.Save(comp)
			}
		}
	}

	ScheduleShow(w, r)
}

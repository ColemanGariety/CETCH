package models

import (
	"time"
	"github.com/jinzhu/gorm"

	"github.com/JacksonGariety/cetch/app/utils"
)

type Competition struct {
	gorm.Model
	Name          string
	Description   string
	Position      int
	Date          time.Time
	Solution      float64
	Entries       []Entry
}

type Competitions []Competition

func (competition *Competition) Current() (*Competition, error) {
	c := DB.Order("date asc").Where("date = ?", utils.NextSaturday()).First(competition)
	return competition, c.Error
}

func (competition *Competition) Previous() (*Competition, error) {
	c := DB.Order("date asc").Where("date = ?", utils.LastSaturday()).First(competition)
	return competition, c.Error
}
func (competition Competition) IsCurrent() bool {
	return competition.Date.Equal(utils.NextSaturday())
}

func (competition Competition) AverageExecTime() float64 {
	entries := Entries{}
	DB.Select("exec_time").Where("competition_id = ?", competition.ID).Find(&entries)

	var avg float64

	for _, entry := range entries {
		avg = avg + entry.ExecTime
	}

	return avg / float64(len(entries))
}

func (competition *Competition) Winner() *Entry {
	winner := Entry{}
	DB.Order("exec_time asc").Where("competition_id = ?", competition.ID).First(&winner)
	DB.Model(winner).Related(&winner.User)
	DB.Model(winner).Related(&winner.Competition)
	return &winner
}

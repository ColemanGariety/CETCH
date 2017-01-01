package models

import (
	"time"
	"github.com/jinzhu/gorm"
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
	c := DB.Order("date asc").Where("date > NOW()").First(competition)
	return competition, c.Error
}

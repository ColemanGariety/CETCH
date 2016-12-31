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
	Entry       []Entry
}

type Competitions []Competition

func NewCompetition(name string, description string, position int) *Competition {
	return &Competition{
		Name:        name,
		Description: description,
		Position:    position,
	}
}

func (competition *Competition) Order(query string) (*gorm.DB) {
	return db.Order(query)
}

func (competition *Competition) Current() (*Competition, error) {
	c := competition.Order("date asc").Where("date > NOW()").First(competition)
	return competition, c.Error
}

func (competition *Competition) FirstWhere(query string, vars ...interface{}) (*Competition, error) {
	comp := &Competition{}
	c := db.Where(query, vars...).First(comp)
	return comp, c.Error
}

func (competition *Competition) Find() (*Competition, error) {
	c := db.Where(&competition).First(&competition)
	return competition, c.Error
}

func (competition *Competition) FindById(id int) (*Competition, error) {
	c := db.First(&competition, id)
	return competition, c.Error
}

func (competition *Competition) Exists() (bool, error) {
	c := db.Where(&competition).First(&competition)
	return !(c.RecordNotFound()), c.Error
}

func (competition *Competition) ExistsById(id int) (bool, error) {
	c := db.First(&competition, id)
	return !(c.RecordNotFound()), c.Error
}

func (competition *Competition) Create() (*Competition, error) {
	db.NewRecord(competition)
	c := db.Create(&competition)
	return competition, c.Error
}

func (competition *Competition) Delete() error {
	c := db.Delete(&competition)
	return c.Error
}

func (competition *Competition) Save() error {
	c := db.Save(&competition)
	return c.Error
}

func (competition *Competition) DeleteById(id int) error {
	c := db.Delete(&competition, id)
	return c.Error
}

func (competitions *Competitions) Where(query string, vars ...interface{}) (*Competitions, error) {
	comps := &Competitions{}
	c := db.Where(query, vars...).Find(comps)
	return comps, c.Error
}

func (competitions *Competitions) FindAll() (*Competitions, error) {
	c := db.Find(&competitions)
	return competitions, c.Error
}

func (competitions *Competitions) DeleteAll() error {
	c := db.Unscoped().Find(&competitions).Delete(Competitions{})
	return c.Error
}

package models

import (
	"github.com/jinzhu/gorm"
)

type Competition struct {
	gorm.Model
	Id          int    `gorm:"primary_key"`
	Name        string
	Description string
	Position    int
}

type Competitions []Competition

func NewCompetition(name string, description string, position int) *Competition {
	return &Competition{
		Name: name,
		Description: description,
		Position: position,
	}
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

func (competition *Competition) Create() (*Competition, error) {
	db.NewRecord(competition)
	c := db.Create(&competition)
	return competition, c.Error
}

func (competition *Competition) Delete() (error) {
	c := db.Delete(&competition)
	return c.Error
}

func (competitions *Competitions) FindAll() (*Competitions, error) {
	c := db.Find(&competitions)
	return competitions, c.Error
}

func (competitions *Competitions) DeleteAll() (error) {
	c := db.Unscoped().Find(&competitions).Delete(Competitions{})
	return c.Error
}

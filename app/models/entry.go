package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
	gorm.Model
	Language        string
	Code            string
	ExecTime        float64
	UserID          uint
	CompetitionID   uint
	Competition     Competition
	User            User `gorm:"ForeignKey:UserID"`
}

type Entries []Entry

func (entries *Entries) FindByUserId(id uint) *Entries {
	DB.Where("user_id = ?", id).Find(&entries)
	return entries
}

func (entry *Entry) TimesFaster() float64 {
	comp := new(Competition)
	DB.Model(&entry).Related(comp)
	return comp.AverageExecTime() / entry.ExecTime
}

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
	gorm.Model
	UserID          uint
	CompetitionID   uint
	Competition     Competition
	Language        string
	Code            string
	ExecTime        float64
}

type Entries []Entry

func (entries *Entries) FindByUserId(id uint) *Entries {
	DB.Where("user_id = ?", id).Find(&entries)
	return entries
}

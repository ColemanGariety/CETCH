package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
	gorm.Model
	UserID          uint
	CompetitionID   uint
	Language        string
	Code            string
	ExecTime        float64
}

type Entries []Entry

func (entry *Entry) Create() (*Entry, error) {
	db.NewRecord(entry)
	c := db.Create(&entry)
	return entry, c.Error
}

func (entry *Entry) Find() (*Entry, error) {
	c := db.Where(&entry).First(&entry)
	return entry, c.Error
}

func (entries *Entries) FindByUserId(id uint) *Entries {
	db.Where("user_id = ?", id).Find(&entries)
	return entries
}

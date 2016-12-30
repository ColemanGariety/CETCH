package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Entry struct {
    gorm.Model
    UserID          int
    CompetitionID   int
}

type Entries []Entry

func (entry *Entry) Create() (*Entry, error) {
	db.NewRecord(entry)
	c := db.Create(&entry)
	return entry, c.Error
}


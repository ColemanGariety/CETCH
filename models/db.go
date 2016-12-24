package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", "user=wetch dbname=wetch_development sslmode=disable")

	if err != nil {
		log.Panic(err)
	}
}

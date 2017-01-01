package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(dbstring string) {
	var err error
	DB, err = gorm.Open("postgres", dbstring)

	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}

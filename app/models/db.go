package models

import (
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB(dbname string) {
	var err error
	db, err = gorm.Open("postgres", fmt.Sprintf("user=cetch dbname=%s sslmode=disable", dbname))

	if err != nil {
		log.Panic(err)
	}
}

func CloseDB() {
	db.Close()
}

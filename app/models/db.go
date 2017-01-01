package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB(dbname string) {
	var err error
	DB, err = gorm.Open("postgres", fmt.Sprintf("user=cetch dbname=%s sslmode=disable", dbname))

	if err != nil {
		panic(err)
	}
}

func CloseDB() {
	DB.Close()
}

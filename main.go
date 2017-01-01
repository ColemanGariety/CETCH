package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

var Production = os.Getenv("env") == "production"

func main() {
	utils.InitTemplates()

	var port string
	if Production {
		port = ":80"
	} else {
		port = ":8080"
	}

	var dbstring string
	if Production {
		dbstring = os.Getenv("dbstring")
	} else {
		dbstring = fmt.Sprintf("user=cetch dbname=%s sslmode=disable", os.Getenv("dbname"))
	}

	log.Println(dbstring)

	models.InitDB(dbstring)

	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(port, NewRouter()))
}

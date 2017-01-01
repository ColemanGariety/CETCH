package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JacksonGariety/cetch/app/models"
	"github.com/JacksonGariety/cetch/app/utils"
)

var Production = os.Getenv("env") == "production"

func main() {
	utils.InitTemplates()
	models.InitDB(os.Getenv("dbname"))
	log.Println("Whispering...")

	var port string
	if Production {
		port = ":80"
	} else {
		port = ":8080"
	}

	log.Fatal(http.ListenAndServe(port, NewRouter()))
}

package main

import (
	"log"
	"net/http"
	"github.com/JacksonGariety/cetch/app/models"
)

func main() {
	models.InitDB()
	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}

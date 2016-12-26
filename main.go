package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JacksonGariety/cetch/app/models"
)

func main() {
	models.InitDB(os.Getenv("dbname"))
	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}

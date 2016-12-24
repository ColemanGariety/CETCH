package main

import (
	"log"
	"net/http"
	// "github.com/JacksonGariety/wetch/models"
)

func main() {
	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}

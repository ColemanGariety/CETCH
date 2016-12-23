package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Whispering...")
	log.Fatal(http.ListenAndServe(":8080", NewRouter()))
}

package main

import (
	"log"
	"net/http"

	"./controllers"
)

type User struct {
	Id int
	Name string
	PasswordHash string
	PasswordSalt string
}

type Session struct {

}

var users = []User{
	User{Id: 1, Name: "colby", PasswordHash: "", PasswordSalt: ""},
	User{Id: 2, Name: "clmn", PasswordHash: "", PasswordSalt: ""},
}

func main() {
	// Static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)

	// Serve
	log.Println("Whispering...")
	http.ListenAndServe(":3000", nil)
}

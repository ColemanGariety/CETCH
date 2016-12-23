package main

import (
	"time"
)

type User struct {
	Id           int
	Name         string
	PasswordHash string
	PasswordSalt string
	CreatedAt    time.Time
}

type Users []User

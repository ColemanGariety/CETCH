package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name         string
	PasswordHash string
}

type Users []User

var db, err = gorm.Open("postgres", "user=wetch dbname=wetch_development sslmode=disable")

func UserByName(name string) (*User, error) {
	user := User{}
	dbc := db.First(&user, "name = ?", name)

	return &user, dbc.Error
}

func UserExistsByName(name string) (bool) {
	var user User
	return !(db.Where("name = ?", name).First(&user).RecordNotFound())
}

func UserCreate(name string, password string) (error) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{Name: name, PasswordHash: string(passwordHash)}

	db.NewRecord(user)
	dbc := db.Create(&user)

	return dbc.Error
}

func UserDelete(name string) (error) {
	user, _ := UserByName(name)
	dbc := db.Delete(&user)

	return dbc.Error
}

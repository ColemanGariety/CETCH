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

func UserAll() ([]User) {
	users := []User{}
	db.Find(&users)
	return users
}

func UserByName(name string) (*User, error) {
	user := User{}
	c := db.First(&user, "name = ?", name)

	return &user, c.Error
}

func UserExistsByName(name string) (bool) {
	var user User
	return !(db.Where("name = ?", name).First(&user).RecordNotFound())
}

func UserCreate(name string, password string) (error) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{Name: name, PasswordHash: string(passwordHash)}

	db.NewRecord(user)
	c := db.Create(&user)

	return c.Error
}

func UserDelete(name string) (error) {
	user, err := UserByName(name)
	if err != nil {
		return err
	}

	c := db.Unscoped().Delete(&user)
	return c.Error
}

func UserDeleteAll() {
	users := []User{}
	db.Unscoped().Delete(&users)
}

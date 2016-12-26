package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email        string
	Name         string
	PasswordHash string
}

type Users []User

func (user *User) Find() (*User, error) {
	c := db.Where(&user).First(&user)
	return user, c.Error
}

func (user *User) Exists() (bool, error) {
	c := db.Where(&user).First(&user)
	return !(c.RecordNotFound()), c.Error
}

func (user *User) CreateFromPassword(password string) (*User, error) {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.PasswordHash = string(passwordHash)

	db.NewRecord(user)
	c := db.Create(&user)

	return user, c.Error
}

func (user *User) Delete() (error) {
	c := db.Delete(&user)
	return c.Error
}

func (users *Users) FindAll() (*Users, error) {
	c := db.Find(&users)
	return users, c.Error
}

func (users *Users) DeleteAll() (error) {
	c := db.Unscoped().Find(&users).Delete(Users{})
	return c.Error
}

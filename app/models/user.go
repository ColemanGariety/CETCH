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

func NewUser(email string, name string, password string) *User {
	return &User{
		Email: email,
		Name: name,
		PasswordHash: hashPassword(password),
	}
}

func (user *User) Find() (*User, error) {
	c := db.Where(&user).First(&user)
	return user, c.Error
}

func (user *User) Exists() (bool, error) {
	c := db.Where(&user).First(&user)
	return !(c.RecordNotFound()), c.Error
}

func (user *User) Create() (*User, error) {
	db.NewRecord(user)
	c := db.Create(&user)
	return user, c.Error
}

func (user *User) CreateFromPassword(password string) (*User, error) {
	user.PasswordHash = hashPassword(password)
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

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email        string
	Name         string
	PasswordHash string
	Admin        bool
	Entries      []Entry
}

type Users []User

func (user *User) CreateFromPassword(password string) (*User, error) {
	user.PasswordHash = hashPassword(password)
	DB.NewRecord(user)
	c := DB.Create(&user)
	return user, c.Error
}

func (user *User) Userpath() string {
	return fmt.Sprintf("/user/%s", user.Name)
}

func (user *User) CurrentEntry() *Entry {
	current := new(Entry)
	DB.Order("created_at asc").Select("exec_time, competition_id").Where("user_id = ?", user.ID).First(current)
	DB.Model(&current).Related(&current.Competition)
	return current
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

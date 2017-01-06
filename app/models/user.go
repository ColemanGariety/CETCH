package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
	"github.com/JacksonGariety/cetch/app/utils"
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
	c := DB.Order("exec_time asc").Select("exec_time, created_at, competition_id").Where("user_id = ? AND created_at >= ?", user.ID, utils.LastSaturday()).First(current)
	if c.Error == nil {
		DB.Model(&current).Related(&current.Competition)
		return current
	} else {
		return nil
	}
}

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

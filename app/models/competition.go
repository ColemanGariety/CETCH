package models

import (
	"github.com/jinzhu/gorm"
)

type Competition struct {
	gorm.Model
	Name        string
	Description string
	Position    int
}

type Competitions []Competition


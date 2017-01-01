package models

import (
	"github.com/jinzhu/gorm"
)

func Find(model interface{}) *gorm.DB {
	return DB.Where(model).First(model)
}

func Exists(model interface{}) bool {
	return !(DB.Where(model).First(model).RecordNotFound())
}

func ExistsById(model interface{}, id int) bool {
	return !(DB.First(model, id).RecordNotFound())
}

func Create(model interface{}) *gorm.DB {
	DB.NewRecord(model)
	return DB.Create(model)
}

func FindById(model interface{}, id int) *gorm.DB {
	return DB.First(model, id)
}

func Save(model interface{}) *gorm.DB {
	return DB.Save(model)
}

func All(models interface{}) *gorm.DB {
	return DB.Find(models)
}

func Where(models interface{}, query string, vars ...interface{}) *gorm.DB {
	return DB.Where(query, vars...).Find(models)
}

func DeleteAll(models interface{}) *gorm.DB {
	return DB.Unscoped().Delete(models)
}

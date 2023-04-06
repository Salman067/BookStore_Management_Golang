package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SetDB interface {
	GetDB() *gorm.DB
}
type Database struct {
	db *gorm.DB
}

func (DB Database) GetDB() *gorm.DB {
	return DB.db
}

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "root:Salman12#@/vivasoftlimited?charset=utf8&parseTime=True&loc=Local")
	// d, err := gorm.Open("mysql", "root:Salman12#@/project_golang?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	var DB SetDB = Database{d}
	return DB.GetDB()
}

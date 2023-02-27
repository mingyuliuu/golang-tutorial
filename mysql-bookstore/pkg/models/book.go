package models

import (
	"github.com/jinzhu/gorm"
	"mysql-bookstore/pkg/config"
)

var (
	database *gorm.DB
)

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication`
}

func init() {
	config.Connect()
	database = config.GetDB()
	database.AutoMigrate(&Book{})
}
package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mysql-bookstore/pkg/utils"
)

var (
	database *gorm.DB
)

func Connect() {
	username := utils.GetEnvVariable("USERNAME")
	password := utils.GetEnvVariable("PASSWORD")

	dbString := fmt.Sprintf("%s:%s@/golang_tutorial?charset=utf8&parseTime=True&loc=Local", username, password)
	db, err := gorm.Open("mysql", dbString)

	if err != nil {
		panic(err)
	}

	database = db
	fmt.Println("Connected to MySQL database successfully.")
}

func GetDB() *gorm.DB {
	return database
}

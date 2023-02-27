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

func (book *Book) CreateBook() *Book {
	database.NewRecord(book)
	database.Create(&book)

	return book
}

func GetAllBooks() []Book {
	var books []Book

	database.Find(&books)
	return books
}

func GetBookById(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := database.Where("ID=?", ID).Find(&getBook)

	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	database.Where("ID=?", ID).Delete(book)

	return book
}
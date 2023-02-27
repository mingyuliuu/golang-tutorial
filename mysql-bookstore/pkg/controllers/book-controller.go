package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"mysql-bookstore/pkg/utils"
	"mysql-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(res http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	jsonRes, _ := json.Marshal(newBooks)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func GetBookById(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID)
	jsonRes, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func CreateBook(res http.ResponseWriter, req *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(req, createBook)
	book := createBook.CreateBook()

	jsonRes, _ := json.Marshal(book)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func UpdateBook(res http.ResponseWriter, req *http.Request) {
	updateBook := &models.Book{}
	utils.ParseBody(req, updateBook)
	
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	jsonRes, _ := json.Marshal(bookDetails)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID)
	jsonRes, _ := json.Marshal(book)

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(jsonRes)
}

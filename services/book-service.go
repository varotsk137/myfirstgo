package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/varotsk137/myfirstgo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB
var err error

const dbURL = "root:tooRLQSyM@tcp(localhost:3306)/letsgobook?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	db, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []model.Book
	db.Preload("Author").Find(&books)
	json.NewEncoder(w).Encode(books)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var authors []model.Author
	db.Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	var book model.Book
	db.Preload("Author").Find(&book, params["id"])
	fmt.Println(book)
	json.NewEncoder(w).Encode(book)
}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Author.Aid = book.AuthorID
	result := db.Create(&book)
	fmt.Println(result)
	json.NewEncoder(w).Encode(book)
}

func UpdateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	params, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		panic(err.Error())
	}
	book.Bid = uint(params)
	book.Author.Aid = book.AuthorID
	result := db.Save(&book)
	fmt.Println(result)
	json.NewEncoder(w).Encode(book)
}

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		panic(err.Error())
	}
	var book model.Book
	book.Bid = uint(params)
	db.Delete(&book)
}

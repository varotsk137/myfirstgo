package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/varotsk137/myfirstgo/model"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var books []model.Book

func getBooks( w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook( w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get Params
	// Loop through books
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

func createBooks( w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBooks( w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = mux.Vars(r)["id"]
	for _, item := range books {
		if item.ID == book.ID {
			item = book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&model.Book{})
}

func deleteBooks( w http.ResponseWriter, r *http.Request )  {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)["id"]
	for index, item := range books {
		if item.ID == params {
			books = append(books[:index], books[index+1:]...)
			json.NewEncoder(w).Encode("Delete book id "+params)
			break
		}
	}
}

func main() {
	//Init Router
	r := mux.NewRouter()

	//Mock Data
	books = append(books, model.Book{
		ID:     "1",
		ISBN:   "448783",
		Title:  "First Book in the World",
		Author: &model.Author{
			Fname: "Tomas",
			Lname: "Alva",
		},
	})
	books = append(books, model.Book{
		ID:     "2",
		ISBN:   "418583",
		Title:  "Second Book in the World",
		Author: &model.Author{
			Fname: "Allen",
			Lname: "Shilva",
		},
	})
	books = append(books, model.Book{
		ID:     "3",
		ISBN:   "1933017",
		Title:  "Third Book in the World",
		Author: &model.Author{
			Fname: "Kriengkai",
			Lname: "Klaigum",
		},
	})

	//Route Handlers / Endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

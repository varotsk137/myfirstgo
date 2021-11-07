package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/varotsk137/myfirstgo/services"
	"log"
	"net/http"
)

func InitialRouter() {
	//Init Router
	r := mux.NewRouter()
	//Route Handlers / Endpoint
	r.HandleFunc("/api/books", services.GetBooks).Methods("GET")
	r.HandleFunc("/api/authors", services.GetAuthors).Methods("GET")
	r.HandleFunc("/api/book/{id}", services.GetBook).Methods("GET")
	r.HandleFunc("/api/books", services.CreateBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}", services.UpdateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", services.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	services.InitialMigration()
	InitialRouter()
}

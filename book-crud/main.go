package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
}

var books []Book

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	arg := mux.Vars(r)
	for index, obj := range books{
		if obj.ID == arg["id"]


	}
}

func main() {
	rter := mux.NewRouter()

	books = append(books, Book{ID: "12345", isbn: "1a2b3c", Title: "Book 1", Author: &Author{firstname: "Jeff", lastname: "Dauda"}})
	books = append(books, Book{ID: "24680", isbn: "2a4b6c", Title: "Book 2", Author: &Author{firstname: "Jeff", lastname: "Trojan0x"}})

	rter.HandleFunc("/books", getBooks).Methods("GET")
	rter.HandleFunc("/books/{id}", getBook).Methods("GET")
	rter.HandleFunc("/books", newBook).Methods("POST")
	rter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	rter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server running at port 2000")
	log.Fatal(http.ListenAndServe(":2000", rter))
}

package main

import (
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

var book []Book

func main() {
	rter := mux.NewRouter()

	rter.HandleFunc("/books", getBooks).Methods("GET")
	rter.HandleFunc("/books/{id}", getBook).Methods("GET")
	rter.HandleFunc("/books", newBook).Methods("POST")
	rter.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	rter.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server running at port 2000")
	log.Fatal(http.ListenAndServe(":2000", rter))
}

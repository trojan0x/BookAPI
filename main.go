package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arg := mux.Vars(r)
	for index, obj := range books {
		if obj.ID == arg["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arg := mux.Vars(r)
	for _, obj := range books {
		if obj.ID == arg["id"] {
			json.NewEncoder(w).Encode(obj)
			return
		}
	}
}

func newBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)



}


func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arg := mux.Vars(r)
	for _, obj := range books {
		if obj.ID == arg["id"] {
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&body)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)

		}

}



func main(){
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

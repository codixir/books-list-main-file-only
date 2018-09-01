package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gets all the books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Gets a single the book by its id")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Adds a book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Removes a book")
}

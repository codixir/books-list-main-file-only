package main

import (
	"encoding/json"
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

	books = append(books,
		Book{ID: "1", Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: "2", Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: "3", Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: "4", Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: "5", Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)

	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(&book)
		}
	}
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

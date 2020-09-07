package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Book Struct (Model)
type Book struct {
	ID     uuid.UUID `json:"id"`
	Isbn   string    `json:"isbn"`
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
}

//Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

func init() {
	uuid1, _ := uuid.NewUUID()
	uuid2, _ := uuid.NewUUID()
	//Mock  Data - @todo - implements DB
	books = append(books, Book{
		ID:    uuid1,
		Isbn:  "448743",
		Title: "Book One",
		Author: &Author{
			Lastname:  "Isaac",
			Firstname: "Henrique",
		},
	})
	books = append(books, Book{
		ID:    uuid2,
		Isbn:  "57699",
		Title: "Book Two",
		Author: &Author{
			Lastname:  "Lorena",
			Firstname: "Farias",
		},
	})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	//Loop through books and infd with id
	for _, item := range books {
		if item.ID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//Init Router
	router := mux.NewRouter()

	//Router Handlers and Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Println("Application is available...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

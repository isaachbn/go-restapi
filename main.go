package main

import (
	"fmt"
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

func getBooks(w http.ResponseWriter, r *http.Request)  {

}

func getBook(w http.ResponseWriter, r *http.Request)  {

}

func createBook(w http.ResponseWriter, r *http.Request)  {

}

func updateBook(w http.ResponseWriter, r *http.Request)  {

}

func deleteBook(w http.ResponseWriter, r *http.Request)  {

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

	log.Fatal(http.ListenAndServe(":8000", router))
	fmt.Println("Application is available...")
}

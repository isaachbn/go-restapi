package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restapi/http/hadlers"
)


func main() {
	//Init Router
	router := mux.NewRouter()

	//Router Handlers and Endpoints
	router.HandleFunc("/api/books", hadlers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", hadlers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", hadlers.CreateBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", hadlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", hadlers.DeleteBook).Methods("DELETE")
	log.Println("Application is available...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

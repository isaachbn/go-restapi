package hadlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"restapi/domain/books"
)

//Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books.Books)
}

//Get book find id
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	//Loop through books and infd with id
	for _, item := range books.Books {
		if item.ID.String() == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&books.Book{})
}

//Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book books.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID, _ = uuid.NewUUID() // Mock uuid -not safe
	books.Books = append(books.Books, book)
	json.NewEncoder(w).Encode(book)
}

//Update book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	for index, item := range books.Books {
		if item.ID.String() == params["id"] {
			books.Books = append(books.Books[:index], books.Books[index+1:]...)
			var book books.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID, _ = uuid.NewUUID() // Mock uuid -not safe
			books.Books = append(books.Books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books.Books)
}

//Delete book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	for index, item := range books.Books {
		if item.ID.String() == params["id"] {
			books.Books = append(books.Books[:index], books.Books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books.Books)
}

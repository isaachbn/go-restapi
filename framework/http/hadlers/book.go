package hadlers

import (
	"encoding/json"
	"net/http"
	"restapi/application/usecase"
	model "restapi/domain/books"
)

//Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := usecase.GetAllBooks()

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(w, http.StatusOK, books)
}

//Get book find id
func GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := getIdentifier(w, r)

	if err != nil {
		respondError(w, http.StatusNotFound, "Identifier is not valid")
		return
	}

	book, err := usecase.FindBook(id)

	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if book == nil {
		respondError(w, http.StatusNotFound, "object not found")
		return
	}

	respondJson(w, http.StatusOK, book)
}

//Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	result, err := usecase.CreateNewBook(&book)

	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(result)
}

//Update book
//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) //Get params
//
//	for index, item := range repositories.Books {
//		if item.ID.String() == params["id"] {
//			repositories.Books = append(repositories.Books[:index], repositories.Books[index+1:]...)
//			var book books.Book
//			_ = json.NewDecoder(r.Body).Decode(&book)
//			book.ID, _ = uuid.NewUUID() // Mock uuid -not safe
//			repositories.Books = append(repositories.Books, book)
//			json.NewEncoder(w).Encode(book)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(repositories.Books)
//}

//Delete book
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) //Get params
//
//	for index, item := range repositories.Books {
//		if item.ID.String() == params["id"] {
//			repositories.Books = append(repositories.Books[:index], repositories.Books[index+1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode(repositories.Books)
//}

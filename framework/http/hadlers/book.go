package hadlers

import (
	"encoding/json"
	"net/http"
	"restapi/application/usecase"
	model "restapi/domain/books"
)

//Get all books
func GetBooks(responseWriter http.ResponseWriter, request *http.Request) {
	books, err := usecase.GetAllBooks()

	if err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondJson(responseWriter, http.StatusOK, books)
}

//Get book find id
func GetBook(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := getIdentifier(responseWriter, request)

	if err != nil {
		respondError(responseWriter, http.StatusNotFound, "Identifier is not valid")
		return
	}

	book, err := usecase.FindBook(id)

	if err != nil {
		respondError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	if book == nil {
		respondError(responseWriter, http.StatusNotFound, "object not found")
		return
	}

	respondJson(responseWriter, http.StatusOK, book)
}

//Create new book
func CreateBook(responseWriter http.ResponseWriter, request *http.Request) {
	var book model.Book
	err := json.NewDecoder(request.Body).Decode(&book)
	defer request.Body.Close()

	if err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	result, err := usecase.CreateNewBook(&book)

	if err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(responseWriter).Encode(result)
}

//Update book
func UpdateBook(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := getIdentifier(responseWriter, request)

	if err != nil {
		respondError(responseWriter, http.StatusNotFound, "Identifier is not valid")
		return
	}

	var book model.Book
	err = json.NewDecoder(request.Body).Decode(&book)
	defer request.Body.Close()

	if err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	result, err := usecase.UpdateBook(id, &book)

	if err != nil {
		respondError(responseWriter, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(responseWriter).Encode(result)
}

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

package usecase

import (
	"restapi/application/repositories"
	model "restapi/domain/books"
)

var repository repositories.BookRepository

func init()  {
	repository = &repositories.BookRepositoryImpl{}
}

//Create new book after validate
func CreateNewBook(book *model.Book) (*model.Book, error)  {
	book, err := book.NewBook()

	if err != nil {
		return nil, err
	}

	book, err = repository.Insert(book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

//List all books
func GetAllBooks() ([]*model.Book, error) {
	books, err := repository.All()

	if err != nil {
		return nil, err
	}

	return books, nil
}
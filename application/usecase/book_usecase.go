package usecase

import (
	"github.com/gofrs/uuid"
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

//Find book by id
func FindBook(uuid uuid.UUID) (*model.Book, error)  {
	return repository.FindById(uuid)
}

func UpdateBook(uuid uuid.UUID, book *model.Book) (*model.Book, error)   {
	bookModel, err := FindBook(uuid)

	if err != nil {
		return nil, err
	}

	book.ID = bookModel.ID
	book, err = book.NewBook()

	if err != nil {
		return nil, err
	}

	repository.Update(book)

	return book, nil
}

func DeleteBook(uuid uuid.UUID) error {
	bookModel, err := FindBook(uuid)

	if err != nil {
		return err
	}

	err = repository.Delete(bookModel)

	if err != nil {
		return err
	}

	return nil
}
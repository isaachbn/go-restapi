package repositories

import (
	"github.com/gofrs/uuid"
	"log"
	model "restapi/domain/books"
	"restapi/framework/db"
)

type BookRepository interface {
	Insert(book *model.Book) (*model.Book, error)
	All() ([]*model.Book, error)
    FindById(uuid uuid.UUID) (*model.Book, error)
	Update(book *model.Book) (*model.Book, error)
}

type BookRepositoryImpl struct {}

//Insert book in database
func (_ *BookRepositoryImpl) Insert(book *model.Book) (*model.Book, error) {
	err := db.ConnectDB().Create(book).Error

	if err != nil {
		log.Fatalf("Error to persist book: %v", err)
		return nil, err
	}

	return book, nil
}

//List all books
func (_ *BookRepositoryImpl) All() ([]*model.Book, error) {
	var book model.Book
	rows, err := db.ConnectDB().Find(&book).Rows()
	defer rows.Close()
	var books []*model.Book

	if err != nil {
		log.Fatalf("Error to list book: %v", err)
		return nil, err
	}

	for rows.Next() {
		db.ConnectDB().ScanRows(rows, &book)
		books = append(books, &book)
	}

	return books, nil
}

func (_ *BookRepositoryImpl) FindById(uuid uuid.UUID) (*model.Book, error) {
	var book model.Book
	row := db.ConnectDB().Find(&book, "id = ?", uuid).Row()

	if row.Err() != nil {
		return nil, row.Err()
	}

	return &book, nil
}

func (_ *BookRepositoryImpl) Update(book *model.Book) (*model.Book, error)  {
	err := db.ConnectDB().Updates(book).Error

	if err != nil {
		log.Fatalf("Error to update book: %v", err)
		return nil, err
	}

	return book, nil
}
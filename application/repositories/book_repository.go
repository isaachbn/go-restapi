package repositories

import (
	"log"
	model "restapi/domain/books"
	"restapi/framework/db"
)

type BookRepository interface {
	Insert(book *model.Book) (*model.Book, error)
	All() ([]*model.Book, error)
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
	rows, err := db.ConnectDB().Table("books").Rows()
	defer rows.Close()
	var books []*model.Book

	if err != nil {
		log.Fatalf("Error to list book: %v", err)
		return nil, err
	}

	var book model.Book

	for rows.Next() {
		db.ConnectDB().ScanRows(rows, &book)
		books = append(books, &book)
	}

	return books, nil
}

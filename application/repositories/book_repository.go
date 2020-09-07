package repositories

import (
	"gorm.io/gorm"
	"log"
	model "restapi/domain/books"
)

type BookRepository interface {
	Insert(book *model.Book) (*model.Book, error)
	All() (*model.Book, error)
}

type BookRepositoryDb struct {
	Db *gorm.DB
}

func (repo BookRepositoryDb) Insert(book *model.Book) (*model.Book, error)  {
	err := repo.Db.Create(book).Error

	if err != nil {
		log.Fatalf("Error to persist book: %v", err)
		return nil, err
	}

	return book, nil
}

func (repo BookRepositoryDb) All() ([]model.Book, error)  {
	rows, err := repo.Db.Table("books").Rows()
	defer rows.Close()
	var books []model.Book

	if err != nil {
		log.Fatalf("Error to list book: %v", err)
		return nil, err
	}

	var book model.Book

	for rows.Next() {
		repo.Db.ScanRows(rows, books)
		books = append(books, book)
	}

	return books, nil
}
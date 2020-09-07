package books

import (
	. "restapi/domain/base"
	"restapi/framework/validator"
)

type BookDomain interface {
	NewBook() (*Book, error)
}

//Book Struct (Model)
type Book struct {
	Base
	Isbn  string `json:"isbn" gorm:"type:text" validate:"required"`
	Title string `json:"title" gorm:"type:text" validate:"required"`
}

func (book *Book) validate() error  {
	validate := validator.GetValidator()
	return validate.Struct(book)
}

func (book *Book) NewBook() (*Book, error)  {
	book.GenerateIdentifier()
	err := book.validate()

	if err != nil {
		return nil, err
	}

	return book, nil
}
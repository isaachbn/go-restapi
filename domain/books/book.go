package books

import (
	"github.com/gofrs/uuid"
	. "restapi/domain/base"
	"restapi/framework/validator"
)

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
	book.ID, _ = uuid.NewV4()
	err := book.validate()

	if err != nil {
		return nil, err
	}

	return book, nil
}

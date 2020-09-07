package books

import (
	. "restapi/domain/base"
)

//Book Struct (Model)
type Book struct {
	Base
	Isbn   string  `json:"isbn" gorm:"type:text"`
	Title  string  `json:"title" gorm:"type:text"`
}

func NewBook() *Book {
	return &Book{}
}

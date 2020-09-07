package books

import "github.com/google/uuid"

//Book Struct (Model)
type Book struct {
	ID     uuid.UUID `json:"id"`
	Isbn   string    `json:"isbn"`
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
}

//Author Struct (Model)
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

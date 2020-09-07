package books

import "github.com/google/uuid"

//Init Books var as a slice Book struct
var Books []Book

func init() {
	uuid1, _ := uuid.NewUUID()
	uuid2, _ := uuid.NewUUID()
	//Mock  Data - @todo - implements DB
	Books = append(Books, Book{
		ID:    uuid1,
		Isbn:  "448743",
		Title: "Book One",
		Author: &Author{
			Lastname:  "Isaac",
			Firstname: "Henrique",
		},
	})
	Books = append(Books, Book{
		ID:    uuid2,
		Isbn:  "57699",
		Title: "Book Two",
		Author: &Author{
			Lastname:  "Lorena",
			Firstname: "Farias",
		},
	})
}

package books

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewBook(t *testing.T) {
	book := Book{Title: "test", Isbn: "test"}
	_, err := book.NewBook()
	require.Nil(t, err)

	book = Book{Title: "", Isbn: "test"}
	_, err = book.NewBook()
	require.Error(t, err)

	book = Book{Title: "test", Isbn: ""}
	_, err = book.NewBook()
	require.Error(t, err)
}

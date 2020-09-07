package books

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestNewBook(t *testing.T) {
	_, err := NewBook("46677", "Test Book")
	require.Nil(t, err)

	_, err = NewBook("", "Test Book")
	require.Error(t, err)

	_, err = NewBook("33322", "")
	require.Error(t, err)
}

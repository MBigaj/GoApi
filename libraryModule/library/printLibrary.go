package library

import "fmt"

type PrintLibrary struct {
	Library *Library
}

func (PrintLibrary PrintLibrary) SetLibrary(library *Library) {
	PrintLibrary.Library = library
}

func (PrintLibrary PrintLibrary) GetBookSelection() string {
	var bookString string

	for index, book := range PrintLibrary.Library.Books {
		bookString += fmt.Sprintf("%d - %s", index+1, book) + "\n"
	}

	return bookString
}

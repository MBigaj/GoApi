package library

import "fmt"

type Library struct {
	Orders []map[*User]Order
	Books  []Book
}

func (Library Library) ShowBooks() string {
	var bookString string

	for _, book := range Library.Books {
		bookString += fmt.Sprint(book) + "\n"
	}

	return bookString
}

func (Library Library) GetRemainingBooksCount() uint32 {
	var totalCount uint32

	for _, book := range Library.Books {
		totalCount += uint32(book.TotalNumber)
	}

	return totalCount
}

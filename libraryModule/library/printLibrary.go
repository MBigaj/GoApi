package library

import (
	"fmt"
)

type PrintLibrary struct {
	Library *Library
}

func GetMenu() string {
	menu := `
		1 - Rent a book
		2 - Add a book to the library
		3 - Give a rented book back
		4 - Check all current orders
		9 - Exit
	`

	return menu
}

func (PrintLibrary PrintLibrary) GetBookSelection() string {
	var bookString string

	for index, book := range PrintLibrary.Library.Books {
		bookString += fmt.Sprintf("%d - %s", index+1, book) + "\n"
	}

	return bookString
}

func (PrintLibrary PrintLibrary) GetRemainingBooksCount(book ...*Book) string {
	remainingCount := ""

	if book == nil {
		remainingCount = fmt.Sprintf("We have a total of %d books in our library!\n", PrintLibrary.Library.GetRemainingBooksCount())
	} else {
		remainingCount = fmt.Sprintf("%s - %d remaining for renting", *book[0], book[0].TotalNumber)
	}

	return remainingCount
}

func (PrintLibrary PrintLibrary) GetAllOrders() string {
	orderList := ""

	for _, orders := range PrintLibrary.Library.Orders {
		for user, order := range orders {
			orderList += fmt.Sprintf("%s rented: %s, has time till %s \n", user.Name, order.Book.Name, order.GiveBackDate)
		}
	}

	return orderList
}

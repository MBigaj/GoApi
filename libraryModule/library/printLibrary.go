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
	if book == nil {
		return fmt.Sprintf("We have a total of %d books in our library!\n", PrintLibrary.Library.GetRemainingBooksCount())
	} else {
		return fmt.Sprintf("%s - %d remaining for renting", *book[0], book[0].TotalNumber)
	}
}

func (PrintLibrary PrintLibrary) GetAllOrders(user ...*User) string {
	orderList := ""

	if len(user) <= 0 {
		for _, orders := range PrintLibrary.Library.Orders {
			for userFromOrder, order := range orders {
				if !order.Completed {
					orderList += fmt.Sprintf("%s rented: %s, has time till %s \n", userFromOrder.Name, order.Book.Name, order.GiveBackDate)
				}
			}
		}

		return orderList
	}

	for index, orders := range PrintLibrary.Library.Orders {
		for userFromOrder, order := range orders {
			if user[0] == userFromOrder {
				if !order.Completed {
					orderList += fmt.Sprintf("%d - %s rented: %s, has time till %s \n", index+1, userFromOrder.Name, order.Book.Name, order.GiveBackDate)
				}
			}
		}
	}

	if orderList == "" {
		return fmt.Sprintf("No orders available for user %s", user[0].Name)
	}

	return orderList
}

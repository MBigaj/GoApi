package library

import (
	"fmt"
	"time"
)

type LibraryHandler struct {
	Library *Library
}

func (LibraryHandler LibraryHandler) HandleRental(user *User, book *Book) {
	order := Order{
		Id:           uint16(LibraryHandler.Library.GetOrderCount()),
		User:         user,
		Book:         book,
		GiveBackDate: time.Now().AddDate(0, 0, 30),
		Completed:    false,
	}

	if book.TotalNumber <= 0 {
		return
	}

	user.AddOrder(&order)

	LibraryHandler.Library.Orders[user] = append(LibraryHandler.Library.Orders[user], &order)
	book.TotalNumber -= 1
}

func (LibraryHandler LibraryHandler) HandleNewBook() {
	var name string
	var writer string
	var state uint8
	var howMany uint16

	fmt.Println("Book name: ")
	fmt.Scan(&name)

	fmt.Println("Writer: ")
	fmt.Scan(&writer)

	fmt.Println("Book state: ")
	fmt.Scan(&state)

	fmt.Println("Books amount: ")
	fmt.Scan(&howMany)

	book := Book{
		Name:        name,
		Writer:      writer,
		State:       state,
		TotalNumber: howMany,
	}

	LibraryHandler.Library.Books = append(LibraryHandler.Library.Books, book)
}

func (LibraryHandler LibraryHandler) HandleBookReturn(user *User, book *Book) {
	usersOrders := LibraryHandler.Library.Orders[user]

	for _, order := range usersOrders {
		if order.Book == book && !order.Completed {
			order.Completed = true
			book.TotalNumber++
			break
		}
	}
}

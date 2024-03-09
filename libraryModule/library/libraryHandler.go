package library

import "time"

type LibraryHandler struct {
	Library *Library
}

func (LibraryHandler LibraryHandler) HandleRental(user *User, book *Book) {
	order := Order{
		Id:           uint16(LibraryHandler.Library.GetOrderCount()),
		UserId:       user.Id,
		Book:         book,
		GiveBackDate: time.Now().AddDate(0, 0, 30),
	}

	if book.TotalNumber <= 0 {
		return
	}

	newOrder := map[*User]*Order{
		user: &order,
	}

	LibraryHandler.Library.Orders = append(LibraryHandler.Library.Orders, newOrder)
	book.TotalNumber -= 1
}

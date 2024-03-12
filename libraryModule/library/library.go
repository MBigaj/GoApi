package library

type Library struct {
	Orders map[*User][]*Order
	Books  []Book
}

func CreateLibrary() *Library {
	return &Library{
		Orders: make(map[*User][]*Order),
		Books:  loadBooks(),
	}
}

func (Library Library) GetRemainingBooksCount() uint32 {
	var totalCount uint32

	for _, book := range Library.Books {
		totalCount += uint32(book.TotalNumber)
	}

	return totalCount
}

func (Library Library) GetOrderCount() uint32 {
	orderCount := 0

	for _, order := range Library.Orders {
		orderCount += len(order)
	}

	return uint32(orderCount)
}

func (Library Library) GetBookById(id uint16) *Book {
	for _, book := range Library.Books {
		if book.Id == id {
			return &book
		}
	}

	return nil
}

func loadBooks() []Book {
	return []Book{
		{Id: 1, Name: "The newest book", Writer: "Nick", State: 7, TotalNumber: 10},
		{Id: 2, Name: "The crazy book", Writer: "Big", State: 5, TotalNumber: 20},
		{Id: 3, Name: "The not so good book", Writer: "Very", State: 3, TotalNumber: 37},
		{Id: 4, Name: "The odd book", Writer: "Nice", State: 7, TotalNumber: 20},
		{Id: 5, Name: "The greatest book", Writer: "Guy", State: 10, TotalNumber: 15},
	}
}

package library

type Library struct {
	Orders []map[*User]*Order
	Books  []Book
}

func (Library Library) GetRemainingBooksCount() uint32 {
	var totalCount uint32

	for _, book := range Library.Books {
		totalCount += uint32(book.TotalNumber)
	}

	return totalCount
}

func (Library Library) GetOrderCount() uint32 {
	return uint32(len(Library.Orders))
}

func (Library Library) GetBookByIndex(index uint16) *Book {
	return &Library.Books[index]
}

package library

import "fmt"

type Book struct {
	Name        string
	Writer      string
	State       uint8
	TotalNumber uint16
}

func (Book Book) String() string {
	return fmt.Sprintf("Book: %s, Writer: %s, Amount: %d", Book.Name, Book.Writer, Book.TotalNumber)
}

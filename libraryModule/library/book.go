package library

import "fmt"

type Book struct {
	Id          uint16
	Name        string
	Writer      string
	State       uint8
	TotalNumber uint16
}

func (Book Book) String() string {
	return fmt.Sprintf("Book: %s, Writer: %s, State: %d", Book.Name, Book.Writer, Book.State)
}

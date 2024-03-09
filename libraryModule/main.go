package main

import (
	lib "Library/library"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the library!")

	library := lib.Library{Books: loadBooks()}
	printLibrary := lib.PrintLibrary{Library: &library}
	libraryHandler := lib.LibraryHandler{Library: &library}
	user := lib.User{Id: 1, Name: "Nice", Age: 23}

	var output rune

	for output != 9 {
		fmt.Print(lib.GetMenu())
		fmt.Scan(&output)

		switch output {
		case 1:
			fmt.Println(printLibrary.GetBookSelection())
			var selection uint16

			fmt.Scan(&selection)

			selectedBook := library.GetBookByIndex(selection - 1)

			fmt.Println(printLibrary.GetRemainingBooksCount(selectedBook))

			libraryHandler.HandleRental(&user, selectedBook)
		case 2:
			// TODO IMPLEMENT
		case 3:
			// TODO IMPLEMENT
		case 4:
			fmt.Println(printLibrary.GetAllOrders())
		}
	}
}

func loadBooks() []lib.Book {
	return []lib.Book{
		{Name: "The newest book", Writer: "Nick", State: 7, TotalNumber: 10},
		{Name: "The crazy book", Writer: "Big", State: 5, TotalNumber: 20},
		{Name: "The not so good book", Writer: "Very", State: 3, TotalNumber: 37},
		{Name: "The odd book", Writer: "Nice", State: 7, TotalNumber: 20},
		{Name: "The greatest book", Writer: "Guy", State: 10, TotalNumber: 15},
	}
}

package main

import (
	lib "Library/library"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the library!")

	library := lib.CreateLibrary()
	printLibrary := lib.PrintLibrary{Library: library}
	libraryHandler := lib.LibraryHandler{Library: library}
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

			selectedBook := library.GetBookById(selection)

			fmt.Println(printLibrary.GetRemainingBooksCount(selectedBook))

			libraryHandler.HandleRental(&user, selectedBook)
		case 2:
			libraryHandler.HandleNewBook()
		case 3:
			var selection uint16

			fmt.Println(printLibrary.GetAllOrders(&user))

			fmt.Scan(&selection)

			libraryHandler.HandleBookReturn(&user, selection)
		case 4:
			fmt.Println(printLibrary.GetAllOrders())
		}
	}
}

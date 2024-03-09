package main

import (
	lib "Library/library"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the library!")

	var library lib.Library

	library.Books = []lib.Book{
		{Name: "The newest book", Writer: "Nick", State: 7, TotalNumber: 10},
		{Name: "The crazy book", Writer: "Big", State: 5, TotalNumber: 20},
		{Name: "The not so good book", Writer: "Very", State: 3, TotalNumber: 37},
		{Name: "The odd book", Writer: "Nice", State: 7, TotalNumber: 20},
		{Name: "The greatest book", Writer: "Guy", State: 10, TotalNumber: 15},
	}

	fmt.Printf("We have a total of %d books in our library!", library.GetRemainingBooksCount())
}

/*
	The user can get books from the library
	The user can register new books in the library
	The user has his own account with books rented with dates on dates on when to give them back
	The book's potential cost of losing or not giving back in time is determined by its State
*/

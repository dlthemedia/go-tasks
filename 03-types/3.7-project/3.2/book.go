package main

import "strconv"

type Book struct {
	Title string

	Pages int
}

func Summary(book Book) string {
	return book.Title + " (" + strconv.Itoa(book.Pages) + " стр.)"
}

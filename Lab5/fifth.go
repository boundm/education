package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Book struct {
	Title  string
	Author string
	Year   int
}

func (b Book) String() string {
	return fmt.Sprintf("Книга: %s, Автор: %s, Год: %d", b.Title, b.Author, b.Year)
}

func main() {
	book := Book{Title: "1984", Author: "Джордж Оруелл", Year: 1949}

	fmt.Println(book.String())
}

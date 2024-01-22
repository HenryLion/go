package main

import "fmt"

func main() {
	type Book struct {
		id int
	}
	bookshelf := make(chan Book, 3)

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case bookshelf <- Book{id: 1}:
			fmt.Println("success to put book", i)
		default:
			fmt.Println("failed to put book")
		}
	}

	for i := 0; i < cap(bookshelf)*2; i++ {
		select {
		case book := <-bookshelf:
			fmt.Println("success to get book", book.id)
		default:
			fmt.Println("fail to get book")
		}
	}
}

package main

import "fmt"

type Aboutable interface {
	About() string
}

type Book struct {
	name string
}

type NaiCha struct {
	money int
}

func (book *Book) About() string {
	return "Book: " + book.name
}

func main() {
	var a Aboutable = &Book{"Go 101"}
	fmt.Println(a)

	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i)

	i = a
	fmt.Println(i)

}

package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

type Singer struct {
	Person
	works []string
}

func main() {
	type P = *bool
	type M = map[int]int
	var x struct {
		string
		//error
		*int
		P
		M
		http.Header
	}
	x.string = "Go"
	//x.error = nil // cause panic
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
	fmt.Println(x)

	//The main purpose of type embedding is to extend the functionalities of
	//the embedded types into the embedding type,
	//so that we don't need to re-implement the functionalities of
	//the embedded types for the embedding type.

	p := Person{"pencil", 23}
	p.PrintName()
	p.SetAge(3)
	fmt.Println(p)

	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName()
	gaga.Name = "Lady Gaga"
	gaga.SetAge(31)
	gaga.PrintName()
	fmt.Println(gaga.Age)
}

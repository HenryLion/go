package main

import "fmt"

type FilterFunc func(in int) bool

func (f FilterFunc) Filte(in int) bool {
	return f(in)
}

func FilterEven(in int) bool {
	return in%2 == 0
}

func main() {
	var filter FilterFunc = FilterEven
	fmt.Println(filter.Filte(3))
}

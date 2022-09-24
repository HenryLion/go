package main

import "fmt"

type person struct {
	name string
	age  float32
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 0.43
	return &p // return a pointer to local variable is ok,
	// the compiler will allocate the memory in the heap
}

func main() {
	fmt.Println(newPerson("xingxing"))
}

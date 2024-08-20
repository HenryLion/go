package main

import "fmt"

type IAnimal interface {
	Speak()
	eat()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	fmt.Printf("%s Woof!\n", d.name)
}

func (d *Dog) eat() {
	fmt.Printf("%s eat meat\n", d.name)
}

type MyDog struct {
	*Dog
}

func main() {
	var d IAnimal = MyDog{
		&Dog{"haskel"},
	}
	_, ok := d.(IAnimal)
	if ok {
		println("MyDog is IAnimal")
	} else {
		println("MyDog is not IAnimal")
	}
	d.eat()
	d.Speak()

	fmt.Println(1723538285 + 1800 - 1)
}

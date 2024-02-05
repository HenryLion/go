package main

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	//animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	//for _, animal := range animals {
	//	println(animal.Speak())
	//}

	var d Dog
	println(d.Speak())

	aDog := []Animal{Dog{}}
	for _, animal := range aDog {
		println(animal.Speak())
	}
}

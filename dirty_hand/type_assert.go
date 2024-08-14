package main

import "fmt"

type Animal2 interface {
	dogEat()
}

type Dogs struct {
}

func (d *Dogs) dogEat() {
	fmt.Println("dog eat banana")
}

type Tiger struct {
	*Dogs
}

func (t *Tiger) dogEat() {
	fmt.Println("tiger eat meat")
}

func main() {
	a := Tiger{}
	a.Dogs.dogEat()

	var a1 Animal2 = &Dogs{}
	a1.dogEat()

	var a2 Animal2 = &Tiger{}
	a2.dogEat()
}

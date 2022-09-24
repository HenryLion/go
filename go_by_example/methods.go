package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width += 2
	r.height += 2
	return r.width * r.height
}

func (r rect) perim() int {
	r.width += 3
	r.height += 3
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{2, 3}
	fmt.Println(r)
	fmt.Println(r.area()) // change r value
	fmt.Println(r)
	fmt.Println((&r).perim()) // not change r value
	fmt.Println(r)
}

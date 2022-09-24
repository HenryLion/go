package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = 3
	fmt.Println(d)
	fmt.Printf("%T\n", d) // int

	var e = true
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}

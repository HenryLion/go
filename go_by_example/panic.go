package main

import "fmt"

// when panic the call stack will print
func f2() {
	fmt.Println("call f2")
	panic("f2 panic")
}
func f1() {
	fmt.Println("call f1")
	f2()
}

func main() {
	f1()
}

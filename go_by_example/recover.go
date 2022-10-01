package main

import "fmt"

func mayPanic() {
	panic("a panic")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered Error:\n", r)
		}
	}()
	mayPanic()
	fmt.Println("after mayPanic")

}

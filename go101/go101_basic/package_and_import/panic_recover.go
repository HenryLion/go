package main

import "fmt"

// 1、without recover, the defer func will be execute after panic
// 2、recover() func should be called by defer

func main() {
	defer func() {
		fmt.Println("exit normally.")
	}()

	fmt.Println("hi!")
	defer func() {
		v := recover()
		fmt.Println("recovered:", v)
	}()
	panic("bye!")
	fmt.Println("unreachable")
}

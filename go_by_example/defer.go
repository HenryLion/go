package main

import "fmt"

func err_handler() {
	fmt.Println("handler error before return")
}

func f1() {
	fmt.Println("here we do some business")
	panic("orr le")
	defer err_handler() // after panic defer not run

}

func main() {
	f1()
}

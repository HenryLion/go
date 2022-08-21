package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(1000 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonaccci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		fmt.Printf("%s\n", "tianxia")

		time.Sleep(delay)
		// when main routine over, this "dixia" will not print
		fmt.Printf("%s\n", "dixia")
	}
}

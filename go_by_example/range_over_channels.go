package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// It's possible to close a non-empty channel
	close(queue)

	// After close, can still get the remaining values
	for elem := range queue {
		fmt.Println(elem)
	}
}

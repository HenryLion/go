package main

import "fmt"

func Printer(input <-chan uint64) {
	for {
		x, ok := <-input
		if !ok {
			break
		}
		fmt.Println(x)
	}
}

func main() {
	ch := make(chan uint64)
	go func() {
		for i := uint64(1); i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	Printer(ch)
}

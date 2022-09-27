package main

import (
	"fmt"
	"time"
)

// channels are the pipes that connect concurrent goroutines

func main() {
	message := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		message <- "ping" // after message return, the main and sub routine will run concurrent
		fmt.Println("go func")
		fmt.Println("go func1")
		fmt.Println("go func2")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("go func3")
		fmt.Println("go func4")

	}()

	fmt.Println("main routines")
	fmt.Println("1+1=", 1+1)
	msg := <-message // main routine will block here wait for message
	fmt.Println(msg)
	time.Sleep(time.Second)
}

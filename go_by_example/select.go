package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "write to channel one"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		c2 <- "write to channel two"
	}()

	// select will block when no val write to channel
	for i := 0; i < 2; i++ {
		select { // we will know which routine has returned
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string, 2)

	go func() {
		messages <- "give val 1"
	}()
	time.Sleep(20 * time.Millisecond)
	messages <- "give val 2"
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}

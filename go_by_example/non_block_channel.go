package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message:", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("receive message:", msg)
	default:
		fmt.Println("no message received")
	}

	select {
	case msg := <-messages:
		fmt.Println("receive message:", msg)
	case sig := <-signals:
		fmt.Println("receive signal", sig)
	default:
		fmt.Println("no activity")
	}
}

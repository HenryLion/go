package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func(c chan<- int, x int) {
		time.Sleep(time.Second)
		c <- x * x
	}(c, 3)
	c1 := make(chan struct{})
	go func(c <-chan int) {
		fmt.Println(<-c)
		time.Sleep(time.Second)
		c1 <- struct{}{}
	}(c)
	<-c1
	fmt.Println("bye 1")

	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Println(<-ball, "kicked the ball")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("John")
	go kickBall("Alice")
	go kickBall("Bob")
	go kickBall("Emily")
	ball <- "referee"
	var c2 chan bool
	<-c2
}

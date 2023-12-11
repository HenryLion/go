package main

import "fmt"

func printNumber2(w chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	w <- true
}

func printLetter2(w1, w2 chan bool) {
	<-w1 // 必须等到w1被写入了值才能执行，否则在此处等待
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Printf("%c ", i)
	}
	w2 <- false
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumber2(w1)
	go printLetter2(w1, w2)
	<-w2 // 主程等待w2被写入值，否则在此等待
}

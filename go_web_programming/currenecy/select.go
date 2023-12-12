package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	// 下面几行用法也是可以的，因为从chan读会block住，打印语句执行的时候肯定已经获取到值了
	//msg1, msg2 := <-a, <-b // block here until a,b has values
	//fmt.Printf("%s from A\n", msg1)
	//fmt.Printf("%s from B\n", msg2)
	for i := 0; i < 5; i++ {
		select {
		case msg := <-a:
			fmt.Printf("%s from A\n", msg)
		case msg := <-b:
			fmt.Printf("%s from B\n", msg)
		default:
			time.Sleep(10 * time.Millisecond)
			fmt.Println("Default")
		}
	}

}

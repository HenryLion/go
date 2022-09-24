package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main rountine begin")
	childdone := make(chan int)
	go func() {
		fmt.Println("clildroutine begin")
		time.Sleep(1 * time.Second)
		fmt.Println("clildroutine over")
		childdone <- 10
	}()
	time.Sleep(500 * time.Millisecond)
	var a int = <-childdone
	fmt.Printf("main routine over %d\n", a)
	close(childdone)
}

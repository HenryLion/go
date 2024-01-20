package main

import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready
	println("worker", id, "starts to work")
	time.Sleep(time.Second * time.Duration(id+1))
	println("worker", id, "has done its work")
	done <- T{}
}

func main() {
	ready, done := make(chan T), make(chan T)
	for i := 0; i < 3; i++ {
		go worker(i, ready, done)
	}

	time.Sleep(time.Second * 3 / 2)
	// 可以在main routine给ready赋值，使得三个子routine可以running
	//ready <- T{}
	//ready <- T{}
	//ready <- T{}

	// 也可以在main routine直接close ready，三个子routine都可以拿到值接触block
	close(ready) // infinite values can be received from a closed channel
	for i := 0; i < 3; i++ {
		<-done
	}
	println("main has notified")
}

package main

import "fmt"

func main() {
	mutex := make(chan struct{}, 1)
	counter := 0
	increase := func() {
		mutex <- struct{}{}
		counter++
		<-mutex
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{}) // done纯粹是为了等待子goroutine执行完成
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter)
}

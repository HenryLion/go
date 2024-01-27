package main

import "sync"

var wg1 sync.WaitGroup

func f3() {
	var a, b int
	var c = make(chan bool)

	go func() {
		defer wg1.Done()
		a = 1
		c <- true
		if b != 1 {
			panic("b != 1")
		}
	}()

	go func() {
		defer wg1.Done()
		b = 1
		<-c
		if a != 1 {
			panic("a != 1")
		}
	}()
}

func main() {
	wg1.Add(2)
	f3()
	wg1.Wait()
}

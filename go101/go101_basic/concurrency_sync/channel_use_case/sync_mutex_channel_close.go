package main

import (
	"fmt"
	"sync"
)

type MyChannel1 struct {
	c      chan int
	closed bool
	mutex  sync.Mutex
}

func NewMyChannel1() *MyChannel1 {
	return &MyChannel1{c: make(chan int), closed: false}
}

func (mc *MyChannel1) SafeClose(wg *sync.WaitGroup) {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

	if !mc.closed {
		close(mc.c)
		fmt.Println("close channel")
		mc.closed = true
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	mc := NewMyChannel1()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go mc.SafeClose(&wg)
	}
	wg.Wait()
	fmt.Println("main end!")
}

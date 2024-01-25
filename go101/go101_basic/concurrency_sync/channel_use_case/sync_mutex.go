package main

import (
	"fmt"
	"sync"
)

// sync.Mutex用来保证同一时刻只有一个goroutine可以访问共享内容

func main() {
	var wg sync.WaitGroup
	var s sync.Mutex
	count := uint64(0)

	increment := func() {
		defer wg.Done()
		s.Lock()
		count++
		s.Unlock()
	}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait() // 等待所有 goroutine 执行完毕
	fmt.Println(count)
}

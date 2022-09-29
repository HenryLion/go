package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 10000; c++ {
				atomic.AddUint64(&ops, 1)
				//ops += 1
			}
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}

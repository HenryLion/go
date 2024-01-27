package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const Max = 100000
	const NumReceivers = 100
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers + 1)

	dataCh := make(chan int)

	count := uint32(0)

	go func() {
		defer wgReceivers.Done()
		for {
			if value := rand.Intn(Max); value == 0 {
				log.Println("get a zero!!", count)
				close(dataCh)
				return
			} else {
				count++
				dataCh <- value
			}
		}
	}()

	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			for value := range dataCh { // channel关闭后,会自动退出for循环
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()
}

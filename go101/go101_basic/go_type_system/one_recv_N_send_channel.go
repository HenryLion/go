package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func getLineNumber() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// let the receiver close an additional signal channel
// to notify senders to stop sending values

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const Max = 100000
	const NumSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int)
	stopCh := make(chan struct{})

	for i := 0; i < NumSenders; i++ {
		id := i
		go func() {
			for {
				select {
				case <-stopCh:
					log.Println("have exit", id, getLineNumber())
					return
				default:

				}
				select {
				case <-stopCh:
					log.Println("have exit", id, getLineNumber())
					return
				case dataCh <- rand.Intn(10000):
				default:
				}
			}
		}()
	}
	count := 0
	go func() {
		defer wgReceivers.Done()
		for value := range dataCh {
			if value == 13 {
				log.Println("get value: ", value)
				close(stopCh)
				return
			}
			count++
			log.Println(value)
		}
	}()

	wgReceivers.Wait()

	log.Println("zongshu: ", count)
}

package main

import (
	"fmt"
	"time"
)

type Request interface {
}

func handle(r Request) {
	fmt.Println(r.(int))
}

const (
	RateLimitPeriod = time.Minute
	RateLimit       = 200
)

func handleRequests(requests <-chan Request) {
	quotas := make(chan time.Time, RateLimit)
	go func() {
		tick := time.NewTicker(RateLimitPeriod / RateLimit)
		defer tick.Stop()
		for t := range tick.C { // for 用于不断的从tick.C这个channel中读取数据
			select {
			case quotas <- t:
			default:
			}
		}
	}()

	for r := range requests {
		<-quotas
		go handle(r)
	}
}

func main() {
	requests := make(chan Request)
	go handleRequests(requests)
	for i := 0; ; i++ {
		requests <- i
	}
}

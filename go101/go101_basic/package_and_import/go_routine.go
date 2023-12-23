package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait()
}

package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("customer#", c, " enters the bar")
	seat := <-bar
	log.Print("++ customer#", c, " drinks at seat#", seat)
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- customer#", c, "frees seat#", seat)
	bar <- seat
}

func main() {
	rand.Seed(time.Now().UnixNano())
	bar24x7 := make(Bar, 10)
	for i := 0; i < cap(bar24x7); i++ {
		bar24x7 <- Seat(i)
	}
	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)

		// 此处会造成在程序生命周期有无数的goroutine
		go bar24x7.ServeCustomer(customerId)
	}

	for {
		time.Sleep(time.Second)
	}
}

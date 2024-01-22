package main

import (
	"errors"
	"time"
)

func doRequest(c chan<- int) {
	time.Sleep(3 * time.Second)
	c <- 66
}

func requestWithTimeout(timeout time.Duration) (int, error) {
	c := make(chan int)
	go doRequest(c)
	select {
	case data := <-c:
		return data, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	data, err := requestWithTimeout(4 * time.Second)
	if err != nil {
		println(err.Error())
	} else {
		println(data)
	}
}

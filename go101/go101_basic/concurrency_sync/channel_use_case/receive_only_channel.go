package main

import (
	"fmt"
	"math/rand"
	"time"
)

// <-chan 表示一个只能读取数据的 channel，
// 也就是说你只能从这个channel中读取数据，不能向这个 channel 中写入数据。
func longTimeRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()
	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}

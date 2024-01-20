package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 次函数的参数r是一个只写channel,
// 也就是说函数实现只能给这个channel写数据,不能从这个channel读数据
func longTimeRequest1(r chan<- int32) {
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares1(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ra, rb := make(chan int32), make(chan int32)
	go longTimeRequest1(ra)
	go longTimeRequest1(rb)
	fmt.Println(sumSquares1(<-ra, <-rb))
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Once的一个实例调用其Do方法只能执行一次，多次调用Do方法也不会执行多次
// 即便多次调用Do方法执行的函数不同，也只会执行第一次调用Do方法时执行的函数
var once sync.Once

func sayHello() {
	fmt.Println("hello, world")
}

func sayHelloV2() {
	fmt.Println("hello, world 2")
}

func main() {
	ch := make(chan int)
	close(ch)
	fmt.Println("channel closed")
	//close(ch) // will panic here
	//fmt.Println("close channel again")
	go once.Do(sayHello)
	go once.Do(sayHelloV2) // sayHelloV2不会执行

	time.Sleep(time.Second * 2)

}

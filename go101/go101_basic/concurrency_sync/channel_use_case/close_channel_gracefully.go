package main

import (
	"fmt"
	"time"
)

type T2 int

// IsClosed 此函数并不能区分通道中原本就有数据还是通道已经被关闭的情况
func IsClosed(ch <-chan T2) bool {
	select {
	case <-ch:
		fmt.Println("1")
		return true
	default:
	}
	fmt.Println("2")
	return false
}

func main() {
	c := make(chan T2)

	go func() {
		c <- T2(3)
	}()
	time.Sleep(time.Second)
	fmt.Println(IsClosed(c)) // 这里输出true，但是通道并没有关闭
	close(c)
	fmt.Println(IsClosed(c))
}

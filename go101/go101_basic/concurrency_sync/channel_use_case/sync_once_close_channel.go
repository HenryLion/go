package main

import (
	"fmt"
	"sync"
)

// 由于多次关闭一个channel会导致panic
// 所以这里使用了sync.Once来保证只关闭一次

type MyChannel struct {
	c    chan int
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{
		c: make(chan int),
	}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		fmt.Println("will safe close the channel")
		close(mc.c)
	})
}

func main() {
	mc := NewMyChannel()
	mc.SafeClose()
	mc.SafeClose()
}

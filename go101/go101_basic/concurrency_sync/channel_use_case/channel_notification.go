package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

// use channels for notifications
func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()
	<-done // 要读数据，但是没数据可读，所以block在这里等待有数据写入
	fmt.Println(values[0], values[len(values)-1])

	fmt.Println("=============")
	done1 := make(chan struct{})
	go func() {
		fmt.Println("Hello")
		time.Sleep(time.Second * 2)
		<-done1
	}()

	// 主goroutine想要发送数据，但是对于unbuffered channel来说，
	// 只有当有其他goroutine在读取数据时，才能发送数据，所以主goroutine会block在这里
	done1 <- struct{}{}
	fmt.Println("World")
}

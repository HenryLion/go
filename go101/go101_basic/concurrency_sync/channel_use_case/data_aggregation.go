package main

import (
	"fmt"
	"sync"
	"time"
)

// Aggregator 将入参channels中的数据聚合到一个输出通道中
func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(in <-chan uint64) {
			defer wg.Done()
			for x := range in {
				out <- x
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {

	in1 := make(chan uint64)
	in2 := make(chan uint64)

	// 使用 Aggregator 函数将输入通道聚合到一个输出通道
	out := Aggregator(in1, in2)

	// 启动两个 goroutine 分别向输入通道发送数据
	go func() {
		for i := uint64(0); i < 5; i++ {
			in1 <- i
			time.Sleep(1000 * time.Millisecond)
		}
		close(in1) // 发送完数据后关闭通道
	}()

	go func() {
		for i := uint64(10); i < 15; i++ {
			in2 <- i
			time.Sleep(2000 * time.Millisecond)
		}
		close(in2) // 发送完数据后关闭通道
	}()

	// 从输出通道接收并打印数据
	for num := range out {
		fmt.Println(num)
	}
}

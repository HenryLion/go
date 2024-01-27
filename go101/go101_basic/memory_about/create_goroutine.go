package main

import (
	"fmt"
	"sync"
	"time"
)

// 本程序想说明go routine创建之前的语句是先于go routine创建执行的

var x, y int
var wg sync.WaitGroup

var x1, y1 int

func f1() {
	x, y = 123, 789 // 这句赋值肯定要早于下面goroutine的执行
	go func() {
		fmt.Println(x) // 123 这句打印肯定要早于下面goroutine的执行
		wg.Done()
		go func() {
			fmt.Println(y) // 789
			wg.Done()
		}()
	}()
}

func f2() {
	go func() {
		fmt.Println(x1)
		wg.Done()
	}()
	go func() {
		fmt.Println(y1)
		wg.Done()
	}()
	time.Sleep(time.Millisecond * 2)
	x1, y1 = 123, 789
}
func main() {
	wg.Add(4)
	f1()
	f2()
	wg.Wait()

}

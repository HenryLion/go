package main

import (
	"fmt"
	"runtime"
	"time"
)

func f0() int {
	var x = 1
	defer fmt.Println("exits normally:", x)
	x++
	return x
}

func f1() {
	var x = 1
	defer fmt.Println("exits normally:", x)
	x++
}

func f2() {
	var x, y = 1, 0
	defer fmt.Println("exits for panicking:", x)
	x = x / y
	x++
}

func f3() int {
	x := 1
	defer fmt.Println("exits for Goexiting:", x)
	x++
	runtime.Goexit()
	return x + x
}
func f4() {
	panic(4)
}

func main() {
	f0()
	f1()
	//f2()
	go f3()
	time.Sleep(time.Second)

	fmt.Println("===============")
	defer func() {
		fmt.Println(recover())
	}()

	defer panic(3)
	defer panic(2)
	//panic(0)

	go f4() // 会导致主携程crash

	time.Sleep(time.Second)
	fmt.Println("after f4 goroutine panic")
}

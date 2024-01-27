package main

import (
	"log"
	"runtime"
)

// 关于编译器对指令顺序的优化

var a string
var done bool

func setup() {
	a = "hello, world" // 这两行赋值语句的顺序可能被编译器优化
	done = true        // 哪条语句先执行可以认为是undefined
	if done {
		log.Println(len(a))
	}
}

// 使用go build -race编译程序并执行可执行文件，会自动检测是否发生了data race

func main() {
	go setup()
	for !done { // 在循环中等待done变量为true
		runtime.Gosched() // 让出当前goroutine的执行，以便其他goroutine可运行
	}
	log.Println(a)
}

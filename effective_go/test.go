package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

//需要注意的是，defer 延迟执行的函数的参数在 defer 语句被执行时就被确定了。这意味着当 defer 带有函数调用时，函数的结果会被缓存，并在实际执行时使用缓存的结果
//entering: b
//in b
//entering: a
//in a
//leaving: a
//leaving: b

func main() {
	b()
}

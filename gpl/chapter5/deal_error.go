package main
import (
	"fmt"
	"strings"
)

func add1(r rune) rune {return r + 1}

// anonymous function usage: function return a function
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main () {
	var zi rune = rune('中') // 20013
	fmt.Println(int(zi)) // 20013
	fmt.Println(string(zi)) // 中
	fmt.Println(strings.Map(add1,"中国")) // 丮图

	// anonymous function
	fmt.Println(strings.Map(func(r rune) rune { return r + 1}, "中国"))

	f := squares()
	fmt.Println(f())
	fmt.Println(f())

	var ops []func()int // ops is a slice of funcs
	for i := 0; i < 3; i++ {
		ops = append(ops, func()int {return i})
	}

	// NOTE: the ops all return 3.
	for _, op := range ops {
		fmt.Println(op ()) // 3 \n 3 \n 3
	}
}
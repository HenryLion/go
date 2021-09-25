package main

import "fmt"
import "os"

// 可变参数函数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func myerrorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main () {
	fmt.Println(sum()) // 0
	fmt.Println(sum(3)) // 3
	fmt.Println(sum(1,2,3,4)) // 10

	// test myerror
	linenum , name := 12, "count"
	myerrorf(linenum, "undefined: %s", name) // Line 12: undefined: count
}
package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("binary: %b\n", 123)
	fmt.Printf("char: %c\n", 98)
	fmt.Printf("hex: %x\n", 230567)
	fmt.Printf("float1: %f\n", 45.6)
	fmt.Printf("float2: %e\n", 123456789.0)
	fmt.Printf("float2: %E\n", 123456789.0)
	fmt.Printf("str1: %s\n", "\"nageren\"")
	fmt.Printf("str2: %q\n", "\"shenmewanyi\"")
	fmt.Printf("str3: %x\n", "hex this")
	fmt.Printf("pointer: %p\n", &p)
	fmt.Printf("width1: |%-6d|%-6d|\n", 12, 345)
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.456)
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.345)

}

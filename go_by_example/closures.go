package main

import "fmt"

func intSeq() func() int {
	fmt.Println("call intSeq")
	i := 0
	j := 3
	return func() int {
		fmt.Printf("call inner func i = %d\n", i)
		i++
		return i + j
	}
}

func main() {
	netxInt := intSeq()
	fmt.Println(netxInt())
	fmt.Println(netxInt())
	fmt.Println(netxInt())
}

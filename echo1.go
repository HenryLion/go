package main

import (
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercise1_1() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercise1_2() {
	for index, arg := range os.Args[:] {
		fmt.Println(index, arg)
	}
}

func main() {
	echo1()
	echo2()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println("exercise below =====")
	exercise1_1()
	exercise1_2()
}

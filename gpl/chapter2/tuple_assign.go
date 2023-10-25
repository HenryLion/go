package main

import (
	"fmt"
	"strconv"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	fmt.Println(gcd(21, 7))
	fmt.Println(fib(5))

	nameAge := make(map[string]int)
	nameAge["hanbin"] = 33
	nameAge["zhangtao"] = 29

	// when get a map value, can this style
	v, ok := nameAge["hanbin"]
	if ok {
		fmt.Println(v, " ok")
	} else {
		fmt.Println("not ok")
	}

	// when get a map value, also can this style
	val := nameAge["zhangtao"]
	fmt.Println(val)

	// test if int can convert to string
	age := 3000
	fmt.Println(string(age)) // this is wrong

	strAge := strconv.Itoa(age) // this is right
	fmt.Println(strAge)

	var a = b + c
	var b, c = 1, 2
	fmt.Println(a)

}

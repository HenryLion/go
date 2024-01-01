package main

import "fmt"

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n)
}

func main() {
	fmt.Printf("%T\n", Double)
	var f func(n int) int
	f = Double
	g := Apply
	fmt.Printf("%T\n", g)

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12

	// function returns a function (a closure)
	isMultipleOfX := func(x int) func(int) bool {
		return func(n int) bool {
			return n%x == 0
		}
	}
	var isMultipleOf3 = isMultipleOfX(3)
	fmt.Println(isMultipleOf3(6))

	var isMultipleOf5 = isMultipleOfX(5)
	isMultipleOf15 := func(n int) bool {
		return isMultipleOf3(n) && isMultipleOf5(n)
	}
	fmt.Println(isMultipleOf15(60))

}

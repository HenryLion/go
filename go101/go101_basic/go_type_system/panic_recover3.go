package main

import "fmt"

// use panic recover calls to simulate long jump statement

func main() {
	n := func() (result int) {
		defer func() {
			if v := recover(); v != nil {
				if n, ok := v.(int); ok {
					result = n
				}
			}
		}()

		func() {
			fmt.Println("func 1")
			func() {
				fmt.Println("func 2")
				func() {
					fmt.Println("func 3")
					panic(123)
				}()
			}()
		}()
		fmt.Println("never here!")
		return 0
	}()
	fmt.Println(n)
}

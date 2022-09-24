package main

import "fmt"

// variadic params can be use like a slice
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func sum_str(strs ...string) {
	t_str := ""
	for _, str := range strs {
		t_str += str + " "
	}
	fmt.Println(t_str)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
	sum_str("hello", "xingxing")
	strs := []string{"xing", "zai", "yu", "le", "quan"}
	sum_str(strs...)
}

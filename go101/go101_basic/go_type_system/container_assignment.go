package main

import "fmt"

func main() {
	m0 := map[int]int{0: 7, 1: 8, 2: 9}
	m1 := m0
	m1[0] = 2
	fmt.Println(m0, m1) // map[0:2 1:8 2:9] map[0:2 1:8 2:9]

	s0 := []int{7, 8, 9}
	s1 := s0 // s0 and s1 share the same memory segment
	s1[0] = 2
	fmt.Println(s0, s1) // [2 8 9] [2 8 9]

	s1 = append(s1, 10) // s0 and s1 no longer share the same memory segment
	fmt.Println(s0, s1)
	s1[0] = 11 // s0 will not change
	fmt.Println(s0, s1)

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println(a0, a1) // [7 8 9] [2 8 9]

	// []string(nil) 表示类型转换
	// []string{} 表示创建一个空切片
	var s = append([]string(nil), "array", "slice")
	fmt.Println(s)      // [array slice]
	fmt.Println(cap(s)) // 2
}

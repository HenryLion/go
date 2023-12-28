package main

import "fmt"

func main() {
	s0 := []int{1, 2, 3, 4, 5, 6, 7}
	s1 := s0[0:3]
	s1[1] = 22
	s1 = append(s1, 9)
	fmt.Println(s0, s1) // [1 22 3 9 5 6 7] [1 22 3 9] the result astonish me

	s1 = append(s1, 10)
	fmt.Println(s0, s1) // [1 22 3 9 10 6 7] [1 22 3 9 10]

	s1 = append(s1, 11, 12)
	fmt.Println(s0, s1)
	// 下面打印的地址不一样，说明两个slice的地址不一样，但是他们数据指针指向的地址应该是共享的
	fmt.Printf("%p, %p\n", &s0, &s1)

	s1 = append(s1, 12, 13, 14, 15, 16)
	fmt.Println(s0, s1)

	s1[0] = 33
	fmt.Println(s0, s1)

	for k, v := range s1 {
		fmt.Println(k, v)
	}
}

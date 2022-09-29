package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("hello world, my baby xingxing.")

	go func() {
		s1 := make([]int, 0)
		for i := 0; i < 15; i++ {
			s1 = append(s1, rand.Intn(5))
		}
		fmt.Println("s1", s1)
	}()

	go func() {
		s2 := make([]int, 0)
		for i := 0; i < 15; i++ {
			s2 = append(s2, rand.Intn(5))
		}
		fmt.Println("s2", s2)
	}()

	time.Sleep(time.Second)
	var state = make(map[int]int)
	fmt.Println(state[3]) // 0

}

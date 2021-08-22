package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pc [256]byte

var numbers []uint64

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	for i := 0; i < 10000000; i++ {
		numbers = append(numbers, rand.Uint64())
	}
}

// PopCount ...
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop ... exercise 2.3
func PopCountLoop(x uint64) int {
	var rst int
	for i := 0; i < 8; i++ {
		rst += int(pc[byte(x>>(i*8))])
	}
	return rst
}

// PopCountShift ... exercise 2.4
func PopCountShift(x uint64) int {
	var rst int
	for i := 0; i < 64; i++ {
		if (x>>i)&1 == 1 {
			rst++
		}
	}
	return rst
}

// PopCountRightmost ... Exercise 2.5
func PopCountRightmost(x uint64) int {
	var rst int
	for x != 0 {
		rst++
		x = x & (x - 1)
	}
	return rst
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	start := time.Now()
	for _, num := range numbers {
		PopCount(num)
	}
	secs := time.Since(start).Seconds()
	fmt.Println("PopCount: ", secs)

	start = time.Now()
	for _, num := range numbers {
		PopCountLoop(num)
	}
	secs = time.Since(start).Seconds()
	fmt.Println("PopCountLoop: ", secs)

	start = time.Now()
	for _, num := range numbers {
		PopCountShift(num)
	}
	secs = time.Since(start).Seconds()
	fmt.Println("PopCountShift: ", secs)

	start = time.Now()
	for _, num := range numbers {
		PopCountRightmost(num)
	}
	secs = time.Since(start).Seconds()
	fmt.Println("PopCountRightmost: ", secs)

}

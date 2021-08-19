package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var numbers []uint32
	for i := 0; i < 1000000000; i++ {
		numbers = append(numbers, rand.Uint32())
	}

	start := time.Now()
	sumCh := make(chan uint64)
	for j := 0; j < 10; j++ {
		go getSum(numbers[j*100000000:(j+1)*100000000], sumCh)
	}
	var fastSum uint64
	for j := 0; j < 10; j++ {
		fastSum += <-sumCh
	}
	secs := time.Since(start).Seconds()
	fmt.Println(fastSum+1, secs)

	start = time.Now()
	var sum uint64
	for _, val := range numbers {
		sum += uint64(val)
	}
	secs = time.Since(start).Seconds()
	fmt.Println(sum, secs)
}

func getSum(numbers []uint32, sum chan<- uint64) {
	var temp uint64
	for _, val := range numbers {
		temp += uint64(val)
	}
	sum <- temp
}

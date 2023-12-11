package main

import (
	"fmt"
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
	fmt.Println()
}

func goPrint1() {
	fmt.Println("diyige") // 这句是主程执行
	go printNumbers1()
	go printLetters1()
	fmt.Println("dierge") // 这句是主程执行
}

func main() {
	print1()
	goPrint1()
	time.Sleep(1 * time.Second)
}

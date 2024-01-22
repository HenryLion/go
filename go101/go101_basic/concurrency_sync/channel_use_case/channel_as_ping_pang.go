package main

import (
	"fmt"
	"os"
	"time"
)

type Ball uint64

func Play(playerName string, table chan Ball) {
	var lastValue Ball = 1
	for {
		ball := <-table
		fmt.Println(playerName, ball)
		ball += lastValue
		if ball < lastValue {
			os.Exit(0)
		}
		lastValue = ball
		table <- ball
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball)
	go func() {
		// 如果不在go func()中给table赋值，
		// 那么主goroutine会阻塞在此处，
		// 没有机会执行下面的go Play函数
		table <- 1
	}()
	go Play("A: ", table)
	Play("B: ", table)
}

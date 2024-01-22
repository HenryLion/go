package main

import (
	"fmt"
	"os"
	"time"
)

// 从nil值的channel发送或读取会被阻塞。

type Ball1 uint8

func Play1(playerName string, table chan Ball1, serve bool) {
	var receive, send chan Ball1
	if serve {
		receive, send = nil, table
	} else {
		receive, send = table, nil
	}
	var lastValue Ball1 = 1
	for {
		select {
		case send <- lastValue:
		case value := <-receive:
			fmt.Println(playerName, value)
			value += lastValue
			if value < lastValue {
				os.Exit(1)
			}
			lastValue = value
		}
		receive, send = send, receive
		time.Sleep(time.Second)
	}
}

func main() {
	table := make(chan Ball1)
	go Play1("A:", table, false)
	Play1("B:", table, true)
}

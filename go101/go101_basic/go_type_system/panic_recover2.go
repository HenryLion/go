package main

import (
	"log"
	"time"
)

// Automatically Restart a Crashed Goroutine

func shouldNotExist() {
	for {
		time.Sleep(time.Second)
		if time.Now().UnixNano()&0x3 == 0 {
			panic("unexpected situation")
		}
	}
}

func NeverExit(name string, f func()) {
	defer func() {
		if v := recover(); v != nil {
			log.Println(name, "is crashed, Restart it now.")
			go NeverExit(name, f) // 如果没有这一句，那么程序会陷入deadlock
		}
	}()
	f()
}

func main() {
	log.SetFlags(0)
	go NeverExit("job#A", shouldNotExist)
	go NeverExit("job#B", shouldNotExist)
	select {}
}

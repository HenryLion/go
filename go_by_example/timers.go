package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	select {
	case <-timer1.C:
		fmt.Println("Time 1 fired")
	}

	timer2 := time.NewTimer(time.Second)
	num := 3
	go func() { // inner func can access num
		<-timer2.C
		fmt.Printf("Time 2 fired, %d\n", num)
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 Stopped")
	}

	time.Sleep(3 * time.Second)
}

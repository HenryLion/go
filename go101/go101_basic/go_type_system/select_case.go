package main

import "fmt"

func main() {
	c := make(chan int)
	select {
	case <-c: // 由于c不带缓冲，所以此处读操作和下面的写操作相互阻塞了，会执行default
		fmt.Println("read from c")
	case c <- 1:
		fmt.Println("write to c")
	default:
		fmt.Println("Go here.")
	}

	// try send and try receive
	c1 := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c1 <- v:
		default:
			fmt.Println("send default branch")
		}
	}

	tryReceive := func() string {
		select {
		case v := <-c1:
			return v
		default:
			return "-"
		}
	}

	trySend("Hello!")
	trySend("Hi!")
	trySend("Bye!")

	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
	fmt.Println(tryReceive())
}

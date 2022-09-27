package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// work goroutine
	go func() {
		for {
			j, more := <-jobs // more will be false when jobs closed
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("send all jobs")

	<-done
}

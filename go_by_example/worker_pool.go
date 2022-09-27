package main

import (
	"fmt"
	"time"
)

func worker1(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs { // block for send to jobs
		fmt.Println("worker ", id, " started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finish job", j)
		result <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//for w := 1; w <= 3; w++ {
	//	go worker1(w, jobs, results)
	//}

	// the follow will be easy to understand
	go worker1(1, jobs, results)
	go worker1(2, jobs, results)
	go worker1(3, jobs, results)

	time.Sleep(time.Second)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

package main

import "fmt"
import "time"
import "flag"

var period = flag.Duration("p", 2*time.Second, "sleep peroid")

func main () {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
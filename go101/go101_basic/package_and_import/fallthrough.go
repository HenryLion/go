package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	switch n := rng.Intn(100) % 5; n {
	case 0, 1, 2, 3, 4:
		fmt.Println("n =", n)
		// The "fallthrough" statement makes the
		// execution slip into the next branch.
		fallthrough
	case 5, 6, 7, 8:
		// A new declared variable also called "n",
		// it is only visible in the current
		// branch code block.
		n := 99
		fmt.Println("n =", n) // 99
		fallthrough
	default:
		// This "n" is the switch expression "n".
		fmt.Println("n =", n)
	}
}

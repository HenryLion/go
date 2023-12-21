package main // specifies the name of the package the source file belong

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Go has", 25, "keywords.")
	fmt.Printf("Next pseudo-random number is %v.\n", rand.Uint32())
}

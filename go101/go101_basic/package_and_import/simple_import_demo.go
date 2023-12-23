package main // specifies the name of the package the source file belong

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Go has", 25, "keywords.")
	var num = rand.Uint32()
	fmt.Printf("Next pseudo-random number is %v, %x, %x.\n", num, num, "hello")
}

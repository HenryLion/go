package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("now known")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
		fmt.Println("You can out play")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) bool {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
		return true
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Println(whatAmI("4"))
}

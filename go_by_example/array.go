package main

import "fmt"
import "regexp"

func pureNumber(s string) bool {
	pattern := "^[0-9]+$"
	match, _ := regexp.MatchString(pattern, s)
	return match
}

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 0}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	fmt.Println(pureNumber("3456789."))
	fmt.Println(pureNumber("a234235366"))
}

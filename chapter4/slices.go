package main
import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) { // x 还有空间
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2 * len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main () {
	// slices 1
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// slices 2
	s1 := make([]int, 2) // make a slice
	for _, v := range s1 {
		fmt.Println(v)
	}
	fmt.Println(len(s1),cap(s1))
	s1 = append(s1, 4)
	fmt.Println(len(s1),cap(s1))
	s1 = append(s1, 5)
	fmt.Println(len(s1),cap(s1))
	s1 = append(s1, 6)
	fmt.Println(len(s1),cap(s1))
}
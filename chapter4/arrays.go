package main
import (
	"fmt"
	"crypto/sha256"
)

func getSha256 (str string) [32]uint8{
	return sha256.Sum256([]byte(str))
}

func changeArr3 (a *[3]int) {
	a[1] = 999
}


// PopCountRightmost ... Exercise 2.5
func PopCountRightmost(x byte) int {
	var rst int
	for x != 0 {
		rst++
		x = x & (x - 1)
	}
	return rst
}


// Exercise 4.1
func getBits (arrs [32]byte) int {
	var rst int
	for _, v := range arrs {
		rst += PopCountRightmost(v)
	}
	return rst
}



func main () {
	var a [3]int
	fmt.Println(a[0]) // default 0
	fmt.Println(a[len(a)-1])
	// fmt.Println(a[len(a)]) // out of bounds

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	var q [3]int = [3]int{1,2,3}
	for i, v := range q {
		fmt.Printf("%d-%d\n", i,v)
	}

	weekday := [...]int{0,1,2,3,4,5,6}
	fmt.Printf("len = %d, %d\n", len(weekday),weekday[3])

	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	// arrays can be initia with index and value
	symbol := [...]string{EUR:"ouyuan", USD:"meiyuan", RMB:"renminbi"}
	for _, v := range symbol {
		fmt.Printf ("%s  ", v)
	}

	fmt.Printf("\n%x\n", getSha256("x"))

	fmt.Printf("%x\n", getSha256("x"))

	arr := [...]int{5,6,7}

	fmt.Println(arr[1])

	changeArr3(&arr)
	fmt.Println(arr[1])

	fmt.Println(getBits(getSha256("x")))
	fmt.Println(getBits(getSha256("X")))
}
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type info struct {
	name string
	age  int
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func addSlice(x *[]int) {
	*x = append(*x, 123)
}

func returnSlice() []info {
	tempinfo := make([]info, 0)
	tempinfo = append(tempinfo, info{"pencil", 23})
	tempinfo = append(tempinfo, info{"zt", 21})
	return tempinfo
}

func main() {
	s := make([]string, 3)
	fmt.Println(s)
	s[0] = "bin"
	s[1] = "tao"
	s[2] = "xing"
	fmt.Println(s)

	s = append(s, "yun")
	fmt.Println(s)
	s = append(s, "yun", "qi")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
	fmt.Printf("%p - %p\n", &s, &c) // the address is different

	c1 := c
	fmt.Printf("%p - %p\n", &c, &c1) // the address is different

	t := []string{"han", "xing", "yun"}
	fmt.Println(t)

	if strings.Contains("abcdef", "") {
		fmt.Println("contain")
	} else {
		fmt.Println("not contain")
	}

	var numbers []int
	printSlice(numbers)

	addSlice(&numbers)
	printSlice(numbers)

	myslice := returnSlice()
	fmt.Println("-------")
	fmt.Println(len(myslice))
	for _, v := range myslice {
		fmt.Println(v)
	}
	num := 123
	mystr := "hello " + strconv.Itoa(num)
	fmt.Println(mystr)

	contents := make([][]string, 0)
	content1 := []string{"a", "b", "c"}
	content2 := []string{"1", "2", "3"}
	contents = append(contents, content1)
	contents = append(contents, content2)
	fmt.Println(contents)
}

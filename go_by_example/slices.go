package main

import "fmt"

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
}

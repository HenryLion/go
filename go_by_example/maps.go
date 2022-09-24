package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 23
	fmt.Println("map:", m)

	v1, _ := m["k1"]
	fmt.Println(v1)
	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println("map:", m)
	v2, ok := m["k2"]
	if ok {
		fmt.Println(v2)
	} else {
		fmt.Println("k2 not exists")

	}
	fmt.Println(m["k2"]) // 0
	fmt.Println(len(m))

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)

}

package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	s := make([]string, 4)
	s = append(s, "han")
	for i, v := range s {
		fmt.Println(i, v)
	}

	//kvs := map[string]string{"a": "apple", "b": "banana"}
	kvi := map[string]string{}
	for k, v := range kvi {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "hello,world" {
		fmt.Println(i, c)
	}
}

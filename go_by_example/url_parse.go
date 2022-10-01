package main

import (
	"fmt"
	"net/url"
)

func main() {
	s := "https://gobyexample.com/url-parsing"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
}

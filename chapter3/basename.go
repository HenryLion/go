package main

import (
	"fmt"
	"strings"
)

func basename (s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename_use_lib (s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main () {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))

	fmt.Println(basename_use_lib("a/b/c.go"))
	fmt.Println(basename_use_lib("c.d.go"))
	fmt.Println(basename_use_lib("abc"))

}

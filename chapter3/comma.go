package main

import (
	"fmt"
	"bytes"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteRune('[') // also can use buf.WriteByte();
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// exercise 3.10
func comm_non_recursive (s string) string {
	prefix := len(s) % 3
	var buf bytes.Buffer
	buf.WriteString(s[0:prefix])

	isbegin := false
	left := s[prefix:]
	if len(left) == len(s) {
		isbegin = true
	}
	for i := 0; i < len(left) / 3; i++ {
		if !isbegin {
			buf.WriteByte(',')
		}
		buf.WriteString(left[i*3:i*3+3])
		isbegin = false
	}
	return buf.String()
}

// exercise 3.11 comma deal with optional sign
func comma_deal_sign(s string) string {
	n := len(s)
	if s[0] == '-' {
		if n <= 4{
			return s
		}
	} else if n <= 3 {
		return s
	}
	return comma_deal_sign(s[:n-3]) + "," + s[n-3:]
}

// exercise 3.12
func string_anagrams (s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if !strings.Contains(s2, s1[i:i+1]) {
			return false
		}
	}

	for i := 0; i < len(s2); i++ {
		if !strings.Contains(s1, s2[i:i+1]) {
			return false
		}
	}
	return true
}




func main () {
	//fmt.Println(comma("123445866"))
	//fmt.Println(comm_non_recursive("3211876"))

	fmt.Println(comma_deal_sign("-14012323"))

	fmt.Println(string_anagrams("1243","2314"))

	//fmt.Println(intsToString([]int{1,2,3,44}))
}
package main
import (
	"fmt"
	"strings"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main () {
	//fmt.Println(comma("12345"))
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(s, b, s2)

	s3 := strings.Fields("today \n is a very love day")
	fmt.Println(s3[3]) // very

	fmt.Println(strings.Count("sokme one look it.", "ok")) // 2

	fmt.Println(strings.Index("so what are you?","a")) //5
}
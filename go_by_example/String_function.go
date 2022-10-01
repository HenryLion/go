package main

import "fmt"
import s "strings"

var p = fmt.Println // 这样也行!!

func main() {
	p("Contains: ", s.Contains("test", "es"))
}

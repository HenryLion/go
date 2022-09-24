package main

import (
	"fmt"
	"unicode/utf8"
)

// strings: a read-only slice of bytes, encoded as utf-8
// rune: like other language's character

func main() {
	const s = "สวัสดี"
	fmt.Println(len(s)) // 18

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U start at %d\n", runeValue, idx)
	}
}

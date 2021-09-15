package main
import (
	"fmt"

)

// TO string 4 important packages:
// bytes, strings, strconv, unicode

func main () {
	// normal string with "", process escapes
	str := "abc什么情况\fdef\f好的哦\x6a\176\n"
	fmt.Println(str)

	// normal string with ``, not process escapes
	str = `abc什么情况\fdef\f好的哦\x6a\176\n`
	fmt.Println(str)

	const GoUsage = `Go is a tool for managing Go source code.
	Usage:
		  go command [option]
		  ...`
	fmt.Println(GoUsage)


	// utf-8
	fmt.Println("世界")
	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c") // "世" 的utf-8编码是e4 b8 96 三个字节 即二进制是这么存的
	fmt.Println("\u4e16\u754c") // 汉子 "世" 序号是 4e16 即 19990 如果要使用序号打印的话需要加\u
	fmt.Println("\U00004e16\U0000754c")

	var zhongwen = "世界"
	byteZhong := []byte(zhongwen)
	fmt.Printf("% x\n",byteZhong) // e4 b8 96 e7 95 8c

}
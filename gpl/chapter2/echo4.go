package main

import (
	"flag"
	"fmt"
	"strings"
)

//生成一个变量nVal 表示一个默认的 -n 命令行选项值 true
var nVal = flag.Bool("n", false, "omit trailing newline")

// 生成一个变量sepVal 表示一个默认的 -s 命令行选项值 "#"
var sepVal = flag.String("s", "#", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sepVal))
	if !*nVal {
		fmt.Println("\n")
	}

	p := new(int)
	fmt.Println(p)
	fmt.Println(*p)
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	/*
		lineCounts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			lineCounts[input.Text()]++
		}
		for line, n := range lineCounts {
			fmt.Printf("%s - %d\n", line, n)
		}
	*/

	var b bool
	var u uint32 = 54
	var i int32 = -56
	var s string = "中国人"
	fmt.Printf("%v\t%v\t%v\t%v\n", b, u, i, s)

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			old_len := len(counts)
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			new_len := len(counts)
			if old_len != new_len {
				fmt.Printf("file %s has multi line.\n", arg)
			}
			f.Close()
		}
	}
	for line, n := range counts {
		fmt.Printf("%d:\t%s\n", n, line)
	}

	count2 := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFile %v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			count2[line]++
		}
	}
	for line, n := range count2 {
		fmt.Printf("%d:\t%s\n", n, line)
	}
}

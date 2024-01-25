package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputFile := "/Users/maegeek/Desktop/log1.txt" // 源文件名
	outputFile := "output.txt"                     // 输出文件名

	// 读取文件内容
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 按行分割，统计数字出现次数
	lines := strings.Split(string(data), "\n")
	counts := make(map[int]int)
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err == nil {
			counts[num]++
		}
	}

	// 将数字和出现次数存入切片，并按出现次数排序
	type numCount struct {
		Number int
		Count  int
	}
	numCounts := make([]numCount, 0, len(counts))
	for num, count := range counts {
		numCounts = append(numCounts, numCount{num, count})
	}
	sort.Slice(numCounts, func(i, j int) bool {
		return numCounts[i].Count > numCounts[j].Count
	})

	// 将排序后的结果输出到新文件
	f, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	for _, nc := range numCounts {
		line := fmt.Sprintf("%d: %d\n", nc.Number, nc.Count)
		writer.WriteString(line)
	}
	writer.Flush()
}

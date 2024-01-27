package main

import (
	"log"
)

func main() {
	mp := make(map[int]string)
	if mp == nil {
		log.Println("map is nil") // 使用make创建的map不为nil
	}

	var nil_map map[int]string
	if nil_map == nil {
		log.Println("map is nil here")
	}

	var nil_slice []int
	if nil_slice == nil {
		log.Println("slice is nil")
	}
}

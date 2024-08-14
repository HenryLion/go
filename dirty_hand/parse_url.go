package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func main() {
	// 定义要解析的字符串
	str := "?version=1&src_type=app&isNeedReply=true&SwitchExtMsg=1&isExt=false&QidianPubID=0&QidianKfUin=3009074258&origExt=2881289492&ToExt=2881289491&reptype=4&scene=5"
	// 解析url
	u, err := url.Parse(str)
	if err != nil {
		panic(err)
	}

	// 获取query参数
	queryParams := u.Query()

	// 解析各个参数
	version := queryParams.Get("version")
	if version == "" {
		version = "1.0"
	}
	srcType := queryParams.Get("src_type")
	isNeedReply := queryParams.Get("isNeedReply")
	switchExtMsg := queryParams.Get("SwitchExtMsg")
	isExt := queryParams.Get("isExt")
	qidianPubID := queryParams.Get("QidianPubID")
	qidianKfUin := queryParams.Get("QidianKfUin")
	origExt := queryParams.Get("origExt")
	toExt := queryParams.Get("ToExt")
	reptype := queryParams.Get("reptype")
	scene := queryParams.Get("scene")
	abc := queryParams.Get("abc")
	if abc == "" {
		fmt.Println("abc is empty")
	}

	abcNum, err := strconv.ParseUint(abc, 10, 64)
	if err != nil {
		fmt.Println("abc is not a number: err: %+v", err)
	} else {
		fmt.Printf("abcNum: %d\n", abcNum)
	}

	// 输出解析结果
	fmt.Printf("version: %s\n", version)
	fmt.Printf("src_type: %s\n", srcType)
	fmt.Printf("isNeedReply: %s\n", isNeedReply)
	fmt.Printf("SwitchExtMsg: %s\n", switchExtMsg)
	fmt.Printf("isExt: %s\n", isExt)
	fmt.Printf("QidianPubID: %s\n", qidianPubID)
	fmt.Printf("QidianKfUin: %s\n", qidianKfUin)
	fmt.Printf("origExt: %s\n", origExt)
	fmt.Printf("ToExt: %s\n", toExt)
	fmt.Printf("reptype: %s\n", reptype)
	fmt.Printf("scene: %s\n", scene)

	fmt.Printf("ToExt value 1: %s\n", toExt)

}

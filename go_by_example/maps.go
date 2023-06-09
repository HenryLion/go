package main

import "fmt"

type UserInfo struct {
	Name string
}

func ReturnMap() (sessionInfo map[string]*UserInfo) {
	resultSessionInfo := make(map[string]*UserInfo)
	resultSessionInfo["one"] = &UserInfo{"default1"}
	resultSessionInfo["two"] = &UserInfo{"default2"}
	return resultSessionInfo
}

func ModifyMapInfo(sessionInfo map[string]*UserInfo) {
	if info, ok := sessionInfo["one"]; ok {
		info.Name = "lili"
	}
	if info, ok := sessionInfo["two"]; ok {
		info.Name = "meili"
	}
}

func PrintMap(sessionInfo map[string]*UserInfo) {
	for _, v := range sessionInfo {
		fmt.Println(*v)
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	sessioonInfos := ReturnMap()
	PrintMap(sessioonInfos)

	ModifyMapInfo(sessioonInfos)
	PrintMap(sessioonInfos)

	//mapkey := make([]string, 0)
	//map1 := make(map[string]string)
	//map1["hello"] = "echo hello"
	//mapkey = append(mapkey, "hello")
	//map1["world"] = "echo world"
	//mapkey = append(mapkey, "world")
	//map1["go"] = "echo go"
	//mapkey = append(mapkey, "go")
	//map1["is"] = "echo is"
	//mapkey = append(mapkey, "is")
	//map1["cool"] = "echo cool"
	//mapkey = append(mapkey, "cool")
	//
	//for _, v := range mapkey {
	//	fmt.Println(v, "->", map1[v])
	//}

	seqNum := makeRange(0, 134)
	loopCount := 30 // 每次遍历个数
	loopRound := 0
	if len(seqNum)%loopCount == 0 {
		loopRound = len(seqNum) / loopCount
	} else {
		loopRound = len(seqNum)/loopCount + 1
	}

	fmt.Printf("total: %d, loopCount: %d, loopNum: %d\n", len(seqNum), loopCount, loopRound)
	for i := 0; i < loopRound; i++ {
		fmt.Println("time: ", i+1)
		batchNum := make([]int, 0)
		if (i+1)*loopCount < len(seqNum) {
			batchNum = seqNum[i*loopCount : (i+1)*loopCount]
		} else {
			batchNum = seqNum[i*loopCount:]
		}
		fmt.Println(batchNum)
	}

}

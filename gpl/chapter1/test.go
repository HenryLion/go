package main

import (
	"fmt"
	"strings"
)
import "strconv"

func buildRoleFilterStr(roleids []uint32) (roleFilterstr string) {
	fmt.Printf("buildRoleFilterStr roleids (%v)\n", roleids)
	if len(roleids) == 0 {
		return
	}
	var strId string
	for i, v := range roleids {
		strId = strconv.Itoa(int(v))
		if i != len(roleids)-1 {
			//roleFilterstr += " knowledgeRoleList eq \"" + strId + "\" or "
			roleFilterstr += `knowledgeRoleList eq "` + strId + `" or `
		} else {
			//roleFilterstr += " knowledgeRoleList eq \"" + strId + "\""
			roleFilterstr += `knowledgeRoleList eq "` + strId + `"`
		}
	}
	return
}

func main() {
	ids := []uint32{12345, 654567, 23456134}
	fmt.Printf("xiediansha (%s)\n", buildRoleFilterStr(ids))

	strnum := "15678"
	num, _ := strconv.Atoi(strnum)
	fmt.Println(num)

	newStrNum := fmt.Sprintf("%03d", num+1)
	fmt.Println(newStrNum)

	str1 := "RW-0007"
	if strings.Contains(str1, "-") {
		pos := strings.LastIndex(str1, "-")
		purenum := str1[pos+1:]
		fmt.Println(purenum)
	}

	fmt.Println(len("1234erdf"))
	fmt.Println(len("中国人"))
}

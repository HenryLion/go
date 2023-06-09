package main

import (
	"encoding/json"
	"fmt"
)

func changeStrValue(str *string) {
	*str = *str + "hello"
}

func changeStrValue2(str string) (strRsp string) {
	str = str + "hello"
	return str
}

func main() {
	str := "world"
	changeStrValue(&str)
	fmt.Println(str)

	str2 := changeStrValue2(str)
	fmt.Println(str2)

	str3 := "shenme"
	str3 = changeStrValue2(str3)
	fmt.Println(str3)

	s := []int{}
	fmt.Println(len(s))

	req := make(map[string]interface{})
	//pack from & size
	req["from"] = 0
	req["size"] = 100

	//pack source
	source := make(map[string]interface{})
	var neededFields []string
	neededFields = append(neededFields, "FCUin")
	source["include"] = neededFields
	req["_source"] = source

	//pack sort
	var sort []interface{}
	sortField := make(map[string]interface{})
	sortCond := make(map[string]interface{})
	sortCond["order"] = "asc"
	sortField["FCUin"] = sortCond
	sort = append(sort, sortField)
	req["sort"] = sort

	//pack query
	boolCond := make(map[string]interface{})
	filter := make(map[string][]interface{})

	filterField1 := make(map[string]interface{})
	termCond1 := make(map[string]interface{})
	termCond1["CorpUin"] = 123
	filterField1["term"] = termCond1

	filterField2 := make(map[string]interface{})
	termCond2 := make(map[string]interface{})
	termCond2["FStatus"] = 0
	filterField2["term"] = termCond2
	filter["filter"] = append(filter["filter"], filterField1)
	filter["filter"] = append(filter["filter"], filterField2)
	boolCond["bool"] = filter
	req["query"] = boolCond

	reqStr, _ := json.Marshal(req)
	fmt.Printf("%s\n", string(reqStr))
}

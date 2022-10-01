package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(bolB)         // []byte
	fmt.Println(string(bolB)) // byte可以string化

	intB, _ := json.Marshal(1)
	fmt.Println(intB)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(fltB)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(slcB)
	fmt.Println(string(slcB))

	// 把map弄成json格式：
	// 直接把Marshal的结果[]byte转成string即可
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(mapB)
	fmt.Println(string(mapB))

	str2, _ := json.Marshal("Сайн уу")
	fmt.Println(str2)
	fmt.Println(string(str2))

	// 把struct弄成jaon字符串
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(res1D)
	fmt.Println(res1B)
	fmt.Println(string(res1B))

	res2D := response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(res2B)
	fmt.Println(string(res2B))

	var dat map[string]any

	// 从jsoon字符串获取结构体信息
	// 先把字符串[]byte
	// 使用Unmarshal将[]弄成map
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	fmt.Println(byt)
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Unmarshal可以把[]byte解析成map或者一个struct
	str := `{"page":1, "fruits":["apple","banana","pear"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}

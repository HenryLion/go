package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(tt)
	return nil
}

// input 不符合go默认的time.Time格式，所以需要自定义一个类型，然后实现UnmarshalJSON方法
var input = `{"create_at": "Thu May 31 00:00:01 +0000 2012"}`

// input2 符合go默认的time.Time格式，所以不需要自定义类型，直接使用time.Time即可
var input2 = `{"create_at": "1999-03-05T17:04:05Z"}`

func main() {
	//var vau map[string]interface{}
	var vau map[string]Timestamp
	if err := json.Unmarshal([]byte(input), &vau); err != nil {
		panic(err)
	}
	fmt.Println(vau)
	for k, v := range vau {
		fmt.Println(k, reflect.TypeOf(v))
	}

	realTime := time.Time(vau["create_at"])
	fmt.Println(realTime.Date())
	fmt.Println(realTime.Day())

	var val map[string]time.Time
	if err := json.Unmarshal([]byte(input2), &val); err != nil {
		panic(err)
	}
	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
}

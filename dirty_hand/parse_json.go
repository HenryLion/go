package main

import (
	"encoding/json"
	"fmt"
)

// 处理动态JSON数据
func ProcessDynamicJSON(jsonStr string) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	processValue("root", data)
}

func processValue(key string, value interface{}) {
	switch v := value.(type) {
	case map[string]interface{}:
		fmt.Printf("%s: Object with %d fields\n", key, len(v))
		for k, val := range v {
			processValue(k, val)
		}
	case []interface{}:
		fmt.Printf("%s: Array with %d elements\n", key, len(v))
		for i, val := range v {
			processValue(fmt.Sprintf("%s[%d]", key, i), val)
		}
	case string:
		fmt.Printf("%s: String = %q\n", key, v)
	case float64:
		fmt.Printf("%s: Number = %.2f\n", key, v)
	case bool:
		fmt.Printf("%s: Boolean = %t\n", key, v)
	case nil:
		fmt.Printf("%s: Null\n", key)
	default:
		fmt.Printf("%s: Unknown type %T\n", key, v)
	}
}

func main() {
	jsonData := `{
        "name": "张三",
        "age": 30,
        "active": true,
        "scores": [85, 92, 78],
        "address": {
            "city": "北京",
            "zipcode": "100000"
        },
        "spouse": null
    }`

	ProcessDynamicJSON(jsonData)
}

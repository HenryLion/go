package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// 从 Redis 读取的字符串
	//redisData := "{\"json\":\"\",\"kfuin\":0,\"vId\":\"305139177084059\",\"roleValue\":0,\"roleData\":0,\"subNavType\":0,\"subNavID\":0,\"subNavTips\":\"\",\"clickid\":\"s9mgbw.krpyow.m6yt65og\",\"wpaType\":0,\"wpaId\":\"1201\",\"isCustomRoute\":0,\"receptionSwitch\":0,\"cLocation\":null,\"webVisitStats\":null,\"robotNavMenuSwitch\":0,\"robotNavMenuID\":0,\"customParams\":{},\"translateSwitch\":0,\"translateLang\":\"\",\"reception_type\":0,\"reception_id\":\"\",\"custom\":null,\"routeId\":0,\"type\":1,\"nav\":{\"menu_name\":\"\xe8\xae\xa2\xe5\x8d\x95\xe5\x92\xa8\xe8\xaf\xa2\"},\"ip\":\"111.122.47.136\"}"
	redisData := ""
	// 解析 JSON 字符串
	var data map[string]interface{}
	err := json.Unmarshal([]byte(redisData), &data)
	if err != nil {
		fmt.Println("JSON 解析失败:", err)
		return
	}
	// 从data中获取wpaId
	wpaId, ok := data["kfuin"]
	if !ok {
		fmt.Println("vId 不存在")
		return
	} else {
		fmt.Printf("vId:%v,type:%v", wpaId, reflect.TypeOf(wpaId))
	}
}

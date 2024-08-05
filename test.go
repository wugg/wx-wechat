package main

import (
	"encoding/json"
	"fmt"
	"wechat/offiaccount"
)

func main() {
	var Jsine *offiaccount.JsSign
	Jsine, _ = offiaccount.GetJsSign("http://www.baidu.com", "", "", "")
	jsineBytes, err := json.Marshal(Jsine)
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	// 输出 JSON 数据
	fmt.Println(string(jsineBytes))
}

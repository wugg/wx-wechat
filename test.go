package main

import (
	"fmt"
	"wechat/offiaccount"
)

func main() {
	var Jsine *offiaccount.JsSign
	Jsine, _ = offiaccount.GetJsSign("http://www.baidu.com", "", "", "")
	fmt.Println(Jsine)
}

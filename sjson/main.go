package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {
	var appid string
	value := `{"status":200,"appid":"123","content":{"type":"normal","speaker":"xiaoi","msg":"asdf 中文"}}`
	json, err := simplejson.NewJson([]byte(value))
	if err != nil {
		return
	}

	if reply, ok := json.Get("appid").String(); ok != nil {
		appid = "default appid"
	} else {
		appid = reply
	}

	fmt.Println(appid)
}

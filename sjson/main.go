package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

type dd struct {
	Type    string `json:"type"`
	Speaker string `json:"speaker"`
	Msg     string `json:"msg"`
}

func main() {
	value := `{"status":200,"appid":"123","content":[{"type":"normal","speaker":"xiaoi","msg":"asdf 中文"}]}`
	sjson, err := simplejson.NewJson([]byte(value))
	if err != nil {
		return
	}

	tmpArr := []dd{}
	if reply, ok := sjson.Get("content").Array(); ok != nil {
		//appid = "default appid"
	} else {
		//d := &dd{}
		b, _ := json.Marshal(reply)
		json.Unmarshal(b, &tmpArr)
		//fmt.Printf("%#v", tmpArr)

		for _, v := range tmpArr {
			fmt.Println(v.Speaker)
		}
	}

}

package routes

import (
	"fmt"
	"testing"
)

func TestWxauthRequestBody_GetMinappAtuh(t *testing.T) {
	body := WxauthRequestBody{
		appid:  "a",
		code:   "c",
		secret: "d",
	}
	r, e := body.GetMinappAtuh()
	fmt.Println(r.MinOpenid, r.Id, e)
}

func BenchmarkWxauthRequestBody_GetMinappAtuh(b *testing.B) {
	body := WxauthRequestBody{
		appid:  "a",
		code:   "c",
		secret: "d",
	}
	r, e := body.GetMinappAtuh()
	fmt.Println(r.MinOpenid, r.Id, e)
}

package utils

import (
	"fmt"
	"testing"
)

func TestGetIntranetIp(t *testing.T) {
	s := GetIntranetIp()
	fmt.Println(s)
}

func TestParseDns(t *testing.T) {
	s := ParseDns("redishost")
	fmt.Println(s)
}

func TestSha1s(t *testing.T) {
	a := `
{"nickName":"香港电影半边天","gender":1,"language":"zh_CN","city":"朝阳","province":"北京","country":"中国","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTICc2r9u5nkrbuElRsrPEZP0ia8Dap0mx85ibtfU8Hic8Ciam8TLCUADPExn7ahAMPqe6Drzjn6AFr2vA/132"}
`
	b := "2o6CQOkpLB+mjYopmdEGyQ=="
	s := Sha1s(fmt.Sprintf("%s%s", a, b))
	fmt.Println("c5610950c484478d0985ffc3ce009aed524a0376")
	fmt.Println(s)
}

func TestTokenValid(t *testing.T) {
	token := "1.eyJ1IjoiZjkzYTczNzNlYjZkYjM0MDQwODU1ZTU4YzdjMWY2ZWIiLCJlIjoxNTQ0NTA4NDkwLCJzayI6IkFsdkd2VEFGN2RwL3o0Wkg4Qzk3c3c9PSJ9.W-gverjIwGf0CcfWUcP61Q"
	secret_key := "92a57168e985cbbc2c5713671c4cc147dd42a020"
	userid, nickname, headimgurl, err := TokenValid(token, secret_key)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(userid, nickname, headimgurl)
}

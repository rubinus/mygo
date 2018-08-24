package main

import (
	"fmt"
	"mygo/queue/q"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.IsEmpty()
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	q.Pop()
	q.IsEmpty()

	str := []byte(`{"status": 503,"msg": "身份已过期,请重新登录."}`)
	str0 := "{\x22status\x22:503,\x22msg\x22:\x22\xE8\xBA\xAB\xE4\xBB\xBD\xE5\xB7\xB2\xE8\xBF\x87\xE6\x9C\x9F,\xE8\xAF\xB7\xE9\x87\x8D\xE6\x96\xB0\xE7\x99\xBB\xE5\xBD\x95.\x22}"

	str1 := "{\x22status\x22:200,\x22data\x22:{\x22contentInfo\x22:{\x22unionid\x22:\x22oCLOAuGNsl5lzmgsvojT-6bUmEZQ\x22,\x22headimgurl\x22:\x22\x22,\x22country\x22:\x22\x22,\x22province\x22:\x22\x22,\x22city\x22:\x22\x22,\x22language\x22:\x22\x22,\x22sex\x22:1,\x22nickname\x22:\x22\x22,\x22openid\x22:\x22\x22},\x22tvmid\x22:\x22wxh59f0bc138a057c595d74f84f\x22,\x22mobile_number\x22:\x2215249251644\x22,\x22withdraw\x22:0,\x22coin\x22:500,\x22cash\x22:58,\x22auth_flag\x22:1,\x22daily_giveout\x22:\x22\x22,\x22invitation_code_old\x22:\x22\x22,\x22invitation_code\x22:\x22A8488822\x22,\x22inviter_tvmid_flag\x22:1,\x22isWelfareNewUser\x22:true,\x22unionid_flag\x22:1,\x22unionid\x22:\x22oCLOAuGNsl5lzmgsvojT-6bUmEZQ\x22,\x22tvmappid\x22:\x22\x22,\x22ttopenid_flag\x22:1,\x22ttopenid\x22:\x22orEt2t7lF-cPqlewMRumidyZvwQA\x22,\x22wechat\x22:\x22oTA1CwBPhkKwdsPUrwyQieywC8Uw\x22,\x22username\x22:\x22\x22,\x22invitation_codeTime\x22:1508949012360,\x22inviter_time\x22:1508949012924,\x22inviter_tvmid\x22:\x22wxh5888994ad82dd4236d327c3c\x22,\x22notes\x22:{\x22maintype\x22:1,\x22subtype\x22:21,\x22source\x22:1,\x22money\x22:1,\x22note\x22:\x22\xE5\xA5\xBD\xE5\x8F\x8B\xE7\x9F\xB3\xE5\xA4\xB4\xE6\x91\x87\xE5\x88\xB0\xE6\x8F\x90\xE7\x8E\xB0\xE9\x94\xA6\xE5\x9B\x8A\xE8\x8E\xB7\xE5\xBE\x97\xE5\xA5\x96\xE5\x8A\xB1\x22,\x22time\x22:1513692439285},\x22isRecommend\x22:2,\x22latitude\x22:34.315224,\x22longitude\x22:108.96518,\x22mtqsign\x22:\x22fe567a7a045ea045609eeeffd91db9e6\x22,\x22sigExpire\x22:1532678863,\x22ttdsbappid\x22:\x22wxd06496bae6bb4a78\x22,\x22yaoSig\x22:\x227f58c190ce9ee38d2bb58b1b3f830816\x22,\x22wxTokenSig\x22:\x22d378cf1259ae2e4ca1785eb1c81253d2\x22,\x22ttdsbwx_token\x22:\x2246497107fa23\x22,\x22userinfo_sign\x22:\x22a51fd8c09fb98ea60670d72015155d91\x22,\x22del_flag\x22:false},\x22token\x22:\x22cdf84a451336bd852ff67983e2f948f7\x22}"
	fmt.Println(str1)
	fmt.Println(len(str0))
	fmt.Println(len(string(str)))
}

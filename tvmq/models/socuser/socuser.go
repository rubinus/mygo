package socuser

import "github.com/json-iterator/go"

type SoconnUser struct {
	UserId string `json:"userId" redis:"userId"`
	Host   string `json:"host" redis:"host"`
	Cid    string `json:"cid" redis:"cid"`
}

type Userinfo struct {
	MinAppid   string `json:"minappid" redis:"minappid"`
	Nickname   string `json:"nickname" redis:"nickname"`
	Headimgurl string `json:"headimgurl" redis:"headimgurl"`
	Id         string `json:"id" redis:"id"`
	City       string `json:"city" redis:"city"`
	Province   string `json:"province" redis:"province"`
	Country    string `json:"country" redis:"country"`
	Token      string `json:"token" redis:"token"`
}

type UserLastSendInfo struct {
	Userid   string `redis:"userid"`
	LastTime int64  `redis:"lasttime"`
}

func (u *Userinfo) StructToMap() map[string]interface{} {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	var m map[string]interface{}
	b, _ := jsonIterator.Marshal(u)
	jsonIterator.Unmarshal(b, &m)
	return m
}

func (u *SoconnUser) StructToMap() map[string]interface{} {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	var m map[string]interface{}
	b, _ := jsonIterator.Marshal(u)
	jsonIterator.Unmarshal(b, &m)
	return m
}

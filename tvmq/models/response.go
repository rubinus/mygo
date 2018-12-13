package models

import (
	"github.com/json-iterator/go"
)

type ResErrBody struct {
	Status int    `json:"status"`
	Key    string `json:"key"`
	Msg    string `json:"msg"`
}

type ResFormanBody struct {
	Status  int    `json:"status"`
	Content string `json:"content"`
	From    string `json:"from"`
}
type ResForman struct {
	Type    int    `json:"type"`
	Speaker string `json:"speaker"`
	Msg     string `json:"msg"`
}

type ResAiBody struct {
	Status  int        `json:"status"`
	Content ResComment `json:"content"`
	From    string     `json:"from"`
}

type ResChatBody struct {
	Status  int        `json:"status"`
	Content ResComment `json:"content"`
	From    string     `json:"from"`
}
type ResComment struct {
	Userid     string `json:"userid"`
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
	City       string `json:"city,omitempty"`
	Province   string `json:"province,omitempty"`
	Country    string `json:"country,omitempty"`
	Message    string `json:"message,omitempty"`
}

type ResGiftBody struct {
	Status  int     `json:"status"`
	Content ResGift `json:"content"`
	From    string  `json:"from"`
}
type ResGift struct {
	Userid     string `json:"userid"`
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
	Giftid     string `json:"giftid"`
	Giftname   string `json:"giftname"`
	Icon       string `json:"icon"`
	Pictures   string `json:"pictures"`
	Count      int    `json:"count"`
	Points     int    `json:"points"`
}

type ResAuthBody struct {
	Status  int      `json:"status"`
	Content AuthBody `json:"content"`
	From    string   `json:"from"`
}
type AuthBody struct {
	Result string `json:"result"`
}

type ResOnlineBody struct {
	Status  int        `json:"status"`
	Content OnlineBody `json:"content"`
	From    string     `json:"from"`
}
type OnlineBody struct {
	Count int64 `json:"count"`
}

func NewResErrBody(status int, msg, key string) string {
	strErr := ResErrBody{
		Status: status,
		Msg:    msg,
		Key:    key,
	}
	reply, _ := Marshal(strErr)
	return reply
}

func Marshal(res interface{}) (string, error) {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	s, err := jsonIterator.Marshal(res)
	return string(s), err
}

func (r *ResFormanBody) read() (interface{}, error) {
	return r, nil
}

func (r *ResAuthBody) read() (interface{}, error) {
	return r, nil
}

func (r *ResChatBody) read() (interface{}, error) {
	return r, nil
}

func (r *ResComment) read() (interface{}, error) {
	return r, nil
}

func (r *ResGift) read() (interface{}, error) {
	return r, nil
}

type Unmarshaler interface {
	read() (interface{}, error)
}

func Unmarshal(message string, unmarshaler Unmarshaler) (interface{}, error) {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	jsonIterator.Unmarshal([]byte(message), &unmarshaler)
	return unmarshaler, nil
}

func UnmarshalNew(message []byte, in interface{}) (interface{}, error) {
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	jsonIterator.Unmarshal(message, in)
	return in, nil
}

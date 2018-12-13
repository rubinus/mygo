package backend

import (
	"context"
	"fmt"

	"code.tvmining.com/tvplay/tvmq/models"
	"github.com/json-iterator/go"
)

type FilterBody struct {
	Items []Items
}

type FilterBodyReply struct {
	Kind  string `json:"kink"`
	Items []Items
}

type Items struct {
	Id     string  `json:"id"`
	UserId string  `json:"user_id"`
	Text   string  `json:"text"`
	Score  float32 `json:"score"`
}

func SendFilter(ctx context.Context, url string, body FilterBody) chan float32 {
	out := make(chan float32)
	go func() {
		defer close(out)

		s, err := Post(url, body)
		if err != nil {
			fmt.Println(err)
			out <- -1
			return
		}
		su := models.ErrorFetchPoint{}
		//time.Sleep(5 * time.Second)
		models.Unmarshal(string(s), &su)
		if su.Error.Code != "" {
			out <- -1
		} else {
			reply := FilterBodyReply{}
			jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
			jsonIterator.Unmarshal([]byte(s), &reply)
			out <- reply.Items[0].Score
		}
		//fmt.Println( string(s))
	}()
	return out
}

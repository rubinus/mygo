package backend

import (
	"context"
	"fmt"
	"strconv"

	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/gift"
)

func GetPoint(ctx context.Context, u string) chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		s, err := SendGet(u)
		if err != nil {
			fmt.Println(err)
		}
		su := models.ErrorFetchPoint{}
		//time.Sleep(5 * time.Second)
		models.Unmarshal(string(s), &su)
		if su.Error.Code != "" {
			out <- -1
		} else {
			reply := models.FetchPoint{}
			models.Unmarshal(string(s), &reply)
			out <- reply.Points
			//out <- 200
		}
		//fmt.Println(u, string(s))
	}()
	return out
}

func SendPoint(ctx context.Context, u string, postParams interface{}) chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		gs := postParams.(*gift.Gift)
		play := make(map[string]interface{})

		play["user_id"] = gs.Userid
		play["giftid"] = gs.Giftid
		play["giftname"] = gs.Giftname
		play["icon"] = gs.Icon
		play["pictures"] = gs.Pictures
		play["point"] = strconv.Itoa(gs.Points)
		play["count"] = strconv.Itoa(gs.Count)

		s, err := Post(u, play)
		if err != nil {
			out <- err.Error()
			return
		}

		//time.Sleep(5 * time.Second)
		//fmt.Println(fmt.Sprintf("%s", s))
		//b,_ := json.Marshal(play)
		//fmt.Println("POST Body \n",string(b))

		su := models.ErrorFetchPoint{}
		models.Unmarshal(string(s), &su)

		if su.Error.Code != "" {
			out <- su.Error.Message
		} else {
			out <- "OK"
		}

	}()
	return out
}

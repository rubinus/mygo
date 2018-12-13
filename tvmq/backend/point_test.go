package backend

import (
	"context"
	"fmt"
	"testing"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/models/gift"
	"github.com/bitly/go-simplejson"
)

func TestSendPoint(t *testing.T) {
	gs := gift.Gift{
		Userid:     "1000",
		Headimgurl: "",
		Nickname:   "test",
		Giftid:     "aaaaaaaaan",
		Giftname:   "",
		Icon:       "",
		Pictures:   "",
		Count:      1,
		Points:     1,
	}
	url := fmt.Sprintf("%s%s?userid=%s&token=%s", config.PointHost, config.PointPostPath, "1000", "1000")

	//url := fmt.Sprintf("%s%s", config.PointHost, config.PointPostPath)
	//url := "http://172.24.0.11:8081/rpcsend"
	fmt.Println(url)
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	out := SendPoint(ctx, url, &gs)
	select {
	case <-ctx.Done():
		fmt.Println("超时")
	case result := <-out:
		fmt.Println(result)

	}

	//n := time.Now().UnixNano()
	//fmt.Println(n / 1e6)
	//fmt.Println(time.Now().Add(60 * time.Second).UnixNano() / 1e6)

}

func TestGetPoint(t *testing.T) {
	var errStr string
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	url := fmt.Sprintf("%s%s%s&token=%s", config.PointHost, config.PointGetPath, "5ba9fc1efd865b000135c8c9", "a37450f85a703f9ef4c368d776b476d8")
	pointCh := GetPoint(ctx, url)
	select {
	case <-ctx.Done():
		errStr = "超时GetPoint"
		fmt.Println(errStr)
	case r := <-pointCh:
		//if r != -1 && r <= count*points {
		//	errStr = "积分不够"
		//}
		fmt.Println(r)
	}
	if errStr != "" {
		fmt.Println(errStr)
	}

	fmt.Println(time.Now().UnixNano() / 1e9)

	r := `{"status":200,"content":{"hd_id":"5ba9af9f56bcab68","title":"测试开始时间","video":"上东方闪电","countdown":"80","icon":"的释放都是","pictures":"随碟附送","type":"1","start_time":1537847220,"question":[{"hd_id":"5ba9af9f56bcab68","title":"以下哪一例餐品为潮骨派主打招牌？","answer":[{"id":"5ba9afb448173b13","title":"酱香大骨饭","state":1},{"id":"5ba9afb4c4b78553","title":"辣骨饭"},{"id":"5ba9afb46f9afcae","title":"麻椒大骨"},{"id":"5ba9afb42b3168a0","title":"酸辣排骨"}]}]}}`

	json, _ := simplejson.NewJson([]byte(r))
	countdown := json.Get("content").Get("countdown")
	fmt.Println("countdown=", countdown)

	content := json.Get("content")
	content.Set("countdown", "1000")
	reply, _ := json.MarshalJSON()
	fmt.Println(string(reply))

}

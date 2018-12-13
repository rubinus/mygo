package comm

import (
	"context"
	"fmt"

	"code.tvmining.com/tvplay/tvmq/models/comment"
	"code.tvmining.com/tvplay/tvmq/models/gift"
	"code.tvmining.com/tvplay/tvmq/models/traceinfodb"
	"code.tvmining.com/tvplay/tvmq/utils"
	"gopkg.in/mgo.v2/bson"
)

func SaveMessageToMongodb(ctx context.Context, c interface{}) chan string {
	//把信息存到mongodb中
	out := make(chan string)
	go func() {
		defer close(out)
		switch c.(type) {
		case *comment.Comment:
			//fmt.Println("当前发的是弹幕消息")
			dc := c.(*comment.Comment)
			_, err := dc.SaveComment() //发到消息队列
			if err != nil {
				fmt.Println("save comment is failed", err)
				out <- err.Error()
				return
			}
			out <- "ok"

		case *gift.Gift:
			//fmt.Println("当前发的是礼物")
			gc := c.(*gift.Gift)
			_, err := gc.SaveGift() //发到消息队列

			if err != nil {
				fmt.Println("save nsgift is failed", err)
				out <- err.Error()
				return
			}

			out <- "ok"

		}

		//fmt.Printf("%v %s", c, "=====save mongo after======")

	}()
	return out
}

func SaveTraceinfoToMongodb(ctx context.Context, reqti *traceinfodb.Traceinfo) chan *traceinfodb.Traceinfo {
	//把信息存到mongodb中
	out := make(chan *traceinfodb.Traceinfo)
	go func() {
		defer close(out)

		ti, err := reqti.GetTraceinfoByTraceId()
		if err != nil || ti.Id == "" {
			//save
			id, err := reqti.SaveTraceinfo()
			if err != nil {
				out <- nil
				return
			} else {
				ti.Id = bson.ObjectId(id)
			}
		} else {
			//update
			id := ti.Id.Hex()
			ti.UpdateTime = utils.GetCurrentTime(13)
			si := ti.StructToMap()
			m := reqti.StructToMap()
			for k, v := range m {
				//fmt.Printf("k: %s, v: %s\n",k,v)
				if k != "id" && k != "traceId" && v != "" {
					si[k] = v
				}
			}
			delete(si, "id")
			err := ti.UpdateById(id, si)
			if err != nil {
				out <- nil
				return
			}
		}
		out <- ti
		//fmt.Printf("%v %s", c, "=====save mongo after======")

	}()
	return out
}

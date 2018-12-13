package lottery

import (
	"context"
	"fmt"
	"os"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/kafka/kafkaCons"
	"github.com/bitly/go-simplejson"
	"github.com/bsm/sarama-cluster"
)

type Lottery struct {
	Brokers []string
	Topic   []string
	GroupId string
}

func (lottery Lottery) Consumer() {
	go kafkaCons.Consumer(lottery.Brokers, lottery.Topic, lottery.GroupId, func(consumer *cluster.Consumer, signals chan os.Signal) {
		for {
			select {
			case part, ok := <-consumer.Partitions():
				if !ok {
					return
				}

				// start a separate goroutine to consume messages
				go func(pc cluster.PartitionConsumer) {
					for msg := range pc.Messages() {
						//fmt.Printf("==lottery===%d %s\n", msg.Offset, msg.Value)
						//					message := `{"status":200,"content":{"level":"0","title":"iphone xs Max","ico":
						//"https://qa.tvplay.tvm.cn/pub/tvmini/malllist/003.png","msg":"","type":"","count":"1","unit":"个","
						//nickname":"Dandelion","headimgurl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKUdkK5dVTue0QKMJLt9A
						//vnHR4h5jotic06sgHBvPMTpU0QsMe3KYfro5qTfWVFs5oeyS66MjUzrMQ/132"}}`

						var errStr string
						var appid string
						var tenantId string
						ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
						defer cancel()
						value := fmt.Sprintf("%s", msg.Value)

						json, err := simplejson.NewJson([]byte(msg.Value))
						if err != nil {
							return
						}
						//提取appid
						if reply, ok := json.Get("appid").String(); ok != nil {
							appid = config.Minappid
						} else {
							appid = reply
						}
						//提取tenantId
						if reply, ok := json.Get("tenantId").String(); ok != nil {
							tenantId = config.DefalutTenantId
						} else {
							tenantId = reply
						}

						out := backend.BroadcastByRedis(ctx, "", config.SocEventLottery, value, appid, "", tenantId)
						select {
						case <-ctx.Done():
							errStr = "超时BroadcastByRedis"
						case ok := <-out:
							if !ok {
								errStr = "广播失败"
							}
						}
						if errStr != "" {
							fmt.Println(errStr)
						} else {
							fmt.Println("broadcast successful")
						}
					}
				}(part)
			case <-signals:
				fmt.Println("lottery no signals ...")

				return
				//case msg, ok := <-consumer.Messages():
				//	if ok {
				//		fmt.Printf("==lottery===%d %s\n", msg.Offset, msg.Value)
				//		//					message := `{"status":200,"content":{"level":"0","title":"iphone xs Max","ico":
				//		//"https://qa.tvplay.tvm.cn/pub/tvmini/malllist/003.png","msg":"","type":"","count":"1","unit":"个","
				//		//nickname":"Dandelion","headimgurl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTKUdkK5dVTue0QKMJLt9A
				//		//vnHR4h5jotic06sgHBvPMTpU0QsMe3KYfro5qTfWVFs5oeyS66MjUzrMQ/132"}}`
				//
				//		var errStr string
				//
				//		ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
				//		defer cancel()
				//		value := fmt.Sprintf("%s", msg.Value)
				//
				//		out := backend.BroadcastByRedis(ctx, "", config.SocEventLottery, value)
				//		select {
				//		case <-ctx.Done():
				//			errStr = "超时BroadcastByRedis"
				//		case ok := <-out:
				//			if !ok {
				//				errStr = "广播失败"
				//			}
				//		}
				//		if errStr != "" {
				//			fmt.Println(errStr)
				//		} else {
				//			fmt.Println("broadcast successful")
				//		}
				//
				//	}
			}
		}
	})
}

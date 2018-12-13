package epg

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

type EPG struct {
	Brokers []string
	Topic   []string
	GroupId string
}

func (epg EPG) Consumer() {
	go kafkaCons.Consumer(epg.Brokers, epg.Topic, epg.GroupId, func(consumer *cluster.Consumer, signals chan os.Signal) {
		for {
			select {
			case part, ok := <-consumer.Partitions():
				if !ok {
					return
				}

				// start a separate goroutine to consume messages
				go func(pc cluster.PartitionConsumer) {
					for msg := range pc.Messages() {
						//fmt.Printf("===epg==%d %s\n", msg.Offset, msg.Value)
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

						out := backend.BroadcastByRedis(ctx, "", config.SocEventAi, value, appid, "", tenantId)

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
				fmt.Println("epg no signals ...")
				return
				//case msg, ok := <-consumer.Messages():
				//	if ok {
				//		fmt.Printf("===epg==%d %s\n", msg.Offset, msg.Value)
				//		ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
				//		defer cancel()
				//
				//		value := fmt.Sprintf("%s", msg.Value)
				//		out := backend.BroadcastByRedis(ctx,"", config.SocEventAi, value)
				//
				//		var errStr string
				//		select {
				//		case <-ctx.Done():
				//			errStr = "超时BroadcastByRedis"
				//		case ok := <-out:
				//			if !ok {
				//				errStr = "广播失败"
				//			}
				//		}
				//
				//		if errStr != "" {
				//			fmt.Println(errStr)
				//		}else{
				//			fmt.Println("broadcast successful")
				//		}
				//
				//	}
			}
		}
	})
}

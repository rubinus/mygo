package activity

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"code.tvmining.com/tvplay/tvmq/backend"
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/kafka/kafkaCons"
	"github.com/bitly/go-simplejson"
	"github.com/bsm/sarama-cluster"
)

type Activity struct {
	Brokers []string
	Topic   []string
	GroupId string
}

func (act Activity) Consumer() {
	go kafkaCons.Consumer(act.Brokers, act.Topic, act.GroupId, func(consumer *cluster.Consumer, signals chan os.Signal) {
		for {
			select {
			case part, ok := <-consumer.Partitions():

				if !ok {
					return
				}
				// start a separate goroutine to consume messages
				go func(pc cluster.PartitionConsumer) {
					for msg := range pc.Messages() {

						//fmt.Printf("==activity===%d %s\n", msg.Offset, msg.Value)

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

						countdown, err := json.Get("content").Get("countdown").String()
						if err != nil {
							return
						}
						ttl, err := strconv.Atoi(countdown)
						if err != nil {
							return
						}

						key := fmt.Sprintf("%s:%s:%s", config.CurrentActivity, appid, tenantId)
						ch := backend.SaveActvityToRedis(ctx, key, value, ttl)
						select {
						case <-ctx.Done():
							errStr = "超时SaveActvityToRedis"
						case ok := <-ch:
							if !ok {
								errStr = "保存互动失败"
							}
						}
						if errStr != "" {
							fmt.Println(errStr)
						} else {
							//fmt.Println("save activity successful")
						}

						out := backend.BroadcastByRedis(ctx, "", config.SocEventActivity, value, appid, "", tenantId)
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
							//fmt.Println("broadcast successful")
						}
					}
				}(part)
			case <-signals:
				fmt.Println("activity no signals ...")
				return

				//case msg, ok := <-consumer.Messages():
				//	if ok {
				//		fmt.Printf("==activity===%d %s\n", msg.Offset, msg.Value)
				//
				//		var errStr string
				//
				//		ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
				//		defer cancel()
				//
				//		value := fmt.Sprintf("%s", msg.Value)
				//
				//		key := fmt.Sprintf("%s:%s", config.CurrentActivity, config.Minappid)
				//		json, err := simplejson.NewJson([]byte(msg.Value))
				//		if err != nil {
				//			return
				//		}
				//
				//		//提取appid
				//		appid, err := json.Get("appid").String()
				//		if err != nil {
				//			return
				//		}
				//		if appid == "" {
				//			appid = config.Minappid
				//		}
				//
				//		countdown, err := json.Get("content").Get("countdown").String()
				//		if err != nil {
				//			return
				//		}
				//		ttl, err := strconv.Atoi(countdown)
				//		if err != nil {
				//			return
				//		}
				//		ch := backend.SaveActvityToRedis(ctx, key, value, ttl)
				//		select {
				//		case <-ctx.Done():
				//			errStr = "超时SaveActvityToRedis"
				//		case ok := <-ch:
				//			if !ok {
				//				errStr = "保存互动失败"
				//			}
				//		}
				//		if errStr != "" {
				//			fmt.Println(errStr)
				//		} else {
				//			//fmt.Println("save activity successful")
				//		}
				//
				//		out := backend.BroadcastByRedis(ctx, "", config.SocEventActivity, value, appid, "")
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
				//			//fmt.Println("broadcast successful")
				//		}
				//
				//	}
				//

			}
		}
	})
}

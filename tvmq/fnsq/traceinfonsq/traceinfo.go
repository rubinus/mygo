package traceinfonsq

import (
	"fmt"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsqService"
	"github.com/nsqio/go-nsq"

	"context"

	"code.tvmining.com/tvplay/tvmq/fnsq/comm"
	"code.tvmining.com/tvplay/tvmq/models"
	"code.tvmining.com/tvplay/tvmq/models/traceinfodb"
)

type Traceinfo struct {
	Topic   string
	Channel string
}

func (c Traceinfo) Consumer() {
	go nsqService.Consumer(c.Topic, c.Channel, config.NsqHostProt, 2, c.Deal)
}

//处理消息
func (c Traceinfo) Deal(msg *nsq.Message) error {
	message := msg.Body
	//fmt.Println("接收到NSQ", msg.NSQDAddress, "Traceinfo:", string(message))

	var errStr string

	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	traceBody := traceinfodb.Traceinfo{}
	models.UnmarshalNew(message, &traceBody)
	in := comm.SaveTraceinfoToMongodb(ctx, &traceBody)
	select {
	case <-ctx.Done():
		errStr = "超时SaveTraceinfoToMongodb"
	case traceBody := <-in:
		if traceBody == nil {
			errStr = "query failed"
		}
	}
	if errStr != "" {
		fmt.Println(errStr)
		return nil
	}
	//fmt.Println("traceBody...", traceBody)

	return nil
}

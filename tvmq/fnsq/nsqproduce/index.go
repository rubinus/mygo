package nsqproduce

import "code.tvmining.com/tvplay/tvmq/fnsq/nsqService"

type Nservice struct {
	Topic string
}

func (s Nservice) Producer(msg string) chan bool {
	return nsqService.Producer(s.Topic, msg)
}

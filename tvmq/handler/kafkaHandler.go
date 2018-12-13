package handler

import (
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/kafka/kafkaCons"
)

type KafkaAuto struct {
	Delay time.Duration
	Topic []string
}

func (hb KafkaAuto) Run() {
	go func() {
		for range time.Tick(hb.Delay) { //another way to get clock signal
			hb.Service()
			time.Sleep(hb.Delay)
		}
	}()
}

func (hb KafkaAuto) Service() {
	//asrcValue := `{"nickname":"涛涛GRT","headimgurl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIOYJOrUGCdDufvEA3te2YrwW0XVRtBpwa8Aic3Kcf6X9Sq6dukJFwxxNeKGo7gdsibhoEMU9DzvpYQ/132","userid":"5b973536a33c300001007749"}`
	asrcValue := `{"status":200,"content":{"type":"normal","speaker":"xiaoi","msg":"asdf 中文"}}`
	kafkaCons.AsyncProducer(hb.Topic[0], asrcValue)
}

func Start() {
	//e := KafkaAuto{
	//	Delay: 15 * time.Second,
	//	Topic: []string{config.TopicEpg},
	//}
	//e.Run()
	//
	//l := KafkaAuto{
	//	Delay: 30 * time.Second,
	//	Topic: []string{config.TopicLottery},
	//}
	//l.Run()
	//
	//f := KafkaAuto{
	//	Delay: 45 * time.Second,
	//	Topic: []string{config.TopicFace},
	//}
	//f.Run()

	a := KafkaAuto{
		Delay: 5 * time.Second,
		Topic: []string{config.TopicForman},
	}
	a.Run()

}

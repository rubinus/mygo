package kafka

import (
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/kafka/activity"
	"code.tvmining.com/tvplay/tvmq/kafka/epg"
	"code.tvmining.com/tvplay/tvmq/kafka/face"
	"code.tvmining.com/tvplay/tvmq/kafka/forman"
	"code.tvmining.com/tvplay/tvmq/kafka/imua"
	"code.tvmining.com/tvplay/tvmq/kafka/lottery"
)

type Dealer interface {
	Consumer()
}

type Services struct {
	Dealer Dealer
}

func Consumers() {
	//消费互动
	a := Services{
		Dealer: activity.Activity{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicActivity},
			GroupId: "group-id-activity",
		},
	}
	a.Dealer.Consumer()

	//消费epg给小i
	e := Services{
		Dealer: epg.EPG{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicEpg},
			GroupId: "group-id-epg",
		},
	}
	e.Dealer.Consumer()

	//消费人脸识别
	f := Services{
		Dealer: face.Face{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicFace},
			GroupId: "group-id-face",
		},
	}
	f.Dealer.Consumer()

	//消费开奖
	l := Services{
		Dealer: lottery.Lottery{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicLottery},
			GroupId: "group-id-lottery",
		},
	}
	l.Dealer.Consumer()

	//人工消息
	fo := Services{
		Dealer: forman.Forman{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicForman},
			GroupId: "group-id-socialtv_im_queue",
		},
	}
	fo.Dealer.Consumer()

	//人工消息
	im := Services{
		Dealer: imua.Imua{
			Brokers: config.KafkaHostPort,
			Topic:   []string{config.TopicIMua},
			GroupId: "group-id-socialtv_im_ua",
		},
	}
	im.Dealer.Consumer()
}

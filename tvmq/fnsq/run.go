package fnsq

import (
	"code.tvmining.com/tvplay/tvmq/config"
	"code.tvmining.com/tvplay/tvmq/fnsq/chat"
	"code.tvmining.com/tvplay/tvmq/fnsq/newAuthUser"
	"code.tvmining.com/tvplay/tvmq/fnsq/nsgift"
	"code.tvmining.com/tvplay/tvmq/fnsq/traceinfonsq"
)

type Dealer interface {
	Consumer()
}

type Services struct {
	Dealer Dealer
}

func Consumers() {
	n := Services{
		Dealer: newAuthUser.NewAuthUser{
			Topic:   config.TopicNewAuthUser,
			Channel: "newAuthUser-channel",
		},
	}
	n.Dealer.Consumer()

	g := Services{
		Dealer: nsgift.Gift{
			Topic:   config.TopicGift,
			Channel: "nsgift-channel",
		},
	}
	g.Dealer.Consumer()

	c := Services{
		Dealer: chat.Chat{
			Topic:   config.TopicChat,
			Channel: "chat-channel",
		},
	}
	c.Dealer.Consumer()

	//链路追踪消息
	ti := Services{
		Dealer: traceinfonsq.Traceinfo{
			Topic:   config.TraceInfo,
			Channel: "traceinfo-channel",
		},
	}
	ti.Dealer.Consumer()

}

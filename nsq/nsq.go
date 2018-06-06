package nsq

import (
	"fmt"
	"time"

	"github.com/nsqio/go-nsq"
)

func Producer(topic string, address string) {
	producer, err := nsq.NewProducer(address, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 5)
		str := time.Now().String()
		err := publish(producer, topic, str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("发送成功", str)
		}

	}
	//关闭
}

//发布消息
func publish(producer *nsq.Producer, topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return nil
		}
		err = producer.Publish(topic, []byte(message)) // 发布消息
		if err != nil {
			fmt.Println(err, "发送失败")
		}
		return err
	}
	return fmt.Errorf("producer is nil", err)
}

//处理消息
func delMsg(msg *nsq.Message) error {
	fmt.Println("接收到NSQ", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

//初始化消费者
func Consumer(topic string, channel string, address string, mode int) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second //设置重连时间
	cfg.MaxInFlight = 2
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		panic(err)
	}
	//c.SetLogger(nil, 0)        //屏蔽系统日志
	c.AddHandler(nsq.HandlerFunc(delMsg)) // 添加消费者接口

	if mode == 1 {
		//建立NSQLookupd连接
		if err := c.ConnectToNSQLookupd(address); err != nil {
			panic(err)
		}
	} else if mode == 2 {
		if err := c.ConnectToNSQD(address); err != nil {
			panic(err)
		}
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接

}

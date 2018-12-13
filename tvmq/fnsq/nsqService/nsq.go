package nsqService

import (
	"fmt"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"github.com/kataras/iris/core/errors"
	"github.com/nsqio/go-nsq"
)

const limitConn = 50

var (
	nsqClientChan  chan *nsq.Producer
	nsqClientQueue []*nsq.Producer
)

func init() {
	nsqClientChan = make(chan *nsq.Producer, 20000)
	nsqClientQueue = []*nsq.Producer{}

	ncChanchan := make(chan chan *nsq.Producer, limitConn)
	go func() {
		for sessionCh := range ncChanchan {
			if p, ok := <-sessionCh; ok {
				nsqClientQueue = append(nsqClientQueue, p)
			}
		}
	}()

	for i := 0; i < limitConn; i++ {
		ncChanchan <- createNsqClient(config.NsqHostProt)
	}

	go func() {
		for {
			if len(nsqClientChan) < 20000 {
				for _, p := range nsqClientQueue {
					if p != nil {
						nsqClientChan <- p
					}
				}
			}
			time.Sleep(limitConn * time.Millisecond)
			//fmt.Println(len(nsqClientChan),"--nsqchan--")
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("init nsq producer to nsqClientChan ...", len(nsqClientChan))
	}()
}

func createNsqClient(address string) chan *nsq.Producer {
	out := make(chan *nsq.Producer)
	go func() {
		pro, err := nsq.NewProducer(address, nsq.NewConfig())
		if err != nil {
			fmt.Println("NewProducer error...", err)
			out <- nil
			return
		}
		pro.SetLogger(nil, 2)

		out <- pro
	}()
	return out
}

func Producer(topic string, msg string) chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)

		producer := <-nsqClientChan
		err := publish(producer, topic, msg)
		if err != nil {
			//fmt.Println(topic,"--发送到NSQ失败--",err)
			out <- false
		} else {
			//fmt.Println(topic,"==发送到NSQ成功==", msg)
			out <- true
		}

	}()
	return out
	//关闭
}

//发布消息
func publish(producer *nsq.Producer, topic string, message string) error {
	var err error
	if producer != nil {
		if message == "" { //不能发布空串，否则会导致error
			return errors.New("message is empty")
		}
		doneChan := make(chan *nsq.ProducerTransaction, 100)
		err = producer.PublishAsync(topic, []byte(message), doneChan) // 发布消息
		<-doneChan
		//fmt.Printf("%s",r)
		return err
	}
	return errors.New("fnsq producer is nil")
}

//初始化消费者
func Consumer(topic string, channel string, address string, mode int, f nsq.HandlerFunc) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second //设置重连时间
	cfg.MaxInFlight = 20
	c, err := nsq.NewConsumer(topic, channel, cfg) // 新建一个消费者
	if err != nil {
		fmt.Println("===NewConsumer===", err)
		return
	}
	//c.SetLogger(nil, 0)        //屏蔽系统日志
	c.SetLogger(nil, 2)
	c.AddHandler(f) // 添加消费者接口

	if mode == 1 {
		//建立NSQLookupd连接
		if err := c.ConnectToNSQLookupd(address); err != nil {
			fmt.Println("===ConnectToNSQLookupd===", err)
		}
	} else if mode == 2 {
		if err := c.ConnectToNSQD(address); err != nil {
			fmt.Println("===ConnectToNSQD===", err)
		}
	}

	//建立多个nsqd连接
	// if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
	//  panic(err)
	// }

	// 建立一个nsqd连接

}

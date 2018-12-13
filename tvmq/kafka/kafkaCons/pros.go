package kafkaCons

import (
	"fmt"
	"os"
	"time"

	"code.tvmining.com/tvplay/tvmq/config"
	"github.com/Shopify/sarama"
)

const limitConn = 20

var (
	apChan  chan *sarama.AsyncProducer
	apQueue []*sarama.AsyncProducer
)

//创建连接池client，每次调用方法时从pool中取client
func init() {
	apChan = make(chan *sarama.AsyncProducer, 10000)
	apQueue = []*sarama.AsyncProducer{}

	ncChanchan := make(chan chan *sarama.AsyncProducer, limitConn)
	go func() {
		for sessionCh := range ncChanchan {
			if p, ok := <-sessionCh; ok {
				apQueue = append(apQueue, p)
			}
		}
	}()

	for i := 0; i < limitConn; i++ {
		ncChanchan <- createAsyncClient(config.KafkaHostPort)
	}

	go func() {
		for {
			if len(apChan) < 10000 {
				for _, p := range apQueue {
					if p != nil {
						apChan <- p
					}
				}
			}
			time.Sleep(50 * time.Millisecond)
			//fmt.Println(len(apChan),"--apChan--")

		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("init kafka async producer to apChan ...", len(apChan))
	}()

}

func createAsyncClient(address []string) chan *sarama.AsyncProducer {
	out := make(chan *sarama.AsyncProducer)
	go func() {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		p, err := sarama.NewAsyncProducer(address, config)
		if err != nil {
			fmt.Printf("sarama.NewSyncProducer err:%s \n", err)
			out <- nil
			return
		}
		out <- &p
	}()
	return out
}

// asyncProducer 异步生产者
// 并发量大时，必须采用这种方式
func AsyncProducer(topics, value string) {
	ptr := <-apChan
	p := *ptr
	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					fmt.Println("asyn send=", err, topics, value)
				}
			case <-success:
				//fmt.Printf("Partition:%d\nOffset:%d\n%s\n%s",s.Partition,s.Offset,s.Value,s.Timestamp)
				fmt.Fprintln(os.Stdout, "\n", value, "==done success==", topics)
			}
		}
	}(p)

	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(value),
	}
	p.Input() <- msg

}

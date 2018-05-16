package kafka

import (
	"fmt"
	"log"
	"os"

	"time"

	"os/signal"

	"github.com/Shopify/sarama" //support automatic consumer-group rebalancing and offset tracking
	"github.com/bsm/sarama-cluster"
)

var (
	topics = "topics_go_test"
)

func GetKafkaInfo(address []string) {
	config := sarama.NewConfig()
	client, err := sarama.NewClient(address, config)
	if err != nil {
		panic("client create error")
	}
	defer client.Close()
	//获取主题的名称集合
	topics, err := client.Topics()
	if err != nil {
		panic("get topics err")
	}
	for _, e := range topics {
		fmt.Println(e)
	}
	//获取broker集合
	brokers := client.Brokers()
	//输出每个机器的地址
	for _, broker := range brokers {
		fmt.Println(broker.Addr())
	}
}

func Consumer(brokers, topics []string, groupId string) {
	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			//fmt.Printf("%t", msg)
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}

// asyncProducer 异步生产者
// 并发量大时，必须采用这种方式
func AsyncProducer(address []string, topics, asrcValue string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Timeout = 2 * time.Second
	p, err := sarama.NewAsyncProducer(address, config)
	if err != nil {
		fmt.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}

	//必须有这个匿名函数内容
	go func(p sarama.AsyncProducer) {
		errors := p.Errors()
		success := p.Successes()
		for {
			select {
			case err := <-errors:
				if err != nil {
					fmt.Println("异步发送后===", err)
				}
			case <-success:
				fmt.Fprintln(os.Stdout, asrcValue, "=====成功返回了====")
			}
		}
	}(p)

	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(asrcValue),
	}
	p.Input() <- msg

}

func TestKafka() {

	Address := []string{"106.15.228.49:9092"}
	topic := []string{"topics__abc_go_test"}

	GetKafkaInfo(Address)
	go Consumer(Address, topic, "group-1100101109")

	asrcValue := "开始发送async-goroutine: this is a message. index=22221"
	AsyncProducer(Address, topic[0], asrcValue)

	time.Sleep(1 * time.Second)
	go AsyncProducer(Address, topic[0], asrcValue)
	time.Sleep(5 * time.Second)

	fmt.Println("5秒了系统退出")
}

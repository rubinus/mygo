package kafkaCons

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
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

func Consumer(brokers, topics []string, groupId string, callback func(*cluster.Consumer, chan os.Signal)) {
	config := cluster.NewConfig()
	//config.Consumer.Return.Errors = true
	//config.Group.Return.Notifications = true

	config.Group.Mode = cluster.ConsumerModePartitions

	consumer, err := cluster.NewConsumer(brokers, groupId, topics, config)
	if err != nil {
		log.Println(brokers, err.Error())
		//panic(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	//go func() {
	//	for err := range consumer.Errors() {
	//		log.Printf("Error: %s\n", err.Error())
	//	}
	//}()

	// consume notifications
	//go func() {
	//	for ntf := range consumer.Notifications() {
	//		log.Printf("Kafka connection: %+v\n", ntf)
	//	}
	//}()

	callback(consumer, signals)
}

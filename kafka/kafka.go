package kafka

import (
	"fmt"
	"log"
	"os"

	"time"

	"os/signal"

	"context"
	"strconv"
	"yaosocket/backend"
	"yaosocket/config"

	"github.com/Shopify/sarama" //support automatic consumer-group rebalancing and offset tracking
	"github.com/bitly/go-simplejson"
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
				fmt.Fprintf(os.Stdout, "=====%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
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

	//Address := []string{"10.20.80.22:9092"}
	//topic := []string{"activity"}

	//GetKafkaInfo(Address)
	//Consumer(Address, topic, "group1-11001011092222444")

	//asrcValue := `{"nickname":"涛涛GRT","headimgurl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTIOYJOrUGCdDufvEA3te2YrwW0XVRtBpwa8Aic3Kcf6X9Sq6dukJFwxxNeKGo7gdsibhoEMU9DzvpYQ/132","userid":"5b973536a33c300001007749"}`
	//
	//AsyncProducer(Address, topic[0], asrcValue)
	//
	//go func() {
	//	for range time.Tick(60 * time.Second) {
	//		AsyncProducer(Address, topic[0], asrcValue)
	//	}
	//}()

	//time.Sleep(500 * time.Second)

	fmt.Println("5秒了系统退出")

	msg := `{"status":200,"content":{"hd_id":"5bac5c2f004f9703","title":"东阿阿胶","video":"https://q-cdn.mtq.tvm.cn/liuxf/tvmini/config/activity/video/deej.mp4","countdown":"60","icon":"https://q-cdn.mtq.tvm.cn/liuxf/tvmini/config/activity/imgs/deej1.png","pictures":"https://q-cdn.mtq.tvm.cn/liuxf/tvmini/config/activity/imgs/deej.png","type":"0","start_time":1538029727,"question":null}}`

	var errStr string
	ctx, cancel := context.WithTimeout(context.Background(), config.GoroutineTimeout)
	defer cancel()

	value := fmt.Sprintf("%s", msg)

	key := fmt.Sprintf("%s:%s", config.CurrentActivity, config.Minappid)
	json, err := simplejson.NewJson([]byte(msg))
	if err != nil {
		return
	}
	countdown, err := json.Get("content").Get("countdown").String()
	if err != nil {
		return
	}
	ttl, err := strconv.Atoi(countdown)
	if err != nil {
		return
	}
	ch := backend.SaveActvityToRedis(ctx, key, value, ttl)
	select {
	case <-ctx.Done():
		errStr = "超时SaveActvityToRedis"
	case ok := <-ch:
		if !ok {
			errStr = "保存互动失败"
		}
	}
	if errStr != "" {
		fmt.Println(errStr)
	} else {
		fmt.Println("save activity successful")
	}

}

//Nsq发送测试
package main

import (
	"fmt"

	"time"

	"github.com/nsqio/go-nsq"
)

// 主函数
func main() {
	strIP1 := "localhost:4150"
	producer, err := nsq.NewProducer(strIP1, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	for {
		time.Sleep(time.Second * 5)
		str := time.Now().String()
		err := Publish(producer, "test1", str)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("发送成功", str)
		}

	}
	//关闭
}

//发布消息
func Publish(producer *nsq.Producer, topic string, message string) error {
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

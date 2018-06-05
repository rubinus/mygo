//Nsq发送测试
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nsqio/go-nsq"
)

// 主函数
func main() {
	strIP1 := "localhost:4150"
	producer, err := nsq.NewProducer(strIP1, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	running := true

	//读取控制台输入
	reader := bufio.NewReader(os.Stdin)
	for running {
		data, _, _ := reader.ReadLine()
		command := string(data)
		if command == "stop" {
			running = false
		}

		err := Publish(producer, "test1", command)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("发送成功", command)
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

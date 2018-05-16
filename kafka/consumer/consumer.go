package main

import (
	"mygo/kafka"
)

func main() {
	Address := []string{"106.15.228.49:9092"}
	topic := []string{"topics__abc_go_test"}
	kafka.Consumer(Address, topic, "group-110010109")
	//time.Sleep(10 * time.Second)
}

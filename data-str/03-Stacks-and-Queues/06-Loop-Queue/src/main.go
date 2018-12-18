package main

import (
	"fmt"
	"mygo/data-str/03-Stacks-and-Queues/06-Loop-Queue/src/ArrayQueue"
)

func main() {
	queue := ArrayQueue.GetArrayQueue(20)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}

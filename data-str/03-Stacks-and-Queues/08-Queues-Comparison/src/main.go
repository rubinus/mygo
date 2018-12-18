package main

import (
	"fmt"
	"math/rand"
	"mygo/data-str/03-Stacks-and-Queues/08-Queues-Comparison/src/ArrayQueue"
	"mygo/data-str/03-Stacks-and-Queues/08-Queues-Comparison/src/LoopQueue"
	"mygo/data-str/03-Stacks-and-Queues/08-Queues-Comparison/src/Queue"
	"time"
)

func testQueue(queue Queue.Queue, opCount int) float64 {
	startTime := time.Now()

	for i := 0; i < opCount; i++ {
		queue.Enqueue(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	return time.Now().Sub(startTime).Seconds()
}

func main() {
	opCount := 100000

	loopQueue := LoopQueue.GetLoopQueue(20)
	time1 := testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time1, "s")

	arrayQueue := ArrayQueue.GetArrayQueue(20)
	time := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time, "s")

	/**
	 * 测试结果：
	 * ArrayQueue, time: 7.94818435 s
	 * LoopQueue, time: 0.020557388 s
	 */
}

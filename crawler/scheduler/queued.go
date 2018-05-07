package scheduler

import (
	"fmt"
	"mygo/crawler/engine"
	"time"
)

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (qs *QueueScheduler) WokeChan() chan engine.Request {
	return make(chan engine.Request)
}

func (qs *QueueScheduler) Submit(r engine.Request) {
	qs.RequestChan <- r
}

func (qs *QueueScheduler) WorkReady(w chan engine.Request) {
	qs.WorkerChan <- w
}

//A调用 B接口，对B进行go并发时，只需要把 Bi（输入），Bo（输出）
//把Bi放到一个channel的队列中，一直在for中接收这个channel的值，然后调用B
//B的处理结果放到 Bo这个channel中，由A一直接着

func (qs *QueueScheduler) Start() {
	//初始化2个chan chan集装箱
	qs.RequestChan = make(chan engine.Request)
	qs.WorkerChan = make(chan chan engine.Request) //存放--执行收到的请求

	//在调度中开goruntine，只有一个并开始循环
	go func() {
		//创建队列
		var queueRequest []engine.Request
		var queueWork []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(queueRequest) > 0 && len(queueWork) > 0 { //存放的，和执行的都有，各取出一个
				activeRequest = queueRequest[0]
				activeWork = queueWork[0]
			}
			select {
			case r := <-qs.RequestChan: //接收到有了请求，先放到队列中
				queueRequest = append(queueRequest, r)
			case w := <-qs.WorkerChan: //work接收到了，也放到队列中
				queueWork = append(queueWork, w)
			case activeWork <- activeRequest: //把当前的请求给处理的work chan 从队列中去掉
				queueRequest = queueRequest[1:]
				queueWork = queueWork[1:]
			case <-time.After(time.Millisecond * 100):
				fmt.Println("============================time out============================time out==")
			}
		}
	}()
}

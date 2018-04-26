package scheduler

import "mygo/crawler/engine"

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
//把Bi放到到一个channel的队列中，在B中for它，一直接收这个channel的值
//处理结果放到 Bo这个channel中，由A一直接着

func (qs *QueueScheduler) Run() {
	qs.WorkerChan = make(chan chan engine.Request)
	qs.RequestChan = make(chan engine.Request)
	//在调度中开goruntine
	go func() {
		//创建队列
		var queueRequest []engine.Request
		var queueWork []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(queueRequest) > 0 && len(queueWork) > 0 {
				activeRequest = queueRequest[0]
				activeWork = queueWork[0]
			}
			select {
			case r := <-qs.RequestChan: //接收到有了请求
				queueRequest = append(queueRequest, r)
			case w := <-qs.WorkerChan: //work接收到了
				queueWork = append(queueWork, w)
			case activeWork <- activeRequest: //从队列中去掉
				queueRequest = queueRequest[1:]
				queueWork = queueWork[1:]
			}
		}
	}()
}

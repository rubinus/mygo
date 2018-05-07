package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Schduler //结构体中使用接口变量
	WorkCount int
}

type Schduler interface {
	Submit(Request)
	WokeChan() chan Request
	Start()
	ReadyNotify //接口套接口
}

type ReadyNotify interface {
	WorkReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult) //create接收的chan
	c.Scheduler.Start()           //队列先运行

	for i := 0; i < c.WorkCount; i++ { //输入的chan有多个，输出的chan只有一个
		CreateWorker(c.Scheduler.WokeChan(), out, c.Scheduler)
	}

	for _, v := range seeds {
		c.Scheduler.Submit(v) //调用结构体的变量中的方法，送出请求到scheduler，放入到in的chan中
	}

	itemCount := 0
	for {
		result := <-out
		for _, v := range result.Requests {
			c.Scheduler.Submit(v)
		}
		for _, v := range result.Items {
			fmt.Printf("%d-----items-----%v\n", itemCount, v)
			itemCount++
		}
	}

}

func CreateWorker(in chan Request, out chan ParseResult, notify ReadyNotify) {
	go func() {
		//并发创建worker
		for { //创建了100个输入chan，和 100个chan的集装箱，一直创建下去
			notify.WorkReady(in) //创建一个输入chan 然后送给chan 集装箱
			r := <-in            //从in中接收，谁向里面放，由请求时，放入到这个in中
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}

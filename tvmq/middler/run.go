package middler

type Scheduler interface {
	MakeWorkChan() chan *RequestBody
	WorkReady(chan *RequestBody)
	Start()
}

type Submiter interface {
	Submit(*RequestBody)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	Submiter    Submiter
	WorkerCount int
}

func (c *ConcurrentEngine) Run() {
	go func() {
		c.Scheduler.Start() //调度队列先运行
		//fmt.Println("-------仅启动时跑一次--------")
		for i := 0; i < c.WorkerCount; i++ { //准备work chan chan用来收集chan body
			CreateWorker(c.Scheduler.MakeWorkChan(), c.Scheduler)
		}
	}()

}

func CreateWorker(in chan *RequestBody, s Scheduler) {
	go func() {
		//并发创建worker
		for { //创建了100个输入chan，和 100个chan的集装箱，一直创建下去
			s.WorkReady(in) //把创建好的workchan通知start已经准备好了
			DoWorker(in)
		}

	}()
}

func MiddRun() (Submiter, Submiter, Submiter, Submiter) {
	//auth
	authPr := &RequestAuth{}
	authEngine := ConcurrentEngine{
		Scheduler:   authPr,
		Submiter:    authPr,
		WorkerCount: 32,
	}
	authEngine.Run()

	//chat
	chatPr := &RequestChat{}
	chatEngine := ConcurrentEngine{
		Scheduler:   chatPr,
		Submiter:    chatPr,
		WorkerCount: 32,
	}
	chatEngine.Run()

	//gift
	giftPr := &RequestGift{}
	giftEngine := ConcurrentEngine{
		Scheduler:   giftPr,
		Submiter:    giftPr,
		WorkerCount: 32,
	}
	giftEngine.Run()

	//diss
	dissPr := &RequestDiss{}
	dissEngine := ConcurrentEngine{
		Scheduler:   dissPr,
		Submiter:    dissPr,
		WorkerCount: 32,
	}
	dissEngine.Run()

	return authEngine.Submiter, chatEngine.Submiter, giftEngine.Submiter, dissEngine.Submiter
}

package scheduler

import "mygo/crawler/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) WokeChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {

}

func (s *SimpleScheduler) Start() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}

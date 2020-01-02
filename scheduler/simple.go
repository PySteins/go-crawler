package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	WorkChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkChan(in chan engine.Request) {
	s.WorkChan = in
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() {
		s.WorkChan <- r
	}()
}

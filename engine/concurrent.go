package engine

import "log"

type Scheduler interface {
	Submit(Request)
	GetWorkerChan() chan Request
	Run()
	Ready
}

type Ready interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, request := range result.Requests {
			if request.Url != "" {
				e.Scheduler.Submit(request)
			} else {
				go func() {
					e.ItemChan <- request.Item
					log.Printf("Got item %v", request.Item)
				}()
			}
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, r Ready) {
	go func() {
		for {
			r.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

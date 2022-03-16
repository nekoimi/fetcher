package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type ReadyNotifier interface {
	WorkerReady(worker chan Request)
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

func (engine *ConcurrentEngine) Run(seeds []Request) {
	out := make(chan ParseResult)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(engine.Scheduler.WorkerChan(), out, engine.Scheduler)
	}

	for _, s := range seeds {
		engine.Scheduler.Submit(s)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				engine.ItemChan <- item
			}()
		}
		for _, request := range result.Requests {
			engine.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, reader ReadyNotifier) {
	go func() {
		for {
			reader.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}

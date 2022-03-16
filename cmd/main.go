package main

import (
	"strconv"

	"github.com/nekoimi/fetcher/engine"
	"github.com/nekoimi/fetcher/meilele"
	"github.com/nekoimi/fetcher/persist"
	"github.com/nekoimi/fetcher/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		ItemChan:    persist.ItemServer(),
		WorkerCount: 10,
	}

	var requests []engine.Request
	for i := 1; i < 47; i++ {
		requests = append(requests, engine.Request{
			Url:       "https://www.meilele.com/category-keting/list-p" + strconv.Itoa(i) + "/?from=page",
			ParseFunc: meilele.ParseGoodList,
		})
	}

	e.Run(requests)
}

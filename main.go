package main

import (
	"fmt"
	. "go-learning/concurrency-patterns/worker"
	"sync"
)

func main() {
	TestWorkerPattern()
}

//func TestLoadBalancer() {
//	// 	import . "go-learning/loadbalancer" . "time"
//	p := make(Pool, 5)
//	done := make(chan *Worker)
//	for j := 0; j < 5; j++ {
//		requests := make(chan Request, 10)
//		p[j] = NewWorker(requests, 0, 0)
//		go p[j].Work(done)
//	}
//	b := NewBalancer(p, done)
//	req := make(chan Request)
//	go Requester(req)
//	go b.Balance(req)
//	Sleep(100 * Second)
//}

func TestWorkerPattern() {
	bufferSize := 100
	dispatcher := NewDispatcher(bufferSize)
	workers := 3
	for i := 0; i < workers; i++ {
		w := NewPreffixSuffixWorker(i, fmt.Sprintf("WorkerID: %d ->", i), "World")
		dispatcher.LaunchWorker(w)
	}

	requests := 10
	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest("(Msg_id: %d) -> Hello", i, &wg)
		dispatcher.MakeRequest(req)
	}
	dispatcher.Stop()
	wg.Wait()
}

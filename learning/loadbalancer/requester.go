package loadbalancer

import (
	"fmt"
	"math/rand"
	. "time"
)

var nWorker = 100

type Request struct {
	fn func() int
	c  chan int
}

func Requester(work chan<- Request) {
	c := make(chan int)
	for {
		Sleep(Duration(rand.Int63n(20)) * Second)
		work <- Request{workFn, c}
		result := <-c
		furtherProcess(result)
	}
}

func workFn() int {
	fmt.Println("work fn")
	return 0
}

func furtherProcess(result int) {
	fmt.Println("result", result)
	fmt.Println()
}

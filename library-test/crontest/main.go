package main

import (
	"fmt"
	//. "time"
	//
	"github.com/robfig/cron"
	"time"
)

var stopCh chan struct{}

func main() {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("Hello")
	})
	stopCh = make(chan struct{})
	go c.Start()
	go stop()
	<-stopCh
	fmt.Println("entries:", len(c.Entries()))
	c.Stop()
	e := c.Entries()
	fmt.Println(e[0].Prev)
	fmt.Println(e[0].Next)

	go c.Start()
	go stop()
	<-stopCh
	c.Stop()
}

func stop() {
	time.Sleep(4 * time.Second)
	stopCh <- struct{}{}
}

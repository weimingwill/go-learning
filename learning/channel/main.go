package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"go-learning/learning/channel/logger"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	const grs = 10
	var d device

	// Before this would fail because something would block on device, when
	// using the stdlib logger

	// var l log.Logger
	// l := log.New(&d, "", 0)
	// l.SetOutput(&d)

	// After we use our custom logger, and the app doesn't fail when device is
	// backed up.
	l := logger.New(&d, 10)

	for i := 0; i < grs; i++ {
		go func(i int) {
			for {
				l.Println(fmt.Sprintf("%d: log data\n", i))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	for {
		<-sigCh
		d.problem = !d.problem
	}
}

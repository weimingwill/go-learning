package worker

import "time"

type Dispatcher interface {
	LaunchWorker(w WorkderLauncher)
	MakeRequest(Request)
	Stop()
}

type dispatcher struct {
	in chan Request
}

func NewDispatcher(b int) Dispatcher {
	return &dispatcher{
		in: make(chan Request, b),
	}
}

func (d *dispatcher) LaunchWorker(w WorkderLauncher) {
	w.LaunchWorker(d.in)
}

func (d *dispatcher) Stop() {
	close(d.in)
}

func (d *dispatcher) MakeRequest(r Request) {
	timeout := time.After(5 * time.Second)
	select {
	case d.in <- r:
	case <-timeout:
		return
	}
}

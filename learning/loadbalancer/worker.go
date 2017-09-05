package loadbalancer

type Worker struct {
	requests chan Request
	pending  int
	index    int
}

func (w *Worker) Work(done chan *Worker) {
	for {
		req := <-w.requests
		req.c <- req.fn()
		done <- w
	}
}

func NewWorker(requests chan Request, pending int, index int) *Worker {
	return &Worker{
		requests: requests,
		pending:  pending,
		index:    index,
	}
}

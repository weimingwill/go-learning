package subscriber

import (
	"fmt"
	"io"
	"os"
	"time"
)

type writerSubscriber struct {
	id     int
	Writer io.Writer
	in     chan interface{}
}

func NewWriterSubscirber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	s := &writerSubscriber{
		id:     id,
		Writer: out,
		in:     make(chan interface{}),
	}

	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v/n", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}
	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}

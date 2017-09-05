package logger

import (
	"fmt"
	"io"
	"sync"
)

type Logger struct {
	ch chan string
	wg sync.WaitGroup
}

// New good practice to create a factory function to create Logger. Important
// to use a buffered channel so we can implement the "Drop Pattern". So we
// have a cap parameter.
func New(w io.Writer, cap int) *Logger {
	l := Logger{
		ch: make(chan string, cap),
	}

	// The goroutine to consume ch
	// Waitgroup is to manage the below goroutine
	l.wg.Add(1)
	go func() {
		for v := range l.ch {
			fmt.Fprintf(w, v)
		}
		l.wg.Done()
	}()

	return &l
}

func (l *Logger) Stop() {
	// We can just close the channel and wait for the consuming goroutine to
	// be done.
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Println(s string) {
	// Non-blocking send. Because we have the default case we will not block
	// if `l.ch` is at capacity, and just drop the log message.
	select {
	case l.ch <- s:
	default:
		fmt.Println("DROP")
	}
}

package subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	testingFunc func(string)
}

// Implement Write function to mock itself as io.Writer
func (m *mockWriter) Write(p []byte) (n int, err error) {
	m.testingFunc(string(p))
	return len(p), nil
}

func TestWriter(t *testing.T) {
	sub := NewWriterSubscirber(0, nil)
	msg := "Hello"

	var wg sync.WaitGroup
	wg.Add(1)
	stdoutPrinter := sub.(*writerSubscriber)
	stdoutPrinter.Writer = &mockWriter{
		testingFunc: func(res string) {
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Incorrect string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		// use wg.Done() to end the waiting, instead of using t.Fatal() to exist the program directly
		wg.Done()
		t.Error(err)
	}

	wg.Wait()
	sub.Close()
}

package worker

import (
	"fmt"
	"strings"
)

type WorkderLauncher interface {
	LaunchWorker(in chan Request)
}

type PreffixSuffixWorker struct {
	id      int
	prefixS string
	suffixS string
}

func NewPreffixSuffixWorker(id int, prefixS string, suffixS string) *PreffixSuffixWorker {
	return &PreffixSuffixWorker{
		id:      id,
		prefixS: prefixS,
		suffixS: suffixS,
	}
}

func (w *PreffixSuffixWorker) LaunchWorker(in chan Request) {
	w.prefix(w.append(w.uppercase(in)))
}

func (w *PreffixSuffixWorker) uppercase(in <-chan Request) <-chan Request {
	out := make(chan Request)

	go func() {
		for msg := range in {
			s, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = strings.ToUpper(s)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) append(in <-chan Request) <-chan Request {
	out := make(chan Request)
	go func() {
		for msg := range in {
			uppercaseStr, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Data = fmt.Sprintf("%s%s", uppercaseStr, w.suffixS)
			out <- msg
		}
		close(out)
	}()
	return out
}

func (w *PreffixSuffixWorker) prefix(in <-chan Request) {
	go func() {
		for msg := range in {
			uppercaseSuffixStr, ok := msg.Data.(string)
			if !ok {
				msg.Handler(nil)
				continue
			}
			msg.Handler(fmt.Sprintf("%s%s", w.prefixS, uppercaseSuffixStr))
		}
	}()
}

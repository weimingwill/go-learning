package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *GlobalHook) Fire(e *logrus.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func main() {
	l := logrus.New()
	l.Out = os.Stdout
	l.Formatter = &logrus.TextFormatter{DisableTimestamp: true, DisableColors: true}
	l.AddHook(&GlobalHook{})
	mystring = "first value"
	l.Info("first log")
	mystring = "another value"
	l.Info("second log")
}

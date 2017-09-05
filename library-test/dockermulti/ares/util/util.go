package util

import (
	"flag"
	"fmt"
)

type Options struct {
	Input string
}

func HelloPrint() {
	fmt.Println("Hello")
}

func Flag() *Options {
	opts := &Options{}
	flag.StringVar(&opts.Input, "in", "", "Path to source JSON file")
	flag.Parse()
	return opts
}

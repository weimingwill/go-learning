package worker

import (
	"fmt"
	"log"
	"sync"
)

type RequestHandler func(interface{})

type Request struct {
	Data    interface{}
	Handler RequestHandler
}

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	return Request{
		Data: "Hello",
		Handler: func(i interface{}) {
			defer wg.Done()
			s, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(s)
		},
	}
}
